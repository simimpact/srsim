import { useQueries } from "@tanstack/react-query";
import { ReactNode, useContext } from "react";
import { AvatarConfig } from "@/bindings/AvatarConfig";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/Primitives/Tooltip";
import { SimControlContext } from "@/providers/SimControl";
import { cn } from "@/utils/classname";
import API, { characterIconUrl } from "@/utils/constants";
import { elementVariants, rarityVariants } from "@/utils/variants";

interface Props {
  isEnemy?: boolean;
  header?: ReactNode;
  onCharacterSelect?: (characterData: AvatarConfig, index: number) => void;
}
const CharacterLineup = ({ isEnemy = false, header, onCharacterSelect }: Props) => {
  const { simulationConfig } = useContext(SimControlContext);

  const characterQueries = useQueries({
    queries: getQueries(simulationConfig?.characters ?? []),
  });
  // console.log(characterQueries[0].data);

  if (characterQueries.map(e => e.data).every(e => !e)) return null;

  return (
    <div className={cn("flex flex-col rounded-md p-2", isEnemy ? "bg-destructive" : "bg-accent")}>
      {/* NOTE: CharacterCard is based for now, not yet implemented */}
      <div className="flex justify-center">{header}</div>
      <div className="flex gap-2">
        {characterQueries.map(({ data }, index) =>
          data ? (
            // <TempCharCard key={code} name={name} code={code} />
            <TooltipProvider key={index} delayDuration={0}>
              <Tooltip>
                <TooltipTrigger>
                  <div
                    onClick={() => {
                      if (onCharacterSelect) onCharacterSelect(data, index);
                    }}
                  >
                    <img
                      src={characterIconUrl(data.avatar_id)}
                      alt={data.avatar_name}
                      className={cn(
                        "box-content max-h-12 rounded-full border",
                        elementVariants({ border: data.damage_type }),
                        rarityVariants({ rarity: data.rarity as 1 | 2 | 3 | 4 | 5 | null })
                      )}
                    />
                  </div>
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
      queryFn: async () => await API.characterSearch.get(character.key),
      enabled: !!character.key,
    };
  });
}

export { CharacterLineup };
