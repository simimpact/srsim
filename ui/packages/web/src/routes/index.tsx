import { Route, Routes } from "react-router-dom";
import { EditorRoutes } from "../features/editor";
import { ViewerRoutes } from "../features/viewer";

export const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<></>} />
      <EditorRoutes />
      <ViewerRoutes />
      <Route path="*" element={<></>} />
    </Routes>
  );
};
