import { useRoutes } from "react-router-dom";
import { EditorRoutes } from "../features/editor";
import { ViewerRoutes } from "../features/viewer";

export const AppRoutes = () => {
  const routes = useRoutes([
    { path: "/", element: <></> },
    ...EditorRoutes,
    ...ViewerRoutes,
    { path: "*", element: <></> },
  ]);

  return <>{routes}</>;
};
