import { ReactNode } from "react";
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
    <div className={cn("flex flex-col rounded-md p-2", isEnemy ? "bg-red-500" : "bg-blue-500")}>
      {/* NOTE: CharacterCard is based for now, not yet implemented */}
      <div className="flex justify-center">{header}</div>
      <div className="flex">
        {charCodes.map(({ name, code }) => (
          <TempCharCard key={code} name={name} code={code} />
        ))}
      </div>
    </div>
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
