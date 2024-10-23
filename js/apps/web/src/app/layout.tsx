import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "@ui/styles/globals.css";
import { ThemeProvider } from "../components/theme-provider";
import { cn } from "ui/src/lib/utils";
import { Toaster } from "ui";
import { ViewerProvider } from "./viewer/provider";
import { ExecutorProvider } from "./exec/provider";
import Navigation from "./navigation";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "srsim",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={cn("bg-background h-screen", inter.className)}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <ExecutorProvider>
            <ViewerProvider>
              <>
                <Navigation />
                {children}
              </>
            </ViewerProvider>
          </ExecutorProvider>
          <Toaster />
        </ThemeProvider>
      </body>
    </html>
  );
}
