import { Meta, StoryObj } from "@storybook/react";
import { Plus } from "lucide-react";
import { Button } from "../Button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "../Tooltip";

const meta = {
  title: "Primitives/Tooltip",
  component: Tooltip,
} satisfies Meta<typeof Tooltip>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default = {
  render: () => (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button variant="outline" className="w-10 rounded-full p-0">
            <Plus className="h-4 w-4" />
            <span className="sr-only">Add</span>
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>Add to library</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  ),
} satisfies Story;
