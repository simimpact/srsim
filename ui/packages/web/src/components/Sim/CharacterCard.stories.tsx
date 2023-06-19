import { Meta, StoryObj } from "@storybook/react";
import { CharacterCard } from "./CharacterCard";

const meta = {
  title: "Sim/CharacterCard",
  component: CharacterCard,
} satisfies Meta<typeof CharacterCard>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  args: {
    name: "Seele",
    rarity: 5,
  },
} satisfies Story;
