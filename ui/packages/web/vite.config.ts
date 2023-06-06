import * as path from "path";
import react from "@vitejs/plugin-react";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

export default () => {
  return defineConfig({
    plugins: [react(), tsconfigPaths(), visualizer()],
    server: {
      proxy: {
        "/api": {
          target: "https://srsim.app",
          changeOrigin: true,
        },
        "/static": {
          target: "https://srsim.app",
          changeOrigin: true,
        },
      },
    },
    resolve: {
      alias: [{ find: "@", replacement: path.resolve(__dirname, "src") }],
    },
  });
};
