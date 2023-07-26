import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactNode } from "react";
import { ErrorBoundary } from "react-error-boundary";
import { HelmetProvider } from "react-helmet-async";
import { BrowserRouter } from "react-router-dom";
import { Toaster } from "@/components/Primitives/Toast/Toaster";
import { TooltipProvider } from "@/components/Primitives/Tooltip";
import { SimControl } from "./SimControl";

interface Props {
  children: ReactNode;
}

const ErrorFallback = () => (
  <div className="flex h-screen w-screen flex-col items-center justify-center gap-5">
    <h2 className="text-xl">something went wrong {":("}</h2>
    <button className="font-blod rounded bg-blue-500 px-4 py-2 text-white">Refresh</button>
  </div>
);
const queryClient = new QueryClient({
  defaultOptions: { queries: { refetchOnWindowFocus: false, retry: false } },
});

export const AppProvider = ({ children }: Props) => {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>
      <HelmetProvider>
        <BrowserRouter>
          <QueryClientProvider client={queryClient}>
            <SimControl>
              <TooltipProvider delayDuration={0}>
                {children}
                <Toaster />
              </TooltipProvider>
            </SimControl>
          </QueryClientProvider>
        </BrowserRouter>
      </HelmetProvider>
    </ErrorBoundary>
  );
};
