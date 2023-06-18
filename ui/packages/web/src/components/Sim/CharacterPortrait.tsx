import { VariantProps } from "class-variance-authority";
import { cn } from "@/utils/classname";
import { elementVariants, rarityVariants } from "@/utils/variants";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../Primitives/Dialog";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "../Primitives/Tooltip";

interface Props extends VariantProps<typeof elementVariants> {
  name: string;
  code: number;
  rarity: number;
}

const CharacterPortrait = ({ code, name, element, rarity }: Props) => {
  let stringrarity: "green" | "blue" | "purple" | "gold" | "silver" | undefined = undefined;
  if (rarity === 1) stringrarity = "silver";
  if (rarity === 2) stringrarity = "green";
  if (rarity === 3) stringrarity = "blue";
  if (rarity === 4) stringrarity = "purple";
  if (rarity === 5) stringrarity = "gold";

  return (
    <TooltipProvider delayDuration={0}>
      <Tooltip>
        <TooltipTrigger>
          <Dialog>
            <DialogTrigger>
              <img
                src={url(code)}
                alt={name}
                className={cn(
                  "max-h-32 max-w-32 rounded-full box-content p-1 border-2",
                  elementVariants({ border: element }),
                  rarityVariants({ rarity: stringrarity })
                )}
              />
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{name}</DialogTitle>
                <DialogDescription>{rarity} ✦ Erudition</DialogDescription>
              </DialogHeader>
              <img
                src={url(code)}
                alt={name}
                className={cn(
                  "h-32 w-32 rounded-full box-content p-1 border-2",
                  elementVariants({ border: element }),
                  rarityVariants({ rarity: stringrarity })
                )}
              />
            </DialogContent>
          </Dialog>
        </TooltipTrigger>
        <TooltipContent>{name} (click on me!)</TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
};

function url(characterCode: number): string {
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/character/${characterCode}.png`;
}
export { CharacterPortrait };
