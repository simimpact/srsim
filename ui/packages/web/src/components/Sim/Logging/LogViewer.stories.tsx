import { Meta, StoryObj } from "@storybook/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { LogViewer } from "./LogViewer";

const queryClient = new QueryClient();
const meta = {
  title: "Sim/LogViewer",
  component: LogViewer,
  decorators: [Story => <QueryClientProvider client={queryClient}> {Story()}</QueryClientProvider>],
} satisfies Meta<typeof LogViewer>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  args: {
    placeholder: "test",
  },
} satisfies Story;
