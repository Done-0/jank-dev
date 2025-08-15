/**
 * RBAC 权限检查 - Hooks
 * 采用主流批量权限检查方式，简化使用
 */

import { useMemo, useEffect, useState } from "react";
import { useAuthStore } from "@/stores/auth.store";
import { useUserStore } from "@/stores/user.store";
import { rbacService } from "@/services/rbac.service";
import { useQuery } from "@tanstack/react-query";
import { RBAC_QUERY_KEYS, RBAC_ACTION, type RbacAction } from "@/constants";
import type { CheckPermissionRequest, GetUserRolesRequest } from "@/types";

// ===== 权限项定义 =====
export interface PermissionItem {
  action: RbacAction;
  resource: string;
}

// ===== 主流批量权限检查 Hook =====
export const useBatchPermissions = (permissionItems: PermissionItem[]) => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();
  const [permissions, setPermissions] = useState<Record<string, boolean>>({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!isAuthenticated || !user?.id || permissionItems.length === 0) {
      setIsLoading(false);
      return;
    }

    const checkPermissions = async () => {
      try {
        setIsLoading(true);
        setError(null);

        // 批量检查所有权限
        const permissionResults = await Promise.all(
          permissionItems.map(({ resource, action }) =>
            rbacService.checkPermission({
              user_id: user.id.toString(),
              resource,
              action,
            })
          )
        );

        // 构建权限映射对象
        const perms = permissionItems.reduce(
          (acc, { action, resource }, index) => {
            acc[`${action}:${resource}`] = permissionResults[index]?.allowed || false;
            return acc;
          },
          {} as Record<string, boolean>
        );

        setPermissions(perms);
      } catch (err) {
        console.error("批量权限检查失败:", err);
        setError("权限检查失败");
      } finally {
        setIsLoading(false);
      }
    };

    checkPermissions();
  }, [isAuthenticated, user?.id, permissionItems]);

  return { permissions, isLoading, error };
};

// ===== 简单的单个权限检查 Hook =====
export const usePermission = (resource: string, action: RbacAction) => {
  const permissionItems = useMemo(() => [{ resource, action }], [resource, action]);
  const { permissions, isLoading, error } = useBatchPermissions(permissionItems);
  
  const hasPermission = permissions[`${action}:${resource}`] || false;

  return { hasPermission, isLoading, error };
};

// ===== 便捷的权限检查函数 =====
export const useHasPermission = (resource: string, action: RbacAction) => {
  return usePermission(resource, action);
};

// ===== 权限检查 Hook =====
export const useCheckPermission = (resource: string, action: RbacAction) => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();

  // 构建权限检查请求参数
  const request: CheckPermissionRequest = useMemo(() => ({
    user_id: user?.id?.toString() || "",
    resource,
    action,
  }), [user?.id, resource, action]);

  return useQuery({
    queryKey: [RBAC_QUERY_KEYS.CHECK_PERMISSION, request],
    queryFn: () => rbacService.checkPermission(request),
    enabled: isAuthenticated && !!user?.id,
    staleTime: 5 * 60 * 1000, // 5分钟缓存
    gcTime: 10 * 60 * 1000, // 10分钟垃圾回收
  });
};

// ===== 用户角色检查 Hook =====
export const useUserRoles = () => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();

  // 构建用户角色查询请求参数
  const request: GetUserRolesRequest = useMemo(() => ({
    user_id: user?.id?.toString() || "",
  }), [user?.id]);

  return useQuery({
    queryKey: [RBAC_QUERY_KEYS.USER_ROLES, user?.id],
    queryFn: () => rbacService.getUserRoles(request),
    enabled: isAuthenticated && !!user?.id,
    staleTime: 10 * 60 * 1000, // 10分钟缓存
    gcTime: 30 * 60 * 1000, // 30分钟垃圾回收
  });
};

// ===== Console 访问权限检查 =====
export const useCanAccessConsole = () => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();
  
  // 使用动态权限检查
  const { data: consolePermission, isLoading } = useCheckPermission("console", RBAC_ACTION.GET);

  // 检查 Console 访问权限 - 完全依赖后端动态校验
  const canAccess = useMemo(() => {
    if (!isAuthenticated || !user) return false;
    
    // 所有角色（包括 super_admin）都通过后端动态权限检查
    return consolePermission?.allowed || false;
  }, [isAuthenticated, user, consolePermission?.allowed]);

  return {
    canAccess,
    isLoading,
    user,
  };
};
