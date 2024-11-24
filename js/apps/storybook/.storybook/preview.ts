import type { Preview } from "@storybook/react";
import "@ui/styles/globals.css";
import { INITIAL_VIEWPORTS, MINIMAL_VIEWPORTS } from "@storybook/addon-viewport";

const customViewports = {
  desktop1024: {
    name: "desktop-1024",
    styles: {
      width: "1024",
      height: "768",
    },
  },
  desktop1280: {
    name: "desktop-1280",
    styles: {
      width: "1280",
      height: "1024",
    },
  },
  desktop1366: {
    name: "desktop-1366",
    styles: {
      width: "1366",
      height: "768",
    },
  },
  desktop1920: {
    name: "desktop-1920",
    styles: {
      width: "1920",
      height: "1080",
    },
  },
  discord: {
    name: "discord",
    styles: {
      width: "520",
      height: "250",
    },
  },
};

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
    viewport: {
      viewports: {
        ...INITIAL_VIEWPORTS,
        ...MINIMAL_VIEWPORTS,
        ...customViewports,
      },
      defaultViewport: "desktop",
    },
  },
};

export default preview;
