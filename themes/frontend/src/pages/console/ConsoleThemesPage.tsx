import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
import { Palette, Download, Settings, Eye } from "lucide-react";

export function ConsoleThemesPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">主题管理</h1>
          <p className="text-muted-foreground">
            管理系统主题和外观设置
          </p>
        </div>
        <Button>
          <Download className="mr-2 h-4 w-4" />
          安装主题
        </Button>
      </div>

      <div className="grid gap-6 md:grid-cols-2">
        {/* 当前激活主题 */}
        <Card className="border-primary">
          <CardHeader>
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Palette className="h-5 w-5 text-primary" />
                <CardTitle>Console 管理主题</CardTitle>
              </div>
              <Badge variant="default">当前使用</Badge>
            </div>
            <CardDescription>
              专为后台管理设计的现代化主题
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="aspect-video bg-gradient-to-br from-primary/20 to-primary/5 rounded-lg flex items-center justify-center">
                <div className="text-center">
                  <Palette className="h-8 w-8 mx-auto mb-2 text-primary" />
                  <p className="text-sm text-muted-foreground">主题预览</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between">
                <div className="space-y-1">
                  <p className="text-sm font-medium">版本 1.0.0</p>
                  <p className="text-xs text-muted-foreground">作者：Jank Team</p>
                </div>
                <div className="flex items-center space-x-2">
                  <Button variant="outline" size="sm">
                    <Eye className="mr-2 h-4 w-4" />
                    预览
                  </Button>
                  <Button variant="outline" size="sm">
                    <Settings className="mr-2 h-4 w-4" />
                    设置
                  </Button>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* 前端用户主题 */}
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Palette className="h-5 w-5" />
                <CardTitle>Frontend 用户主题</CardTitle>
              </div>
              <Badge variant="secondary">可用</Badge>
            </div>
            <CardDescription>
              面向普通用户的前端展示主题
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="aspect-video bg-gradient-to-br from-blue-100 to-purple-100 rounded-lg flex items-center justify-center">
                <div className="text-center">
                  <Palette className="h-8 w-8 mx-auto mb-2 text-blue-600" />
                  <p className="text-sm text-muted-foreground">主题预览</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between">
                <div className="space-y-1">
                  <p className="text-sm font-medium">版本 1.0.0</p>
                  <p className="text-xs text-muted-foreground">作者：Jank Team</p>
                </div>
                <div className="flex items-center space-x-2">
                  <Button variant="outline" size="sm">
                    <Eye className="mr-2 h-4 w-4" />
                    预览
                  </Button>
                  <Button size="sm">
                    启用
                  </Button>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* 主题设置 */}
      <Card>
        <CardHeader>
          <CardTitle>主题设置</CardTitle>
          <CardDescription>
            配置主题的全局设置和选项
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="flex items-center justify-between">
            <div className="space-y-0.5">
              <div className="text-sm font-medium">自动切换主题</div>
              <div className="text-xs text-muted-foreground">
                根据系统设置自动切换明暗主题
              </div>
            </div>
            <Switch />
          </div>

          <div className="flex items-center justify-between">
            <div className="space-y-0.5">
              <div className="text-sm font-medium">主题缓存</div>
              <div className="text-xs text-muted-foreground">
                启用主题资源缓存以提升加载速度
              </div>
            </div>
            <Switch defaultChecked />
          </div>

          <div className="flex items-center justify-between">
            <div className="space-y-0.5">
              <div className="text-sm font-medium">主题预编译</div>
              <div className="text-xs text-muted-foreground">
                预编译主题资源以优化性能
              </div>
            </div>
            <Switch defaultChecked />
          </div>
        </CardContent>
      </Card>

      {/* 主题统计 */}
      <div className="grid gap-4 md:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              已安装主题
            </CardTitle>
            <Palette className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">2</div>
            <p className="text-xs text-muted-foreground">
              系统中可用的主题数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              当前主题
            </CardTitle>
            <Settings className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">Console</div>
            <p className="text-xs text-muted-foreground">
              正在使用的主题名称
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              主题版本
            </CardTitle>
            <Download className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">v1.0.0</div>
            <p className="text-xs text-muted-foreground">
              当前主题的版本号
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
