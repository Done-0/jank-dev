// Package impl 用户服务实现
// 创建者：Done-0
// 创建时间：2025-08-10
package impl

import (
	"context"
	"fmt"

	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/user"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/internal/utils/verification"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

var (
	registerLock      sync.Mutex // 用户注册锁，保护并发用户注册的操作
	passwordResetLock sync.Mutex // 修改密码锁，保护并发修改用户密码的操作
	logoutLock        sync.Mutex // 用户登出锁，保护并发用户登出操作
)

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userMapper mapper.UserMapper
}

// NewUserService 创建用户服务实例
func NewUserService(userMapperImpl mapper.UserMapper) service.UserService {
	return &UserServiceImpl{
		userMapper: userMapperImpl,
	}
}

// Register 用户注册逻辑
func (us *UserServiceImpl) Register(c *app.RequestContext, req *dto.RegisterRequest) (*vo.RegisterResponse, error) {
	registerLock.Lock()
	defer registerLock.Unlock()

	if !verification.VerifyEmailCode(c, req.EmailVerificationCode, req.Email) {
		return nil, fmt.Errorf("invalid email verification code")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.BizLogger(c).Errorf("password hashing failed: %v", err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	u := &user.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Role:     cfgs.AppConfig.User.DefaultRole,
	}

	if err := us.userMapper.RegisterUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("user registration failed for '%s': %v", req.Email, err)
		return nil, fmt.Errorf("user registration failed: %w", err)
	}

	userIDStr := strconv.FormatInt(u.ID, 10)
	if _, err := global.Enforcer.AddRoleForUser(userIDStr, u.Role); err != nil {
		us.userMapper.DeleteUser(c, userIDStr)
		return nil, fmt.Errorf("user registration failed due to RBAC system error: %w", err)
	}

	return &vo.RegisterResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Role:     u.Role,
	}, nil
}

// Login 登录用户逻辑
func (us *UserServiceImpl) Login(c *app.RequestContext, req *dto.LoginRequest) (*vo.LoginResponse, error) {
	u, err := us.userMapper.GetUserByEmail(c, req.Email)
	if err != nil {
		logger.BizLogger(c).Errorf("user not found for email '%s': %v", req.Email, err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		logger.BizLogger(c).Errorf("password verification failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("invalid password")
	}

	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	// 生成 access_token 和 refresh_token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		cfgs.AppConfig.JWT.IdentityKey: u.ID,
		"exp":                          time.Now().Add(time.Duration(cfgs.AppConfig.JWT.ExpireTime) * time.Hour).Unix(),
		"iat":                          time.Now().Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("access token generation failed: %v", err)
		return nil, fmt.Errorf("access token generation failed: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		cfgs.AppConfig.JWT.IdentityKey: u.ID,
		"exp":                          time.Now().Add(time.Duration(cfgs.AppConfig.JWT.RefreshExpire) * time.Hour).Unix(),
		"iat":                          time.Now().Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("refresh token generation failed: %v", err)
		return nil, fmt.Errorf("refresh token generation failed: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.UserCacheKeyPrefix, u.ID), accessTokenString, time.Duration(cfgs.AppConfig.JWT.ExpireTime)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to set access token cache: %v", err)
		return nil, fmt.Errorf("failed to set access token cache: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.UserRefreshCacheKeyPrefix, u.ID), refreshTokenString, time.Duration(cfgs.AppConfig.JWT.RefreshExpire)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to set refresh token cache: %v", err)
		return nil, fmt.Errorf("failed to set refresh token cache: %w", err)
	}

	return &vo.LoginResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// Logout 处理用户登出逻辑
func (us *UserServiceImpl) Logout(c *app.RequestContext) (*vo.LogoutResponse, error) {
	logoutLock.Lock()
	defer logoutLock.Unlock()

	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	userID, exists := c.Get(cfgs.AppConfig.JWT.IdentityKey)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get current user ID from context")
		return nil, fmt.Errorf("authentication required")
	}

	err = global.RedisClient.Del(context.Background(), fmt.Sprintf("%s:%d", consts.UserCacheKeyPrefix, userID.(int64))).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to delete access token cache: %v", err)
		return nil, fmt.Errorf("failed to delete access token cache: %w", err)
	}

	err = global.RedisClient.Del(context.Background(), fmt.Sprintf("%s:%d", consts.UserRefreshCacheKeyPrefix, userID.(int64))).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to delete refresh token cache: %v", err)
		return nil, fmt.Errorf("failed to delete refresh token cache: %w", err)
	}

	return &vo.LogoutResponse{
		Message: "登出成功",
	}, nil
}

// RefreshToken 刷新token
func (us *UserServiceImpl) RefreshToken(c *app.RequestContext, req *dto.RefreshTokenRequest) (*vo.RefreshTokenResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	// 解析refresh token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfgs.AppConfig.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		logger.BizLogger(c).Errorf("invalid refresh token: %v", err)
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.BizLogger(c).Errorf("invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims[cfgs.AppConfig.JWT.IdentityKey].(float64)
	if !ok {
		logger.BizLogger(c).Errorf("invalid user ID in token")
		return nil, fmt.Errorf("invalid user ID in token")
	}

	storedRefreshToken, err := global.RedisClient.Get(context.Background(), fmt.Sprintf("%s:%d", consts.UserRefreshCacheKeyPrefix, int64(userID))).Result()
	if err != nil {
		logger.BizLogger(c).Errorf("refresh token not found in cache: %v", err)
		return nil, fmt.Errorf("refresh token expired or invalid")
	}

	if storedRefreshToken != req.RefreshToken {
		logger.BizLogger(c).Errorf("refresh token mismatch")
		return nil, fmt.Errorf("refresh token invalid")
	}

	// 生成新的 access token 和 refresh token
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		cfgs.AppConfig.JWT.IdentityKey: int64(userID),
		"exp":                          time.Now().Add(time.Duration(cfgs.AppConfig.JWT.ExpireTime) * time.Hour).Unix(),
		"iat":                          time.Now().Unix(),
	})
	newAccessTokenString, err := newAccessToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("new access token generation failed: %v", err)
		return nil, fmt.Errorf("new access token generation failed: %w", err)
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		cfgs.AppConfig.JWT.IdentityKey: int64(userID),
		"exp":                          time.Now().Add(time.Duration(cfgs.AppConfig.JWT.RefreshExpire) * time.Hour).Unix(),
		"iat":                          time.Now().Unix(),
	})
	newRefreshTokenString, err := newRefreshToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("new refresh token generation failed: %v", err)
		return nil, fmt.Errorf("new refresh token generation failed: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.UserCacheKeyPrefix, int64(userID)), newAccessTokenString, time.Duration(cfgs.AppConfig.JWT.ExpireTime)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to update access token cache: %v", err)
		return nil, fmt.Errorf("failed to update access token cache: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.UserRefreshCacheKeyPrefix, int64(userID)), newRefreshTokenString, time.Duration(cfgs.AppConfig.JWT.RefreshExpire)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to update refresh token cache: %v", err)
		return nil, fmt.Errorf("failed to update refresh token cache: %w", err)
	}

	return &vo.RefreshTokenResponse{
		AccessToken:  newAccessTokenString,
		RefreshToken: newRefreshTokenString,
	}, nil
}

// GetProfile 获取用户资料逻辑
func (us *UserServiceImpl) GetProfile(c *app.RequestContext) (*vo.GetProfileResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	userID, exists := c.Get(cfgs.AppConfig.JWT.IdentityKey)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get current user ID from context")
		return nil, fmt.Errorf("authentication required")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("user not found: %v", err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &vo.GetProfileResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Role:     u.Role,
	}, nil

}

// Update 更新用户信息逻辑
func (us *UserServiceImpl) Update(c *app.RequestContext, req *dto.UpdateRequest) (*vo.UpdateResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	userID, exists := c.Get(cfgs.AppConfig.JWT.IdentityKey)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get user ID from context")
		return nil, fmt.Errorf("unable to get user ID from context")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user info: %v", err)
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	if req.Nickname != "" && u.Nickname != req.Nickname {
		existingUser, err := us.userMapper.GetUserByNickname(c, req.Nickname)
		if err == nil && existingUser.ID != u.ID {
			logger.BizLogger(c).Errorf("nickname '%s' is already in use", req.Nickname)
			return nil, fmt.Errorf("nickname '%s' is already in use", req.Nickname)
		}
	}

	// 更新用户信息
	if req.Nickname != "" {
		u.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		u.Avatar = req.Avatar
	}

	if err := us.userMapper.UpdateUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("user update failed for '%s': %v", u.Email, err)
		return nil, fmt.Errorf("user update failed: %w", err)
	}

	return &vo.UpdateResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Role:     u.Role,
	}, nil

}

// UpdateUserRole 管理员更新用户角色
func (us *UserServiceImpl) UpdateUserRole(c *app.RequestContext, req *dto.UpdateUserRoleRequest) (*vo.UpdateUserRoleResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	userID, exists := c.Get(cfgs.AppConfig.JWT.IdentityKey)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get current user ID from context")
		return nil, fmt.Errorf("authentication required")
	}

	// 验证当前用户是否有权限更新用户角色
	hasPermission, err := global.Enforcer.Enforce(userID.(string), string(c.Path()), string(c.Method()))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check user permissions: %v", err)
		return nil, fmt.Errorf("failed to verify permissions: %w", err)
	}
	if !hasPermission {
		logger.BizLogger(c).Warnf("user ID %s attempted to update user role without permission", userID.(string))
		return nil, fmt.Errorf("insufficient permissions: you do not have permission to update user roles")
	}

	// 防止用户修改自己的角色
	if userID.(string) == req.ID {
		logger.BizLogger(c).Warnf("user ID %s attempted to modify their own role", userID.(string))
		return nil, fmt.Errorf("cannot modify your own role")
	}

	// 获取目标用户信息
	targetUserID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid target user ID format: %v", err)
		return nil, fmt.Errorf("invalid target user ID format: %w", err)
	}

	targetUser, err := us.userMapper.GetUserByID(c, targetUserID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get target user info: %v", err)
		return nil, fmt.Errorf("failed to get target user info: %w", err)
	}

	// 更新用户角色并同步到 Casbin
	oldRole := targetUser.Role
	targetUser.Role = req.Role

	if err := us.userMapper.UpdateUser(c, targetUser); err != nil {
		logger.BizLogger(c).Errorf("failed to update user role for user '%s': %v", targetUser.Email, err)
		return nil, fmt.Errorf("failed to update user role: %w", err)
	}

	if oldRole != "" {
		if _, err := global.Enforcer.DeleteRoleForUser(req.ID, oldRole); err != nil {
			logger.BizLogger(c).Errorf("failed to remove old role from RBAC: %v", err)
		}
	}
	if _, err := global.Enforcer.AddRoleForUser(req.ID, req.Role); err != nil {
		logger.BizLogger(c).Errorf("failed to add new role to RBAC: %v", err)
		return nil, fmt.Errorf("failed to sync role to RBAC system: %w", err)
	}

	return &vo.UpdateUserRoleResponse{
		ID:       req.ID,
		Email:    targetUser.Email,
		Nickname: targetUser.Nickname,
		Role:     req.Role,
		Message:  fmt.Sprintf("User role successfully updated to %s", req.Role),
	}, nil
}

// ResetPassword 重置密码逻辑
func (us *UserServiceImpl) ResetPassword(c *app.RequestContext, req *dto.ResetPasswordRequest) (*vo.ResetPasswordResponse, error) {
	passwordResetLock.Lock()
	defer passwordResetLock.Unlock()

	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	userID, exists := c.Get(cfgs.AppConfig.JWT.IdentityKey)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get user ID from context")
		return nil, fmt.Errorf("unable to get user ID from context")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("user not found: %v", err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.OldPassword))
	if err != nil {
		logger.BizLogger(c).Errorf("old password verification failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("old password verification failed")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.BizLogger(c).Errorf("password hashing failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	u.Password = string(newPassword)

	if err := us.userMapper.UpdateUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("password update failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("password update failed: %w", err)
	}

	return &vo.ResetPasswordResponse{
		Message: "密码重置成功",
	}, nil
}

// ListUsers 获取用户列表逻辑
func (us *UserServiceImpl) ListUsers(c *app.RequestContext, req *dto.ListUsersRequest) (*vo.ListUsersResponse, error) {
	users, total, err := us.userMapper.ListUsers(c, req.PageNo, req.PageSize, req.Keyword, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user list: %v", err)
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}

	list := make([]*vo.UserItem, 0, len(users))
	for _, u := range users {
		list = append(list, &vo.UserItem{
			ID:       strconv.FormatInt(u.ID, 10),
			Email:    u.Email,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Role:     u.Role,
		})
	}

	return &vo.ListUsersResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}
