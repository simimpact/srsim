import { useRoutes } from "react-router-dom";
import { EditorRoutes } from "../features/editor";
import { ViewerRoutes } from "../features/viewer";
import { Debug } from "./Debug/page";
import { Root } from "./Root/page";

export const AppRoutes = () => {
  const routes = useRoutes([
    { path: "/", element: <Root /> },
    { path: "/debug", element: <Debug /> },
    ...EditorRoutes,
    ...ViewerRoutes,
    { path: "*", element: <></> },
  ]);

  return <>{routes}</>;
};
