"use client";
import React, { createContext, ReactNode } from "react";
import { Executor, ExecutorSupplier, NullExecutor, ServerExecutor } from "@srsim/executor";

type ExecutorContextProviderProps = {
  children: ReactNode;
};

type ExecutorContextType = {
  supplier: ExecutorSupplier<Executor>;
};

export const ExecutorContext = createContext<ExecutorContextType>({
  supplier: () => {
    return new NullExecutor();
  },
});

let exec: ServerExecutor | undefined;
const urlKey = "server-mode-url";
const defaultURL = "http://127.0.0.1:54321";

export const ExecutorProvider = ({ children }: ExecutorContextProviderProps) => {
  const [url, setURL] = React.useState<string>(defaultURL);
  React.useEffect(() => {
    if (typeof window !== "undefined") {
      const saved = window.localStorage.getItem(urlKey);
      if (saved === null) {
        window.localStorage.setItem(urlKey, defaultURL);
        return;
      }
      setURL(saved);
    }
  }, []);
  React.useEffect(() => {
    if (typeof window !== "undefined") {
      window.localStorage.setItem(urlKey, url);
    }
    if (exec != null) {
      exec.set_url(url);
    }
  }, [url]);
  const supplier = React.useRef<ExecutorSupplier<Executor>>(() => {
    if (exec == null) {
      exec = new ServerExecutor(url);
    }
    return exec;
  });
  return (
    <ExecutorContext.Provider value={{ supplier: supplier.current }}>
      {children}
    </ExecutorContext.Provider>
  );
};
