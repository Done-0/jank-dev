import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { FileText, Search, Plus, Edit, Trash2, Eye } from "lucide-react";

export function ConsolePostsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">文章管理</h1>
          <p className="text-muted-foreground">
            管理系统中的所有文章内容
          </p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          新建文章
        </Button>
      </div>

      <Tabs defaultValue="all" className="space-y-4">
        <TabsList>
          <TabsTrigger value="all">全部文章</TabsTrigger>
          <TabsTrigger value="published">已发布</TabsTrigger>
          <TabsTrigger value="draft">草稿</TabsTrigger>
          <TabsTrigger value="archived">已归档</TabsTrigger>
        </TabsList>

        <TabsContent value="all" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>文章列表</CardTitle>
              <CardDescription>
                系统中所有文章的管理界面
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex items-center space-x-2 mb-4">
                <div className="relative flex-1 max-w-sm">
                  <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="搜索文章..."
                    className="pl-8"
                  />
                </div>
                <Select>
                  <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="选择分类" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">全部分类</SelectItem>
                    <SelectItem value="tech">技术</SelectItem>
                    <SelectItem value="life">生活</SelectItem>
                    <SelectItem value="travel">旅行</SelectItem>
                  </SelectContent>
                </Select>
                <Select>
                  <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="排序方式" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="latest">最新发布</SelectItem>
                    <SelectItem value="oldest">最早发布</SelectItem>
                    <SelectItem value="views">浏览量</SelectItem>
                    <SelectItem value="title">标题</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div className="rounded-md border">
                <div className="p-4">
                  <div className="grid grid-cols-6 gap-4 font-medium text-sm text-muted-foreground mb-4">
                    <div className="col-span-2">标题</div>
                    <div>分类</div>
                    <div>状态</div>
                    <div>发布时间</div>
                    <div>操作</div>
                  </div>
                  
                  <div className="space-y-4">
                    {/* 示例文章数据 */}
                    <div className="grid grid-cols-6 gap-4 items-center py-2 border-b">
                      <div className="col-span-2">
                        <div className="flex items-center space-x-3">
                          <div className="w-8 h-8 bg-blue-100 rounded flex items-center justify-center">
                            <FileText className="h-4 w-4 text-blue-600" />
                          </div>
                          <div>
                            <p className="font-medium">React 18 新特性详解</p>
                            <p className="text-sm text-muted-foreground">深入了解 React 18 的并发特性...</p>
                          </div>
                        </div>
                      </div>
                      <div>
                        <Badge variant="outline">技术</Badge>
                      </div>
                      <div>
                        <Badge variant="default">已发布</Badge>
                      </div>
                      <div className="text-sm text-muted-foreground">
                        2024-01-15
                      </div>
                      <div className="flex items-center space-x-1">
                        <Button variant="ghost" size="sm">
                          <Eye className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm">
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm" className="text-destructive">
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>

                    <div className="grid grid-cols-6 gap-4 items-center py-2 border-b">
                      <div className="col-span-2">
                        <div className="flex items-center space-x-3">
                          <div className="w-8 h-8 bg-green-100 rounded flex items-center justify-center">
                            <FileText className="h-4 w-4 text-green-600" />
                          </div>
                          <div>
                            <p className="font-medium">TypeScript 最佳实践</p>
                            <p className="text-sm text-muted-foreground">分享 TypeScript 开发经验...</p>
                          </div>
                        </div>
                      </div>
                      <div>
                        <Badge variant="outline">技术</Badge>
                      </div>
                      <div>
                        <Badge variant="secondary">草稿</Badge>
                      </div>
                      <div className="text-sm text-muted-foreground">
                        2024-01-10
                      </div>
                      <div className="flex items-center space-x-1">
                        <Button variant="ghost" size="sm">
                          <Eye className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm">
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm" className="text-destructive">
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>

                    <div className="grid grid-cols-6 gap-4 items-center py-2 border-b">
                      <div className="col-span-2">
                        <div className="flex items-center space-x-3">
                          <div className="w-8 h-8 bg-purple-100 rounded flex items-center justify-center">
                            <FileText className="h-4 w-4 text-purple-600" />
                          </div>
                          <div>
                            <p className="font-medium">日本旅行攻略</p>
                            <p className="text-sm text-muted-foreground">分享日本旅行的经验和建议...</p>
                          </div>
                        </div>
                      </div>
                      <div>
                        <Badge variant="outline">旅行</Badge>
                      </div>
                      <div>
                        <Badge variant="default">已发布</Badge>
                      </div>
                      <div className="text-sm text-muted-foreground">
                        2024-01-05
                      </div>
                      <div className="flex items-center space-x-1">
                        <Button variant="ghost" size="sm">
                          <Eye className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm">
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="sm" className="text-destructive">
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div className="flex items-center justify-between mt-4">
                <p className="text-sm text-muted-foreground">
                  显示 1-3 条，共 3 条记录
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
        </TabsContent>

        <TabsContent value="published" className="space-y-4">
          <Card>
            <CardContent className="pt-6">
              <div className="text-center py-8 text-muted-foreground">
                <FileText className="h-12 w-12 mx-auto mb-4 opacity-50" />
                <p>已发布文章列表</p>
                <p className="text-sm">显示所有已发布的文章</p>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="draft" className="space-y-4">
          <Card>
            <CardContent className="pt-6">
              <div className="text-center py-8 text-muted-foreground">
                <FileText className="h-12 w-12 mx-auto mb-4 opacity-50" />
                <p>草稿文章列表</p>
                <p className="text-sm">显示所有草稿状态的文章</p>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="archived" className="space-y-4">
          <Card>
            <CardContent className="pt-6">
              <div className="text-center py-8 text-muted-foreground">
                <FileText className="h-12 w-12 mx-auto mb-4 opacity-50" />
                <p>已归档文章列表</p>
                <p className="text-sm">显示所有已归档的文章</p>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      {/* 文章统计 */}
      <div className="grid gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              总文章数
            </CardTitle>
            <FileText className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">3</div>
            <p className="text-xs text-muted-foreground">
              系统中所有文章数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              已发布
            </CardTitle>
            <Eye className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">2</div>
            <p className="text-xs text-muted-foreground">
              已发布的文章数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              草稿
            </CardTitle>
            <Edit className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">1</div>
            <p className="text-xs text-muted-foreground">
              草稿状态的文章
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              总浏览量
            </CardTitle>
            <Plus className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">1.2k</div>
            <p className="text-xs text-muted-foreground">
              所有文章的浏览量
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
