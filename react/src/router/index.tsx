import { createBrowserRouter, Outlet } from "react-router-dom";

import IgntHeader from "../components/IgntHeader";
import DataView from "../views/ResolveView";
import PortfolioView from "../views/PortfolioView";
import ResolveView from "../views/ResolveView";
import SendView from "../views/SendView";

const items = [
  {
    label: "Portfolio",
    to: "/",
  },
  {
    label: "Resolve Domain",
    to: "/resolve",
  },
  {
    label: "Send Transaction",
    to: "/send",
  },
];
const Layout = () => {
  return (
    <>
      <IgntHeader navItems={items} />
      <Outlet />
    </>
  );
};
const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "/", element: <PortfolioView /> },
      { path: "/resolve", element: <ResolveView /> },
      { path: "/send", element: <SendView /> },
    ],
  },
]);

export default router;
