import { VariantProps } from "class-variance-authority";
import { ReactNode } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTrigger,
} from "@/components/Primitives/Dialog";
import { CharacterPortrait } from "@/components/Sim/CharacterPortrait";
import { cn } from "@/utils/classname";
import { elementVariants } from "@/utils/variants";

interface Props {
  isEnemy?: boolean;
  header?: ReactNode;
}
const CharacterLineup = ({ isEnemy = false, header }: Props) => {
  const charCodes: {
    name: string;
    code: number;
    rarity: number;
    element: VariantProps<typeof elementVariants>["element"];
  }[] = [
    { name: "Trailblaze (Fire)", code: 8004, rarity: 5, element: "fire" },
    { name: "Natasha", code: 1105, rarity: 4, element: "physical" },
    { name: "Bronya", code: 1101, rarity: 5, element: "wind" },
    { name: "Kafka", code: 1005, rarity: 5, element: "lightning" },
  ];

  return (
    <Dialog>
      <DialogTrigger>
        <div
          className={cn("flex flex-col rounded-md p-2", isEnemy ? "bg-destructive" : "bg-accent")}
        >
          {/* NOTE: CharacterCard is based for now, not yet implemented */}
          <div className="flex justify-center">{header}</div>
          <div className="flex">
            {charCodes.map(({ name, code, rarity, element }) => (
              // <TempCharCard key={code} name={name} code={code} />
              <CharacterPortrait
                key={code}
                name={name}
                code={code}
                rarity={rarity}
                element={element}
              />
            ))}
          </div>
        </div>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader className="text-card-foreground">Team name 123128</DialogHeader>
        <DialogDescription>
          <ul>
            {charCodes.map(({ name, code }) => (
              <li key={code}>{name}</li>
            ))}
          </ul>
        </DialogDescription>
      </DialogContent>
    </Dialog>
  );
};

export { CharacterLineup };
