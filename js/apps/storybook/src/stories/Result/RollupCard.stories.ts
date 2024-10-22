import { RollupCard } from "@ui/components";
import type { Meta, StoryObj } from "@storybook/react";

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta: Meta<typeof RollupCard> = {
  title: "Result/RollupCard",
  component: RollupCard,
  parameters: {
    // Optional parameter to center the component in the Canvas. More info: https://storybook.js.org/docs/configure/story-layout
    layout: "padded",
  },
  // This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
  tags: ["autodocs"],
  // More on argTypes: https://storybook.js.org/docs/api/argtypes
  argTypes: {},
  // Use `fn` to spy on the onClick arg, which will appear in the actions panel once invoked: https://storybook.js.org/docs/essentials/actions#action-args
  args: {},
};

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
  args: {
    color: "bg-red",
    title: "Damage Per Cycle (DPC)",
    value: "66,000",
    auxStats: [
      { title: "min", value: "23,456" },
      { title: "max", value: "23,456" },
      { title: "std", value: "23,456" },
      { title: "p25", value: "23,456" },
      { title: "p50", value: "23,456" },
      { title: "p75", value: "23,456" },
    ],
    tooltip: "some help tip",
  },
};
