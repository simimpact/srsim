import { Meta, StoryObj } from "@storybook/react";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "../HoverCard";

const meta = {
  title: "Primitives/HoverCard",
  component: HoverCard,
} satisfies Meta<typeof HoverCard>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  render: () => (
    <HoverCard>
      <HoverCardTrigger>Hover</HoverCardTrigger>
      <HoverCardContent>The React Framework – created and maintained by @vercel.</HoverCardContent>
    </HoverCard>
  ),
} satisfies Story;
