/**
 * 文章管理页面
 */

import { useState } from "react";

import { PostsSidebar } from "@/components/console/posts/PostsSidebar";
import { PostsContent } from "@/components/console/posts/PostsContent";

import { usePostsByStatus } from "@/hooks/use-posts";
import { useAllCategories } from "@/hooks/use-categories";
import type { PostStatus } from "@/constants/post";

export function ConsolePostsPage() {
  // ===== 状态管理 =====
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null);
  const [selectedStatus, setSelectedStatus] = useState<PostStatus | null>(null);
  const [searchQuery, setSearchQuery] = useState("");

  // ===== 数据获取 =====
  const requestParams = selectedStatus 
    ? { page_no: 1, page_size: 100, status: selectedStatus }
    : { page_no: 1, page_size: 100 };
  
  const { data: postsData, isLoading } = usePostsByStatus(requestParams);
  const { data: allPostsData } = usePostsByStatus({ page_no: 1, page_size: 100 });
  const { data: categories = [] } = useAllCategories();
  
  // ===== 计算数据 =====
  const allPosts = allPostsData?.list || [];
  const currentPosts = selectedStatus ? (postsData?.list || []) : allPosts;
  const stats = {
    total: allPosts.length,
    published: allPosts.filter(post => post.status === 'published').length,
    draft: allPosts.filter(post => post.status === 'draft').length,
    archived: allPosts.filter(post => post.status === 'archived').length,
  };

  // ===== 事件处理 =====
  const handleStatusChange = (status: PostStatus | null) => setSelectedStatus(status);
  const handleCategoryChange = (category: string | null) => setSelectedCategory(category);
  const handleSearchChange = (query: string) => setSearchQuery(query);
  const handleNewPost = () => console.log('新建文章'); // TODO: 实现新建文章逻辑

  // ===== 渲染 =====
  return (
    <div className="flex h-full">
      <div className="hidden md:block">
        <PostsSidebar
          selectedStatus={selectedStatus}
          selectedCategory={selectedCategory}
          stats={stats}
          categories={categories}
          onStatusChange={handleStatusChange}
          onCategoryChange={handleCategoryChange}
        />
      </div>

      <div className="flex-1 flex flex-col h-full">
        <PostsContent
          selectedStatus={selectedStatus}
          selectedCategory={selectedCategory}
          searchQuery={searchQuery}
          categories={categories}
          allPosts={currentPosts}
          isLoading={isLoading}
          onSearchChange={handleSearchChange}
          onNewPost={handleNewPost}
          onStatusChange={handleStatusChange}
          onCategoryChange={handleCategoryChange}
        />
      </div>
    </div>
  );
}
