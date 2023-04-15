import { createBrowserRouter, Outlet } from "react-router-dom";

import IgntHeader from "../components/IgntHeader";
import DataView from "../views/ResolveView";
import PortfolioView from "../views/PortfolioView";
import ResolveView from "../views/ResolveView";

const items = [
  {
    label: "Portfolio",
    to: "/",
  },
  {
    label: "Resolve",
    to: "/resolve",
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
    ],
  },
]);

export default router;
