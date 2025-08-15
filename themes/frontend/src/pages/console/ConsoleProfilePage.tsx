/**
 * 个人中心页面
 */
import { useEffect } from "react";
import { Loader2 } from "lucide-react";

import {
  BasicInfoCard,
  SecurityCard,
} from "@/components/profile";

import { useUserStore } from "@/stores";
import { userService } from "@/services";

export default function ProfilePage() {
  const { user, setUser } = useUserStore();

  // 页面加载时获取最新用户数据
  useEffect(() => {
    const fetchLatestUserData = async () => {
      if (user?.id) {
        try {
          const latestUser = await userService.getProfile();
          setUser(latestUser);
        } catch (error) {
          console.error("获取用户数据失败:", error);
        }
      }
    };

    fetchLatestUserData();
  }, [user?.id, setUser]);

  if (!user) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <Loader2 className="h-8 w-8 animate-spin" />
      </div>
    );
  }

  return (
    <div className="p-6 space-y-6">
      <BasicInfoCard user={user} />
      <SecurityCard />
    </div>
  );
}
