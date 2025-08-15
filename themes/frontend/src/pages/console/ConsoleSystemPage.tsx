import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Textarea } from "@/components/ui/textarea";
import { Settings, Database, Shield, Mail, Server } from "lucide-react";

export function ConsoleSystemPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">系统设置</h1>
          <p className="text-muted-foreground">
            配置系统的全局设置和参数
          </p>
        </div>
      </div>

      <Tabs defaultValue="general" className="space-y-4">
        <TabsList>
          <TabsTrigger value="general">基本设置</TabsTrigger>
          <TabsTrigger value="security">安全设置</TabsTrigger>
          <TabsTrigger value="email">邮件设置</TabsTrigger>
          <TabsTrigger value="database">数据库</TabsTrigger>
        </TabsList>

        <TabsContent value="general" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Settings className="h-5 w-5" />
                网站基本信息
              </CardTitle>
              <CardDescription>
                配置网站的基本信息和显示设置
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="site-name">网站名称</Label>
                  <Input id="site-name" defaultValue="Jank Console" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="site-url">网站地址</Label>
                  <Input id="site-url" defaultValue="https://example.com" />
                </div>
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="site-description">网站描述</Label>
                <Textarea 
                  id="site-description" 
                  defaultValue="基于 Go + React 的现代化内容管理系统"
                  rows={3}
                />
              </div>

              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">网站维护模式</div>
                    <div className="text-xs text-muted-foreground">
                      启用后，普通用户将无法访问网站
                    </div>
                  </div>
                  <Switch />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">用户注册</div>
                    <div className="text-xs text-muted-foreground">
                      允许新用户注册账号
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">评论功能</div>
                    <div className="text-xs text-muted-foreground">
                      启用文章评论功能
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>
              </div>

              <div className="flex justify-end">
                <Button>保存设置</Button>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="security" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Shield className="h-5 w-5" />
                安全配置
              </CardTitle>
              <CardDescription>
                配置系统的安全策略和访问控制
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="session-timeout">会话超时时间（分钟）</Label>
                  <Input id="session-timeout" type="number" defaultValue="30" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="max-login-attempts">最大登录尝试次数</Label>
                  <Input id="max-login-attempts" type="number" defaultValue="5" />
                </div>
              </div>

              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">强制 HTTPS</div>
                    <div className="text-xs text-muted-foreground">
                      强制所有连接使用 HTTPS 协议
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">双因素认证</div>
                    <div className="text-xs text-muted-foreground">
                      启用双因素认证功能
                    </div>
                  </div>
                  <Switch />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">IP 白名单</div>
                    <div className="text-xs text-muted-foreground">
                      启用管理员 IP 白名单限制
                    </div>
                  </div>
                  <Switch />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="allowed-ips">允许的 IP 地址</Label>
                <Textarea 
                  id="allowed-ips" 
                  placeholder="每行一个 IP 地址或 CIDR 块"
                  rows={4}
                />
              </div>

              <div className="flex justify-end">
                <Button>保存设置</Button>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="email" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Mail className="h-5 w-5" />
                邮件服务配置
              </CardTitle>
              <CardDescription>
                配置系统邮件发送服务
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="smtp-host">SMTP 服务器</Label>
                  <Input id="smtp-host" placeholder="smtp.example.com" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="smtp-port">端口</Label>
                  <Input id="smtp-port" type="number" defaultValue="587" />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="smtp-username">用户名</Label>
                  <Input id="smtp-username" placeholder="your-email@example.com" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="smtp-password">密码</Label>
                  <Input id="smtp-password" type="password" placeholder="••••••••" />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="from-email">发件人邮箱</Label>
                <Input id="from-email" placeholder="noreply@example.com" />
              </div>

              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">启用 TLS</div>
                    <div className="text-xs text-muted-foreground">
                      使用 TLS 加密连接
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">邮件通知</div>
                    <div className="text-xs text-muted-foreground">
                      启用系统邮件通知功能
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>
              </div>

              <div className="flex justify-between">
                <Button variant="outline">测试连接</Button>
                <Button>保存设置</Button>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="database" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Database className="h-5 w-5" />
                数据库信息
              </CardTitle>
              <CardDescription>
                查看数据库状态和统计信息
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label>数据库类型</Label>
                  <div className="text-sm text-muted-foreground">PostgreSQL</div>
                </div>
                <div className="space-y-2">
                  <Label>数据库版本</Label>
                  <div className="text-sm text-muted-foreground">14.9</div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label>连接状态</Label>
                  <div className="flex items-center space-x-2">
                    <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                    <span className="text-sm text-green-600">已连接</span>
                  </div>
                </div>
                <div className="space-y-2">
                  <Label>数据库大小</Label>
                  <div className="text-sm text-muted-foreground">25.6 MB</div>
                </div>
              </div>

              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">自动备份</div>
                    <div className="text-xs text-muted-foreground">
                      每日自动备份数据库
                    </div>
                  </div>
                  <Switch defaultChecked />
                </div>

                <div className="flex items-center justify-between">
                  <div className="space-y-0.5">
                    <div className="text-sm font-medium">查询日志</div>
                    <div className="text-xs text-muted-foreground">
                      记录数据库查询日志
                    </div>
                  </div>
                  <Switch />
                </div>
              </div>

              <div className="flex justify-between">
                <div className="space-x-2">
                  <Button variant="outline">立即备份</Button>
                  <Button variant="outline">优化数据库</Button>
                </div>
                <Button variant="outline" className="text-destructive">
                  <Database className="mr-2 h-4 w-4" />
                  重置数据库
                </Button>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      {/* 系统状态概览 */}
      <div className="grid gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              系统状态
            </CardTitle>
            <Server className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-green-600">正常</div>
            <p className="text-xs text-muted-foreground">
              所有服务运行正常
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              运行时间
            </CardTitle>
            <Settings className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">7天</div>
            <p className="text-xs text-muted-foreground">
              系统连续运行时间
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              内存使用
            </CardTitle>
            <Database className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">45%</div>
            <p className="text-xs text-muted-foreground">
              当前内存使用率
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              磁盘空间
            </CardTitle>
            <Shield className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">78%</div>
            <p className="text-xs text-muted-foreground">
              磁盘空间使用率
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
