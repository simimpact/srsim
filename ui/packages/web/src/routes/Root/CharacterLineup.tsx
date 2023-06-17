import { ReactNode } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTrigger,
} from "@/components/Primitives/Dialog";
import { cn } from "@/utils/classname";

interface Props {
  isEnemy?: boolean;
  header?: ReactNode;
}
const CharacterLineup = ({ isEnemy = false, header }: Props) => {
  const charCodes = [
    { name: "Trailblaze (Fire)", code: 8004 },
    { name: "Natasha", code: 1105 },
    { name: "Bronya", code: 1101 },
    { name: "Kafka", code: 1005 },
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
            {charCodes.map(({ name, code }) => (
              <TempCharCard key={code} name={name} code={code} />
            ))}
          </div>
        </div>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>Team name 123128</DialogHeader>
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

interface CardProps {
  name: string;
  code: number; // for img pathname
}
const TempCharCard = ({ name, code }: CardProps) => {
  const url = (code: number) =>
    `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/character/${code}.png`;
  return (
    <div>
      <img src={url(code)} alt={name} className="max-h-32" />
    </div>
  );
};
export { CharacterLineup };
