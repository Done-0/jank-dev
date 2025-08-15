import { Home } from "lucide-react";
import { SidebarTrigger } from "@/components/ui/sidebar";
import { Separator } from "@/components/ui/separator";
import { ROUTE_TITLES } from "@/constants";

interface AppHeaderProps {
  activeTab: string;
}

export function AppHeader({ activeTab }: AppHeaderProps) {
  const title = ROUTE_TITLES[activeTab as keyof typeof ROUTE_TITLES] || "控制台";

  return (
    <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
      <div className="flex items-center gap-2 px-4">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
        <div className="flex items-center gap-2 text-sm font-medium">
          <Home className="h-4 w-4" />
          <span>{title}</span>
        </div>
      </div>
    </header>
  );
}
