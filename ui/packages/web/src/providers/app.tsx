import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactNode } from "react";
import { ErrorBoundary } from "react-error-boundary";
import { HelmetProvider } from "react-helmet-async";
import { BrowserRouter } from "react-router-dom";

interface Props {
  children: ReactNode;
}

const ErrorFallback = () => (
  <div className="flex flex-col justify-center items-center w-screen h-screen gap-5">
    <h2 className="text-xl">something went wrong :(</h2>
    <button className="rounded bg-blue-500 font-blod px-4 py-2 text-white">Refresh</button>
  </div>
);
const queryClient = new QueryClient();

export const AppProvider = ({ children }: Props) => (
  <ErrorBoundary FallbackComponent={ErrorFallback}>
    <HelmetProvider>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
      </BrowserRouter>
    </HelmetProvider>
  </ErrorBoundary>
);
