import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Shield, Users, Key, Plus } from "lucide-react";

export function ConsoleRBACPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">权限管理</h1>
          <p className="text-muted-foreground">
            管理系统角色和权限配置
          </p>
        </div>
      </div>

      <Tabs defaultValue="roles" className="space-y-4">
        <TabsList>
          <TabsTrigger value="roles">角色管理</TabsTrigger>
          <TabsTrigger value="permissions">权限管理</TabsTrigger>
          <TabsTrigger value="policies">策略管理</TabsTrigger>
        </TabsList>

        <TabsContent value="roles" className="space-y-4">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between">
              <div>
                <CardTitle className="flex items-center gap-2">
                  <Users className="h-5 w-5" />
                  系统角色
                </CardTitle>
                <CardDescription>
                  管理系统中的用户角色定义
                </CardDescription>
              </div>
              <Button>
                <Plus className="mr-2 h-4 w-4" />
                添加角色
              </Button>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center justify-between p-4 border rounded-lg">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-red-100 rounded-lg flex items-center justify-center">
                      <Shield className="h-5 w-5 text-red-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">超级管理员</h3>
                      <p className="text-sm text-muted-foreground">拥有系统所有权限</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="destructive">super_admin</Badge>
                    <Button variant="outline" size="sm">编辑</Button>
                  </div>
                </div>

                <div className="flex items-center justify-between p-4 border rounded-lg">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                      <Users className="h-5 w-5 text-blue-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">普通用户</h3>
                      <p className="text-sm text-muted-foreground">基础用户权限</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="secondary">user</Badge>
                    <Button variant="outline" size="sm">编辑</Button>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="permissions" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Key className="h-5 w-5" />
                权限列表
              </CardTitle>
              <CardDescription>
                系统中所有可用的权限定义
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4 md:grid-cols-2">
                <div className="space-y-2">
                  <h4 className="font-medium">用户管理权限</h4>
                  <div className="space-y-1">
                    <Badge variant="outline">user:read</Badge>
                    <Badge variant="outline">user:write</Badge>
                    <Badge variant="outline">user:delete</Badge>
                  </div>
                </div>

                <div className="space-y-2">
                  <h4 className="font-medium">内容管理权限</h4>
                  <div className="space-y-1">
                    <Badge variant="outline">post:read</Badge>
                    <Badge variant="outline">post:write</Badge>
                    <Badge variant="outline">post:delete</Badge>
                  </div>
                </div>

                <div className="space-y-2">
                  <h4 className="font-medium">系统管理权限</h4>
                  <div className="space-y-1">
                    <Badge variant="outline">system:read</Badge>
                    <Badge variant="outline">system:write</Badge>
                    <Badge variant="outline">theme:manage</Badge>
                  </div>
                </div>

                <div className="space-y-2">
                  <h4 className="font-medium">插件管理权限</h4>
                  <div className="space-y-1">
                    <Badge variant="outline">plugin:read</Badge>
                    <Badge variant="outline">plugin:write</Badge>
                    <Badge variant="outline">plugin:execute</Badge>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="policies" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>访问策略</CardTitle>
              <CardDescription>
                基于角色的访问控制策略配置
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="text-center py-8 text-muted-foreground">
                  <Shield className="h-12 w-12 mx-auto mb-4 opacity-50" />
                  <p>策略管理功能开发中...</p>
                  <p className="text-sm">将支持细粒度的权限策略配置</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
