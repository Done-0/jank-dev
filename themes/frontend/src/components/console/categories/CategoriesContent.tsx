import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Badge } from '@/components/ui/badge';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from '@/components/ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Search, Plus, Edit, Trash2, Eye, EyeOff, MoreHorizontal, FolderOpen } from 'lucide-react';
import { useCreateCategory, useUpdateCategory, useDeleteCategory } from '@/hooks/use-categories';
import { ConfirmDialog } from '@/components/ui/confirm-dialog';
import type { CategoryItem } from '@/types/category';

interface CategoriesContentProps {
  categories: CategoryItem[];
  isLoading: boolean;
}

export function CategoriesContent({ categories, isLoading }: CategoriesContentProps) {
  // ===== State =====
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [editDialogOpen, setEditDialogOpen] = useState(false);
  const [editingCategory, setEditingCategory] = useState<CategoryItem | null>(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [formData, setFormData] = useState({ 
    name: '', 
    description: '', 
    parent_id: 'none', 
    sort: 0 
  });

  // ===== Hooks =====
  const createCategoryMutation = useCreateCategory();
  const updateCategoryMutation = useUpdateCategory();
  const deleteCategoryMutation = useDeleteCategory();

  // 获取分类颜色
  const getCategoryColor = (categoryId: string) => {
    const colors = ['bg-blue-500', 'bg-green-500', 'bg-yellow-500', 'bg-red-500', 'bg-purple-500', 'bg-pink-500', 'bg-indigo-500', 'bg-orange-500'];
    return colors[parseInt(categoryId) % colors.length];
  };

  // ===== 事件处理 =====
  const handleNewCategory = () => {
    setFormData({ name: '', description: '', parent_id: 'none', sort: 0 });
    setCreateDialogOpen(true);
  };

  const handleCreateSubmit = async () => {
    if (!formData.name.trim()) return;
    
    try {
      await createCategoryMutation.mutateAsync({
        name: formData.name.trim(),
        description: formData.description.trim(),
        parent_id: formData.parent_id === 'none' ? '0' : formData.parent_id, // "0"表示顶级分类
        sort: formData.sort,
        is_active: true
      });
      setCreateDialogOpen(false);
      setFormData({ name: '', description: '', parent_id: 'none', sort: 0 });
    } catch (error) {
      console.error('创建分类失败:', error);
    }
  };

  const handleEditCategory = (category: CategoryItem) => {
    setEditingCategory(category);
    setFormData({ 
      name: category.name, 
      description: category.description || '',
      parent_id: category.parent_id === '0' ? 'none' : category.parent_id,
      sort: category.sort || 0
    });
    setEditDialogOpen(true);
  };

  const handleEditSubmit = async () => {
    if (!editingCategory || !formData.name.trim()) return;
    
    try {
      await updateCategoryMutation.mutateAsync({
        id: editingCategory.id,
        name: formData.name.trim(),
        description: formData.description.trim(),
        parent_id: formData.parent_id === 'none' ? '0' : formData.parent_id, // "0"表示顶级分类 
        sort: formData.sort
      });
      setEditDialogOpen(false);
      setEditingCategory(null);
      setFormData({ name: '', description: '', parent_id: 'none', sort: 0 });
    } catch (error) {
      console.error('更新分类失败:', error);
    }
  };

  const handleToggleActive = async (category: CategoryItem) => {
    try {
      await updateCategoryMutation.mutateAsync({
        id: category.id,
        is_active: !category.is_active
      });
    } catch (error) {
      console.error('切换分类状态失败:', error);
    }
  };

  const handleDeleteCategory = async (category: CategoryItem) => {
    try {
      await deleteCategoryMutation.mutateAsync({ id: category.id });
    } catch (error) {
      console.error('删除分类失败:', error);
    }
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center py-12">
        <div className="text-muted-foreground">加载中...</div>
      </div>
    );
  }

  // 过滤分类
  const filteredCategories = categories.filter(category =>
    category.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
    category.description?.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="flex-1 flex flex-col h-full">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索分类..."
              className="pl-9 h-10 rounded-full"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
          </div>
          <Button onClick={handleNewCategory} className="h-10 px-4 shrink-0 rounded-full">
            <Plus className="mr-2 h-4 w-4" />
            <span className="hidden sm:inline">新建分类</span>
            <span className="sm:hidden">新建</span>
          </Button>
        </div>
      </div>

      {/* 分类列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {filteredCategories.length === 0 ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              {searchQuery ? (
                <>
                  <Search className="w-12 h-12 mx-auto text-muted-foreground mb-4" />
                  <h3 className="text-lg font-medium mb-2">未找到匹配的分类</h3>
                  <p className="text-muted-foreground mb-4">尝试使用其他关键词搜索</p>
                </>
              ) : (
                <>
                  <FolderOpen className="w-12 h-12 mx-auto text-muted-foreground mb-4" />
                  <h3 className="text-lg font-medium mb-2">暂无分类</h3>
                  <p className="text-muted-foreground mb-4">创建第一个分类来开始管理内容</p>
                  <Button onClick={handleNewCategory} className="rounded-full">
                    <Plus className="w-4 h-4 mr-2" />
                    新建分类
                  </Button>
                </>
              )}
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-auto">
            <div className="divide-y divide-border">
              {filteredCategories.map((category) => (
            <div key={category.id} className="px-4 py-4 hover:bg-accent/50 transition-colors">
              <div className="flex flex-col gap-3">
                <div className="flex items-start justify-between gap-3">
                  <div className="flex-1 min-w-0">
                    <h3 className="font-medium text-lg line-clamp-2">{category.name}</h3>
                  </div>
                  <Badge variant={category.is_active ? "default" : "secondary"} className="shrink-0">
                    {category.is_active ? '激活' : '禁用'}
                  </Badge>
                </div>
                
                <div className="min-h-[1.25rem]">
                  <p className="text-sm text-muted-foreground line-clamp-2">
                    {category.description || '暂无描述'}
                  </p>
                </div>
                
                <div className="flex items-center justify-between gap-3">
                  <div className="flex items-center gap-2 text-xs text-muted-foreground">
                    <div className="flex items-center gap-1.5">
                      <div className={`w-1.5 h-1.5 rounded-full ${getCategoryColor(category.id)}`} />
                      <span>排序: {category.sort || 0}</span>
                    </div>
                    <div className="flex items-center gap-1.5">
                      <div className="w-1.5 h-1.5 rounded-full bg-slate-400" />
                      <span>{new Date(category.updated_at).toLocaleDateString()}</span>
                    </div>
                  </div>
                  
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="sm" className="h-8 w-8 p-0 rounded-full">
                        <MoreHorizontal className="h-4 w-4" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end" className="w-36 rounded-xl">
                      <DropdownMenuItem onClick={() => handleEditCategory(category)} className="py-2.5">
                        <Edit className="mr-2 h-4 w-4" />
                        编辑分类
                      </DropdownMenuItem>
                      <DropdownMenuItem onClick={() => handleToggleActive(category)} className="py-2.5">
                        {category.is_active ? (
                          <>
                            <EyeOff className="mr-2 h-4 w-4" />
                            禁用分类
                          </>
                        ) : (
                          <>
                            <Eye className="mr-2 h-4 w-4" />
                            启用分类
                          </>
                        )}
                      </DropdownMenuItem>
                      <DropdownMenuSeparator />
                      <ConfirmDialog
                        title="删除分类"
                        description={`确定要删除分类"${category.name}"吗？此操作不可撤销。`}
                        onConfirm={() => handleDeleteCategory(category)}
                        destructive
                      >
                        <DropdownMenuItem 
                          className="text-destructive focus:text-destructive py-2.5"
                          onSelect={(e) => e.preventDefault()}
                        >
                          <Trash2 className="mr-2 h-4 w-4" />
                          删除分类
                        </DropdownMenuItem>
                      </ConfirmDialog>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              </div>
            </div>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* 创建分类弹窗 */}
      <Dialog open={createDialogOpen} onOpenChange={setCreateDialogOpen}>
        <DialogContent className="sm:max-w-lg rounded-2xl">
          <DialogHeader className="text-center">
            <DialogTitle className="text-xl font-semibold">创建新分类</DialogTitle>
            <p className="text-sm text-muted-foreground mt-2">为你的内容添加一个新的分类标签</p>
          </DialogHeader>
          
          <div className="space-y-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">分类名称</label>
              <Input
                value={formData.name}
                onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
                placeholder="输入分类名称"
                className="h-10 rounded-xl"
                autoFocus
              />
            </div>
            
            <div className="space-y-2">
              <label className="text-sm font-medium">分类描述</label>
              <Textarea
                value={formData.description}
                onChange={(e) => setFormData(prev => ({ ...prev, description: e.target.value }))}
                placeholder="简要描述这个分类的用途（可选）"
                className="rounded-xl resize-none"
                rows={3}
              />
            </div>
            
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">父分类</label>
                <Select value={formData.parent_id} onValueChange={(value) => setFormData(prev => ({ ...prev, parent_id: value }))}>
                  <SelectTrigger className="h-10 rounded-xl">
                    <SelectValue placeholder="选择父分类" />
                  </SelectTrigger>
                  <SelectContent className="rounded-xl">
                    <SelectItem value="none">
                      顶级分类
                    </SelectItem>
                    {categories.map(category => (
                      <SelectItem key={category.id} value={category.id}>
                        {category.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              
              <div className="space-y-2">
                <label className="text-sm font-medium">排序权重</label>
                <Input
                  type="number"
                  value={formData.sort}
                  onChange={(e) => setFormData(prev => ({ ...prev, sort: parseInt(e.target.value) || 0 }))}
                  placeholder="0"
                  className="h-10 rounded-xl [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                  min="0"
                />
              </div>
            </div>
          </div>
          
          <DialogFooter className="gap-3">
            <Button variant="outline" onClick={() => setCreateDialogOpen(false)} className="rounded-xl">
              取消
            </Button>
            <Button 
              onClick={handleCreateSubmit}
              disabled={!formData.name.trim() || createCategoryMutation.isPending}
              className="rounded-xl"
            >
              {createCategoryMutation.isPending ? "创建中..." : "创建分类"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 编辑分类弹窗 */}
      <Dialog open={editDialogOpen} onOpenChange={setEditDialogOpen}>
        <DialogContent className="sm:max-w-lg rounded-2xl">
          <DialogHeader className="text-center">
            <DialogTitle className="text-xl font-semibold">编辑分类</DialogTitle>
            <p className="text-sm text-muted-foreground mt-2">修改 &ldquo;{editingCategory?.name}&rdquo; 的信息</p>
          </DialogHeader>
          
          <div className="space-y-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">分类名称</label>
              <Input
                value={formData.name}
                onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
                placeholder="输入分类名称"
                className="h-10 rounded-xl"
                autoFocus
              />
            </div>
            
            <div className="space-y-2">
              <label className="text-sm font-medium">分类描述</label>
              <Textarea
                value={formData.description}
                onChange={(e) => setFormData(prev => ({ ...prev, description: e.target.value }))}
                placeholder="简要描述这个分类的用途（可选）"
                className="rounded-xl resize-none"
                rows={3}
              />
            </div>
            
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">父分类</label>
                <Select value={formData.parent_id} onValueChange={(value) => setFormData(prev => ({ ...prev, parent_id: value }))}>
                  <SelectTrigger className="h-10 rounded-xl">
                    <SelectValue placeholder="选择父分类" />
                  </SelectTrigger>
                  <SelectContent className="rounded-xl">
                    <SelectItem value="none">
                      顶级分类
                    </SelectItem>
                    {categories.filter(cat => cat.id !== editingCategory?.id).map(category => (
                      <SelectItem key={category.id} value={category.id}>
                        {category.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              
              <div className="space-y-2">
                <label className="text-sm font-medium">排序权重</label>
                <Input
                  type="number"
                  value={formData.sort}
                  onChange={(e) => setFormData(prev => ({ ...prev, sort: parseInt(e.target.value) || 0 }))}
                  placeholder="0"
                  className="h-10 rounded-xl [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                  min="0"
                />
              </div>
            </div>
          </div>
          
          <DialogFooter className="gap-3">
            <Button 
              variant="outline" 
              onClick={() => {
                setEditDialogOpen(false);
                setEditingCategory(null);
                setFormData({ name: '', description: '', parent_id: 'none', sort: 0 });
              }}
              className="rounded-xl"
            >
              取消
            </Button>
            <Button 
              onClick={handleEditSubmit}
              disabled={!formData.name.trim() || updateCategoryMutation.isPending}
              className="rounded-xl"
            >
              {updateCategoryMutation.isPending ? "更新中..." : "保存更改"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
