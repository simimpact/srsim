import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { AppProvider } from "./providers/app";
import "./index.css";

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion, @typescript-eslint/no-unsafe-call, @typescript-eslint/no-unsafe-member-access
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <AppProvider>
      <App />
    </AppProvider>
  </React.StrictMode>
);
