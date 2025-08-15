/**
 * RBAC 相关常量定义
 */

// ===== RBAC 动作枚举 =====
export const RBAC_ACTION = {
  ALL: "*",
  GET: "GET",
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
  ACCESS: "access",
} as const;

// ===== RBAC 查询键 =====
export const RBAC_QUERY_KEYS = {
  CHECK_PERMISSION: "checkPermission",
  USER_ROLES: "userRoles",
} as const;

// ===== 类型定义 =====
export type RbacAction = typeof RBAC_ACTION[keyof typeof RBAC_ACTION];
