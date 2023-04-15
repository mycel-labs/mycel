import { createBrowserRouter, Outlet } from "react-router-dom";

import IgntHeader from "../components/IgntHeader";
import DataView from "../views/ResolveView";
import PortfolioView from "../views/PortfolioView";
import ResolveView from "../views/ResolveView";
import SendView from "../views/SendView";
import ExploreView from "../views/ExploreView";

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
  {
    label: "Explore",
    to: "/explore"
  }
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
      { path: "/explore", element: <ExploreView /> }
    ],
  },
]);

export default router;
