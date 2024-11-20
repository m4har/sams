import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { Calendar, Cat, Home, Inbox, Search, Settings } from "lucide-react";

// Menu items.
const items = [
  {
    title: "Home",
    url: "#",
    icon: Home,
  },
  {
    title: "Inbox",
    url: "#",
    icon: Inbox,
  },
  {
    title: "Calendar",
    url: "#",
    icon: Calendar,
  },
  {
    title: "Search",
    url: "#",
    icon: Search,
  },
  {
    title: "Settings",
    url: "#",
    icon: Settings,
  },
];

type Props = {
  children: React.ReactNode;
  title: string;
};

export const SidebarDashboard = ({ children, title }: Props) => {
  return (
    <div className="flex w-full">
      <Sidebar>
        <SidebarContent>
          <SidebarGroup>
            <SidebarHeader>
              <div className="flex items-center space-x-2">
                <Cat className="text-mainAccent" size={40} />
                <span className="font-semibold text-lg">Execu AI</span>
              </div>
            </SidebarHeader>
            <SidebarGroupContent>
              <SidebarMenu>
                {items.map((item) => (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton asChild>
                      <a href={item.url}>
                        <item.icon />
                        <span>{item.title}</span>
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>
      </Sidebar>

      <main className="w-full relative">
        <nav className="sticky top-0 z-50 mx-auto flex h-[88px] w-full items-center border-b border-border bg-white pr-5 m500:h-16 ">
          <div className="flex flex-1 items-center">
            <SidebarTrigger className="mx-3" />
            <h2 className="text-xl font-bold tracking-tight capitalize">
              {title}
            </h2>
            <div className="ml-auto flex items-center space-x-4">
              {/* use navigation profile */}
            </div>
          </div>
        </nav>
        <div className="mx-3">{children}</div>
      </main>
    </div>
  );
};
