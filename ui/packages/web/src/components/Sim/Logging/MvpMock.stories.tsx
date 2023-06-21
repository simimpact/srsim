import { Meta, StoryObj } from "@storybook/react";
import { MvpMock } from "./MvpMock";

const meta = {
  title: "Sim/MvpMock",
  component: MvpMock,
} satisfies Meta<typeof MvpMock>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  args: {
    name: "aaaa",
  },
} satisfies Story;
