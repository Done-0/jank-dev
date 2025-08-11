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
func NewRBACController(rbacService service.RBACService) *RBACController {
	return &RBACController{
		rbacService: rbacService,
	}
}

// AddPolicy 添加策略
// @Router /api/rbac/addPolicy [post]
func (rc *RBACController) AddPolicy(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AddPolicyRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.AddPolicy(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "add policy failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RemovePolicy 删除策略
// @Router /api/rbac/removePolicy [post]
func (rc *RBACController) RemovePolicy(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RemovePolicyRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.RemovePolicy(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "remove policy failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetAllPolicies 获取所有策略
// @Router /api/rbac/getAllPolicies [get]
func (rc *RBACController) GetAllPolicies(ctx context.Context, c *app.RequestContext) {
	response, err := rc.rbacService.GetAllPolicies(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "get all policies failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// AddRoleForUser 为用户添加角色
// @Router /api/rbac/addRoleForUser [post]
func (rc *RBACController) AddRoleForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AddRoleForUserRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.AddRoleForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "add role for user failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RemoveRoleForUser 删除用户角色
// @Router /api/rbac/removeRoleForUser [post]
func (rc *RBACController) RemoveRoleForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RemoveRoleForUserRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.RemoveRoleForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "remove role for user failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetRolesForUser 获取用户角色
// @Router /api/rbac/getRolesForUser [get]
func (rc *RBACController) GetRolesForUser(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetRolesForUserRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "query_params"), errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.GetRolesForUser(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "get roles for user failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Enforce 权限检查
// @Router /api/rbac/enforce [post]
func (rc *RBACController) Enforce(ctx context.Context, c *app.RequestContext) {
	req := new(dto.EnforceRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.Enforce(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "permission check failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
