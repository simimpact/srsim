import { Meta, StoryObj } from "@storybook/react";
import { CharacterPortrait } from "./CharacterPortrait";

const meta = {
  title: "Sim/CharacterPortrait",
  component: CharacterPortrait,
} satisfies Meta<typeof CharacterPortrait>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  args: {
    code: 1102,
    name: "seele",
    rarity: 5,
    element: "quantum",
  },
} satisfies Story;
