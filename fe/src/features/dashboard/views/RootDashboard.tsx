import { SidebarProvider } from "@/components/ui/sidebar";
import { Outlet } from "react-router-dom";
import { SidebarDashboard } from "../components";

export const RootDashboard = () => {
  return (
    <SidebarProvider>
      <SidebarDashboard title="Overview">
        <Outlet />
      </SidebarDashboard>
    </SidebarProvider>
  );
};
