// Package controller RBAC权限管理控制器
// 创建者：Done-0
// 创建时间：2025-08-08
package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/validator"
	"github.com/Done-0/jank/internal/utils/vo"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// RBACController RBAC权限管理控制器
type RBACController struct {
	rbacService service.RBACService
}

// NewRBACController 创建RBAC权限管理控制器
// 参数：
//
//	rbacService: RBAC权限管理服务
//
// 返回值：
//
//	*RBACController: RBAC权限管理控制器
func NewRBACController(rbacService service.RBACService) *RBACController {
	return &RBACController{
		rbacService: rbacService,
	}
}

// AddPolicy 添加策略
// @Router /api/v1/rbac/addPolicy [post]
func (rc *RBACController) AddPolicy(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AddPolicyRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrPolicyExists)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrPolicyExists)))
		return
	}

	response, err := rc.rbacService.AddPolicy(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPolicyExists)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RemovePolicy 删除策略
// @Router /api/v1/rbac/removePolicy [post]
func (rc *RBACController) RemovePolicy(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RemovePolicyRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	response, err := rc.rbacService.RemovePolicy(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// AddRoleForUser 为用户添加角色
// @Router /api/v1/rbac/addRoleForUser [post]
func (rc *RBACController) AddRoleForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AddRoleForUserRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	response, err := rc.rbacService.AddRoleForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RemoveRoleForUser 删除用户角色
// @Router /api/v1/rbac/removeRoleForUser [post]
func (rc *RBACController) RemoveRoleForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RemoveRoleForUserRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	response, err := rc.rbacService.RemoveRoleForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// AddRoleInheritance 添加角色继承关系
// @Router /api/v1/rbac/addRoleInheritance [post]
func (rc *RBACController) AddRoleInheritance(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AddRoleInheritanceRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	response, err := rc.rbacService.AddRoleInheritance(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RemoveRoleInheritance 删除角色继承关系
// @Router /api/v1/rbac/removeRoleInheritance [post]
func (rc *RBACController) RemoveRoleInheritance(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RemoveRoleInheritanceRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	response, err := rc.rbacService.RemoveRoleInheritance(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetRolesForUser 获取用户角色
// @Router /api/v1/rbac/getRolesForUser [get]
func (rc *RBACController) GetRolesForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetRolesForUserRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	response, err := rc.rbacService.GetRolesForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetPoliciesForRole 获取角色策略
// @Router /api/v1/rbac/getPoliciesForRole [get]
func (rc *RBACController) GetPoliciesForRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetPoliciesForRoleRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	response, err := rc.rbacService.GetPoliciesForRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetAllRoles 获取所有角色
// @Router /api/v1/rbac/getAllRoles [get]
func (rc *RBACController) GetAllRoles(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetAllRolesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	response, err := rc.rbacService.GetAllRoles(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrRoleNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetAllPolicies 获取所有策略
// @Router /api/v1/rbac/getAllPolicies [get]
func (rc *RBACController) GetAllPolicies(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetAllPoliciesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	response, err := rc.rbacService.GetAllPolicies(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPolicyNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
