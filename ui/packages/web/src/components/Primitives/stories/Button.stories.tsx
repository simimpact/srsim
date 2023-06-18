import { Meta, StoryObj } from "@storybook/react";
import { Button } from "../Button";

const meta = {
  title: "Primitives/Button",
  component: Button,
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  args: {
    variant: "default",
    children: "This is a normal <Button />",
  },
} satisfies Story;

export const Warning = {
  args: {
    variant: "destructive",
    children: "this is a warning <Button />",
  },
} satisfies Story;

export const SecondaryAndSize = {
  args: {
    variant: "secondary",
    size: "lg",
    children: "this is a secondary <Button /> with custom size",
  },
} satisfies Story;

export const CustomClassName = {
  args: {
    variant: "secondary",
    className: "w-1/2 h-24",
    children: "You can still mix and match with your own tailwind class to the <Button />",
  },
} satisfies Story;
