import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { FolderOpen, Search, Plus, Edit, Trash2, FileText } from "lucide-react";

export function ConsoleCategoriesPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">分类管理</h1>
          <p className="text-muted-foreground">
            管理文章分类和标签系统
          </p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          新建分类
        </Button>
      </div>

      <div className="grid gap-6 md:grid-cols-3">
        {/* 分类列表 */}
        <div className="md:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>分类列表</CardTitle>
              <CardDescription>
                系统中所有文章分类的管理
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex items-center space-x-2 mb-4">
                <div className="relative flex-1 max-w-sm">
                  <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="搜索分类..."
                    className="pl-8"
                  />
                </div>
              </div>

              <div className="space-y-4">
                {/* 示例分类数据 */}
                <div className="flex items-center justify-between p-4 border rounded-lg">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                      <FolderOpen className="h-5 w-5 text-blue-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">技术</h3>
                      <p className="text-sm text-muted-foreground">编程、开发相关内容</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="outline">5 篇文章</Badge>
                    <Button variant="outline" size="sm">
                      <Edit className="mr-2 h-4 w-4" />
                      编辑
                    </Button>
                    <Button variant="outline" size="sm" className="text-destructive">
                      <Trash2 className="mr-2 h-4 w-4" />
                      删除
                    </Button>
                  </div>
                </div>

                <div className="flex items-center justify-between p-4 border rounded-lg">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                      <FolderOpen className="h-5 w-5 text-green-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">生活</h3>
                      <p className="text-sm text-muted-foreground">日常生活、感悟分享</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="outline">3 篇文章</Badge>
                    <Button variant="outline" size="sm">
                      <Edit className="mr-2 h-4 w-4" />
                      编辑
                    </Button>
                    <Button variant="outline" size="sm" className="text-destructive">
                      <Trash2 className="mr-2 h-4 w-4" />
                      删除
                    </Button>
                  </div>
                </div>

                <div className="flex items-center justify-between p-4 border rounded-lg">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
                      <FolderOpen className="h-5 w-5 text-purple-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">旅行</h3>
                      <p className="text-sm text-muted-foreground">旅行攻略、游记分享</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="outline">2 篇文章</Badge>
                    <Button variant="outline" size="sm">
                      <Edit className="mr-2 h-4 w-4" />
                      编辑
                    </Button>
                    <Button variant="outline" size="sm" className="text-destructive">
                      <Trash2 className="mr-2 h-4 w-4" />
                      删除
                    </Button>
                  </div>
                </div>

                <div className="flex items-center justify-between p-4 border rounded-lg opacity-60">
                  <div className="flex items-center space-x-4">
                    <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center">
                      <FolderOpen className="h-5 w-5 text-gray-600" />
                    </div>
                    <div>
                      <h3 className="font-medium">未分类</h3>
                      <p className="text-sm text-muted-foreground">暂未归类的文章</p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <Badge variant="secondary">0 篇文章</Badge>
                    <Button variant="outline" size="sm" disabled>
                      <Edit className="mr-2 h-4 w-4" />
                      编辑
                    </Button>
                    <Button variant="outline" size="sm" disabled>
                      <Trash2 className="mr-2 h-4 w-4" />
                      删除
                    </Button>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* 分类统计 */}
        <div className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">分类统计</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center justify-between">
                <span className="text-sm text-muted-foreground">总分类数</span>
                <span className="font-medium">4</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-muted-foreground">有文章分类</span>
                <span className="font-medium">3</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-muted-foreground">空分类</span>
                <span className="font-medium">1</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-muted-foreground">总文章数</span>
                <span className="font-medium">10</span>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle className="text-lg">热门分类</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-2">
                  <div className="w-3 h-3 bg-blue-500 rounded-full"></div>
                  <span className="text-sm">技术</span>
                </div>
                <span className="text-sm font-medium">5 篇</span>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-2">
                  <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                  <span className="text-sm">生活</span>
                </div>
                <span className="text-sm font-medium">3 篇</span>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-2">
                  <div className="w-3 h-3 bg-purple-500 rounded-full"></div>
                  <span className="text-sm">旅行</span>
                </div>
                <span className="text-sm font-medium">2 篇</span>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle className="text-lg">快速操作</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              <Button variant="outline" className="w-full justify-start">
                <Plus className="mr-2 h-4 w-4" />
                新建分类
              </Button>
              <Button variant="outline" className="w-full justify-start">
                <FileText className="mr-2 h-4 w-4" />
                批量管理
              </Button>
              <Button variant="outline" className="w-full justify-start">
                <FolderOpen className="mr-2 h-4 w-4" />
                分类排序
              </Button>
            </CardContent>
          </Card>
        </div>
      </div>

      {/* 分类概览卡片 */}
      <div className="grid gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              总分类数
            </CardTitle>
            <FolderOpen className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">4</div>
            <p className="text-xs text-muted-foreground">
              系统中所有分类数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              活跃分类
            </CardTitle>
            <FileText className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">3</div>
            <p className="text-xs text-muted-foreground">
              包含文章的分类数量
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              最热分类
            </CardTitle>
            <Plus className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">技术</div>
            <p className="text-xs text-muted-foreground">
              文章数量最多的分类
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              平均文章数
            </CardTitle>
            <Edit className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">3.3</div>
            <p className="text-xs text-muted-foreground">
              每个分类的平均文章数
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
