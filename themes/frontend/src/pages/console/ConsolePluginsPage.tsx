import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Puzzle, Download, Settings, Play, Pause, Trash2 } from "lucide-react";

export function ConsolePluginsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">插件管理</h1>
          <p className="text-muted-foreground">
            管理系统插件和扩展功能
          </p>
        </div>
        <Button>
          <Download className="mr-2 h-4 w-4" />
          安装插件
        </Button>
      </div>

      <Tabs defaultValue="installed" className="space-y-4">
        <TabsList>
          <TabsTrigger value="installed">已安装插件</TabsTrigger>
          <TabsTrigger value="available">可用插件</TabsTrigger>
          <TabsTrigger value="settings">插件设置</TabsTrigger>
        </TabsList>

        <TabsContent value="installed" className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2">
            {/* 示例插件 1 */}
            <Card>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <div className="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
                      <Puzzle className="h-4 w-4 text-green-600" />
                    </div>
                    <CardTitle className="text-lg">内容统计插件</CardTitle>
                  </div>
                  <Badge variant="default">运行中</Badge>
                </div>
                <CardDescription>
                  提供文章和用户数据的统计分析功能
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">版本</span>
                    <span>v1.2.0</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">作者</span>
                    <span>Jank Team</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">状态</span>
                    <Switch defaultChecked />
                  </div>
                  <div className="flex items-center space-x-2">
                    <Button variant="outline" size="sm">
                      <Settings className="mr-2 h-4 w-4" />
                      设置
                    </Button>
                    <Button variant="outline" size="sm">
                      <Pause className="mr-2 h-4 w-4" />
                      停用
                    </Button>
                    <Button variant="outline" size="sm" className="text-destructive">
                      <Trash2 className="mr-2 h-4 w-4" />
                      卸载
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* 示例插件 2 */}
            <Card>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <div className="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                      <Puzzle className="h-4 w-4 text-blue-600" />
                    </div>
                    <CardTitle className="text-lg">SEO 优化插件</CardTitle>
                  </div>
                  <Badge variant="secondary">已停用</Badge>
                </div>
                <CardDescription>
                  自动优化文章的 SEO 设置和元数据
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">版本</span>
                    <span>v2.1.0</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">作者</span>
                    <span>Community</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">状态</span>
                    <Switch />
                  </div>
                  <div className="flex items-center space-x-2">
                    <Button variant="outline" size="sm">
                      <Settings className="mr-2 h-4 w-4" />
                      设置
                    </Button>
                    <Button variant="outline" size="sm">
                      <Play className="mr-2 h-4 w-4" />
                      启用
                    </Button>
                    <Button variant="outline" size="sm" className="text-destructive">
                      <Trash2 className="mr-2 h-4 w-4" />
                      卸载
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="available" className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2">
            {/* 可用插件示例 */}
            <Card>
              <CardHeader>
                <div className="flex items-center space-x-2">
                  <div className="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
                    <Puzzle className="h-4 w-4 text-purple-600" />
                  </div>
                  <CardTitle className="text-lg">评论系统插件</CardTitle>
                </div>
                <CardDescription>
                  为文章添加评论和互动功能
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">版本</span>
                    <span>v1.0.0</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">下载量</span>
                    <span>1.2k</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">评分</span>
                    <span>⭐⭐⭐⭐⭐ 4.8</span>
                  </div>
                  <Button className="w-full">
                    <Download className="mr-2 h-4 w-4" />
                    安装插件
                  </Button>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <div className="flex items-center space-x-2">
                  <div className="w-8 h-8 bg-orange-100 rounded-lg flex items-center justify-center">
                    <Puzzle className="h-4 w-4 text-orange-600" />
                  </div>
                  <CardTitle className="text-lg">邮件通知插件</CardTitle>
                </div>
                <CardDescription>
                  自动发送系统通知和邮件提醒
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">版本</span>
                    <span>v3.2.1</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">下载量</span>
                    <span>856</span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-muted-foreground">评分</span>
                    <span>⭐⭐⭐⭐ 4.2</span>
                  </div>
                  <Button className="w-full">
                    <Download className="mr-2 h-4 w-4" />
                    安装插件
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="settings" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>插件系统设置</CardTitle>
              <CardDescription>
                配置插件系统的全局设置和选项
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <div className="text-sm font-medium">自动更新插件</div>
                  <div className="text-xs text-muted-foreground">
                    自动检查并更新已安装的插件
                  </div>
                </div>
                <Switch defaultChecked />
              </div>

              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <div className="text-sm font-medium">插件安全检查</div>
                  <div className="text-xs text-muted-foreground">
                    安装前检查插件的安全性和兼容性
                  </div>
                </div>
                <Switch defaultChecked />
              </div>

              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <div className="text-sm font-medium">插件日志记录</div>
                  <div className="text-xs text-muted-foreground">
                    记录插件的运行日志和错误信息
                  </div>
                </div>
                <Switch />
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      {/* 插件统计 */}
      <div className="grid gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              已安装插件
            </CardTitle>
            <Puzzle className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">2</div>
            <p className="text-xs text-muted-foreground">
              当前安装的插件数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              运行中插件
            </CardTitle>
            <Play className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">1</div>
            <p className="text-xs text-muted-foreground">
              正在运行的插件数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              可用更新
            </CardTitle>
            <Download className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">0</div>
            <p className="text-xs text-muted-foreground">
              有可用更新的插件
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              插件商店
            </CardTitle>
            <Settings className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">50+</div>
            <p className="text-xs text-muted-foreground">
              商店中可用的插件
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
