/**
 * 文章管理侧边栏组件
 */

import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { FileText, Edit, Eye, Archive, Hash } from "lucide-react";

import { CATEGORY_COLORS } from "@/constants";
import type { PostStatus } from "@/constants/post";
import type { CategoryItem } from "@/types/category";

interface PostsSidebarProps {
  selectedStatus: PostStatus | null;
  selectedCategory: string | null;
  stats: {
    total: number;
    published: number;
    draft: number;
    archived: number;
  };
  categories: CategoryItem[];
  onStatusChange: (status: PostStatus | null) => void;
  onCategoryChange: (category: string | null) => void;
  className?: string;
}

export function PostsSidebar({
  selectedStatus,
  selectedCategory,
  stats,
  categories,
  onStatusChange,
  onCategoryChange,
}: PostsSidebarProps) {

  // 获取分类颜色
  const getCategoryColor = (index: number) => {
    return CATEGORY_COLORS[index % CATEGORY_COLORS.length];
  };

  return (
    <div className="w-64 lg:w-72 border-r border-border bg-background hidden md:flex">
      <div className="h-full flex flex-col w-full">
        {/* 状态筛选区域 */}
        <div className="px-3 sm:px-4 py-3 sm:py-4 border-b border-border">
          <div className="space-y-1">
            <Button
              variant={selectedStatus === null ? "default" : "ghost"}
              className="w-full h-8 justify-start font-normal px-3 text-sm"
              onClick={() => onStatusChange(null)}
            >
              <FileText className="h-4 w-4 mr-2" />
              <span className="flex-1 text-left">全部</span>
              <Badge variant="secondary" className="h-5 px-1.5 text-xs ml-auto">
                {stats.total}
              </Badge>
            </Button>

            <Button
              variant={selectedStatus === 'published' ? "default" : "ghost"}
              className="w-full h-8 justify-start font-normal px-3 text-sm"
              onClick={() => onStatusChange('published')}
            >
              <Eye className="h-4 w-4 mr-2" />
              <span className="flex-1 text-left">已发布</span>
              <Badge variant="secondary" className="h-5 px-1.5 text-xs ml-auto">
                {stats.published}
              </Badge>
            </Button>

            <Button
              variant={selectedStatus === 'draft' ? "default" : "ghost"}
              className="w-full h-8 justify-start font-normal px-3 text-sm"
              onClick={() => onStatusChange('draft')}
            >
              <Edit className="h-4 w-4 mr-2" />
              <span className="flex-1 text-left">草稿</span>
              <Badge variant="secondary" className="h-5 px-1.5 text-xs ml-auto">
                {stats.draft}
              </Badge>
            </Button>

            <Button
              variant={selectedStatus === 'archived' ? "default" : "ghost"}
              className="w-full h-8 justify-start font-normal px-3 text-sm"
              onClick={() => onStatusChange('archived')}
            >
              <Archive className="h-4 w-4 mr-2" />
              <span className="flex-1 text-left">已归档</span>
              <Badge variant="secondary" className="h-5 px-1.5 text-xs ml-auto">
                {stats.archived}
              </Badge>
            </Button>
          </div>
        </div>

        {/* 分类筛选区域 */}
        <div className="flex-1 flex flex-col min-h-0">
          <div className="flex-1 overflow-y-auto scrollbar-hover min-h-0">
            <div className="space-y-1 px-3 sm:px-4 py-3 sm:py-4">
              <Button
                variant={selectedCategory === null ? "default" : "ghost"}
                className="w-full h-8 justify-start font-normal px-3 text-sm"
                onClick={() => onCategoryChange(null)}
              >
                <Hash className="h-4 w-4 mr-2" />
                <span className="flex-1 text-left">全部分类</span>
              </Button>
              
              {categories.length > 0 ? (
                categories
                  .filter(category => category.is_active)
                  .map((category, index) => (
                    <Button
                      key={category.id}
                      variant={selectedCategory === category.id ? "default" : "ghost"}
                      className="w-full h-8 justify-start font-normal px-3 text-sm"
                      onClick={() => onCategoryChange(category.id)}
                    >
                      <div className={`w-2 h-2 rounded-full ${getCategoryColor(index)} mr-2 flex-shrink-0`} />
                      <span className="flex-1 text-left truncate">{category.name}</span>
                    </Button>
                  ))
              ) : (
                <div className="text-sm text-muted-foreground px-3 py-2">暂无分类</div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
