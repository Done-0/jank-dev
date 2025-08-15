import { LayoutDashboard, Users, Shield, Palette, Puzzle, FileText, FolderOpen, Settings } from "lucide-react";
import { Link } from "@tanstack/react-router";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
  useSidebar,
} from "@/components/ui/sidebar";
import type { GetProfileResponse } from "@/types";
import { CONSOLE_ROUTES } from "@/constants";
import { ConsoleUserMenu } from "@/components/layout/ConsoleUserMenu";

interface AppSidebarProps {
  activeTab: string;
  user: GetProfileResponse | null;
  onLogout: () => void;
}

export function AppSidebar({ activeTab, user, onLogout }: AppSidebarProps) {
  const { isMobile, setOpenMobile } = useSidebar();

  // Console 管理菜单配置
  const mainMenuItems = [
    {
      id: CONSOLE_ROUTES.ROOT,
      label: "控制台",
      icon: LayoutDashboard,
      route: CONSOLE_ROUTES.ROOT,
    },
    {
      id: CONSOLE_ROUTES.USERS,
      label: "用户管理",
      icon: Users,
      route: CONSOLE_ROUTES.USERS,
    },
    {
      id: CONSOLE_ROUTES.RBAC,
      label: "权限管理",
      icon: Shield,
      route: CONSOLE_ROUTES.RBAC,
    },
    {
      id: CONSOLE_ROUTES.POSTS,
      label: "文章管理",
      icon: FileText,
      route: CONSOLE_ROUTES.POSTS,
    },
    {
      id: CONSOLE_ROUTES.CATEGORIES,
      label: "分类管理",
      icon: FolderOpen,
      route: CONSOLE_ROUTES.CATEGORIES,
    },
  ] as const;

  // 系统设置菜单配置
  const systemMenuItems = [
    {
      id: CONSOLE_ROUTES.THEMES,
      label: "主题管理",
      icon: Palette,
      route: CONSOLE_ROUTES.THEMES,
    },
    {
      id: CONSOLE_ROUTES.PLUGINS,
      label: "插件管理",
      icon: Puzzle,
      route: CONSOLE_ROUTES.PLUGINS,
    },
    {
      id: CONSOLE_ROUTES.SYSTEM,
      label: "系统设置",
      icon: Settings,
      route: CONSOLE_ROUTES.SYSTEM,
    },
  ] as const;

  // 移动端点击链接后自动关闭侧边栏
  const handleLinkClick = () => {
    if (isMobile) {
      setOpenMobile(false);
    }
  };

  return (
    <Sidebar
      collapsible="offcanvas"
      className="border-r bg-sidebar text-sidebar-foreground"
    >
      <SidebarHeader className="bg-sidebar">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <Link
                to={CONSOLE_ROUTES.ROOT}
                className="flex items-center gap-2"
              >
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <LayoutDashboard className="size-4" />
                </div>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-semibold">Jank Console</span>
                  <span className="truncate text-xs text-sidebar-foreground/70">
                    后台管理系统
                  </span>
                </div>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent className="bg-sidebar">
        {/* 内容管理 */}
        <SidebarGroup>
          <SidebarGroupLabel>内容管理</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {mainMenuItems.map((item) => {
                const Icon = item.icon;
                return (
                  <SidebarMenuItem key={item.id}>
                    <SidebarMenuButton
                      asChild
                      tooltip={item.label}
                      isActive={activeTab === item.id}
                    >
                      <Link to={item.route} onClick={handleLinkClick}>
                        <Icon />
                        <span>{item.label}</span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        {/* 系统设置 */}
        <SidebarGroup>
          <SidebarGroupLabel>系统设置</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {systemMenuItems.map((item) => {
                const Icon = item.icon;
                return (
                  <SidebarMenuItem key={item.id}>
                    <SidebarMenuButton
                      asChild
                      tooltip={item.label}
                      isActive={activeTab === item.id}
                    >
                      <Link to={item.route} onClick={handleLinkClick}>
                        <Icon />
                        <span>{item.label}</span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter className="bg-sidebar">
        <SidebarMenu>
          <ConsoleUserMenu user={user} onLogout={onLogout} />
        </SidebarMenu>
      </SidebarFooter>

      <SidebarRail />
    </Sidebar>
  );
}
