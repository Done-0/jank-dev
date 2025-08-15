import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Search, Plus, MoreHorizontal } from "lucide-react";

export function ConsoleUsersPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">用户管理</h1>
          <p className="text-muted-foreground">
            管理系统用户和权限设置
          </p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          添加用户
        </Button>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>用户列表</CardTitle>
          <CardDescription>
            系统中所有注册用户的管理界面
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex items-center space-x-2 mb-4">
            <div className="relative flex-1 max-w-sm">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="搜索用户..."
                className="pl-8"
              />
            </div>
          </div>

          <div className="rounded-md border">
            <div className="p-4">
              <div className="grid grid-cols-4 gap-4 font-medium text-sm text-muted-foreground mb-4">
                <div>用户信息</div>
                <div>角色</div>
                <div>注册时间</div>
                <div>操作</div>
              </div>
              
              <div className="space-y-4">
                {/* 示例用户数据 */}
                <div className="grid grid-cols-4 gap-4 items-center py-2 border-b">
                  <div className="flex items-center space-x-3">
                    <div className="w-8 h-8 bg-primary/10 rounded-full flex items-center justify-center">
                      <span className="text-sm font-medium">A</span>
                    </div>
                    <div>
                      <p className="font-medium">admin@example.com</p>
                      <p className="text-sm text-muted-foreground">管理员</p>
                    </div>
                  </div>
                  <div>
                    <Badge variant="default">超级管理员</Badge>
                  </div>
                  <div className="text-sm text-muted-foreground">
                    2024-01-01
                  </div>
                  <div>
                    <Button variant="ghost" size="sm">
                      <MoreHorizontal className="h-4 w-4" />
                    </Button>
                  </div>
                </div>

                <div className="grid grid-cols-4 gap-4 items-center py-2 border-b">
                  <div className="flex items-center space-x-3">
                    <div className="w-8 h-8 bg-primary/10 rounded-full flex items-center justify-center">
                      <span className="text-sm font-medium">U</span>
                    </div>
                    <div>
                      <p className="font-medium">user@example.com</p>
                      <p className="text-sm text-muted-foreground">普通用户</p>
                    </div>
                  </div>
                  <div>
                    <Badge variant="secondary">普通用户</Badge>
                  </div>
                  <div className="text-sm text-muted-foreground">
                    2024-01-02
                  </div>
                  <div>
                    <Button variant="ghost" size="sm">
                      <MoreHorizontal className="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="flex items-center justify-between mt-4">
            <p className="text-sm text-muted-foreground">
              显示 1-2 条，共 2 条记录
            </p>
            <div className="flex items-center space-x-2">
              <Button variant="outline" size="sm" disabled>
                上一页
              </Button>
              <Button variant="outline" size="sm" disabled>
                下一页
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
