import { Meta, StoryObj } from "@storybook/react";
import { ComplexForm } from "./ComplexForm";

const meta = {
  title: "ComplexData/FormUsingState",
  component: ComplexForm,
} satisfies Meta<typeof ComplexForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {} satisfies Story;
