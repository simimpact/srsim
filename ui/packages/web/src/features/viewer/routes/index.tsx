import { Route, Routes } from "react-router-dom";

export const ViewerRoutes = () => {
  return (
    <Routes>
      <Route path="web" element={<></>} />
      <Route path="local" element={<></>} />
      <Route path="sh/:id" element={<></>} />
    </Routes>
  );
};
