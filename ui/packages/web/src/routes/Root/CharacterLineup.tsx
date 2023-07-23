import { event } from "@srsim/types";
import { useQueries } from "@tanstack/react-query";
import { ReactNode, useContext } from "react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/Primitives/Tooltip";
import { CharacterPortrait } from "@/components/Sim/CharacterPortrait";
import { SimControlContext } from "@/providers/SimControl";
import { cn } from "@/utils/classname";
import API from "@/utils/constants";

interface Props {
  isEnemy?: boolean;
  header?: ReactNode;
}
const CharacterLineup = ({ isEnemy = false, header }: Props) => {
  const { simulationData } = useContext(SimControlContext);
  const initEvent = simulationData.find(e => e.name == "Initialize")?.event as
    | event.Initialize
    | undefined;
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
  const characters: { key: string; light_cone: { key: string } }[] =
    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
    initEvent?.config.characters ?? [];
  console.log(characters);

  const characterQueries = useQueries({
    queries: getQueries(characters),
  });
  // console.log(characterQueries[0].data);

  if (characterQueries.map(e => e.data).every(e => !e)) return null;

  return (
    <div className={cn("flex flex-col rounded-md p-2", isEnemy ? "bg-destructive" : "bg-accent")}>
      {/* NOTE: CharacterCard is based for now, not yet implemented */}
      <div className="flex justify-center">{header}</div>
      <div className="flex">
        {characterQueries.map(({ data }, index) =>
          data ? (
            // <TempCharCard key={code} name={name} code={code} />
            <TooltipProvider key={index} delayDuration={0}>
              <Tooltip>
                <TooltipTrigger>
                  <CharacterPortrait data={data} />
                </TooltipTrigger>
                <TooltipContent>{data.avatar_name}</TooltipContent>
              </Tooltip>
            </TooltipProvider>
          ) : null
        )}
      </div>
    </div>
  );
};

function getQueries(characters: { key: string }[]) {
  return characters.map(character => {
    return {
      queryKey: ["character", character.key],
      queryFn: async () => await API.character.get(character.key),
      enabled: !!character.key,
    };
  });
}

export { CharacterLineup };

export type Path =
  | "Destruction"
  | "Hunt"
  | "Erudition"
  | "Harmony"
  | "Nihility"
  | "Preservation"
  | "Abundance";

/**
 * metadata for light cones
 */

export interface EquipmentConfig {
  avatar_base_type: Path;
  coin_cost: number;
  equipment_desc: string;
  equipment_id: number;
  equipment_name: string;
  exp_provide: number;
  exp_type: number;
  max_promotion: number;
  max_rank: number;
  rank_up_cost_list: number[];
  rarity: number;
  release: boolean;
  skill_id: number;
}

export type Element = "Fire" | "Ice" | "Physical" | "Wind" | "Lightning" | "Quantum" | "Imaginary";

export interface AvatarConfig {
  avatar_base_type: Path;
  avatar_desc: string;
  avatar_id: number;
  avatar_name: string;
  avatar_votag: string;
  damage_type: Element;
  damage_type_resistance: DamageTypeResistance[];
  rank_idlist: number[];
  rarity: number;
  release: boolean;
  skill_list: number[];
  spneed: number;
}

export interface DamageTypeResistance {
  damage_type: Element;
  value: Param;
}

export interface Param {
  value: number;
}
