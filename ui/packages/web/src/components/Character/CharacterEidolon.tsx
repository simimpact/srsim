import { Fragment, useState } from "react";
import { AvatarRankConfig } from "@/bindings/AvatarRankConfig";
import { sanitizeNewline } from "@/utils/helpers";
import { Badge } from "../Primitives/Badge";
import { Toggle } from "../Primitives/Toggle";

interface Props {
  data: AvatarRankConfig[];
  characterId: number;
}
const CharacterEidolon = ({ data, characterId }: Props) => {
  const [selectedEidolon, setSelectedEidolon] = useState(1);

  const top = data.filter(e => e.rank <= 3).sort((a, b) => a.rank - b.rank);
  const bottom = data.filter(e => e.rank > 3).sort((a, b) => a.rank - b.rank);

  const currentEidolon = data.find(e => e.rank === selectedEidolon);

  return (
    <>
      <EidolonRow
        data={top}
        selectedEidolon={selectedEidolon}
        setSelectedEidolon={setSelectedEidolon}
        characterId={characterId}
      />

      <div className="my-2 min-h-[8rem] whitespace-pre-wrap rounded-md border p-4">
        {currentEidolon?.desc.map((descPart, index) => (
          <Fragment key={index}>
            <span className="whitespace-pre-wrap">{sanitizeNewline(descPart)}</span>
            <span className="font-semibold text-accent-foreground">
              {currentEidolon.param[index]}
            </span>
          </Fragment>
        ))}
      </div>

      <EidolonRow
        data={bottom.reverse()}
        selectedEidolon={selectedEidolon}
        setSelectedEidolon={setSelectedEidolon}
        characterId={characterId}
      />
    </>
  );
};

interface EidolonRowProps {
  data: AvatarRankConfig[];
  selectedEidolon: number;
  setSelectedEidolon: (value: number) => void;
  characterId: number;
}

const EidolonRow = ({
  data,
  selectedEidolon,
  setSelectedEidolon,
  characterId,
}: EidolonRowProps) => {
  return (
    <div className="grid grid-cols-3 gap-2">
      {data.map(eidolon => (
        <Toggle
          key={eidolon.rank_id}
          className="flex h-full flex-1 flex-col justify-start gap-2 py-2 sm:flex-row"
          pressed={selectedEidolon === eidolon.rank}
          onPressedChange={() => setSelectedEidolon(eidolon.rank)}
        >
          <div className="flex flex-col items-center gap-1">
            <img
              src={url(characterId, eidolon.rank)}
              alt={eidolon.name}
              onClick={() => setSelectedEidolon(eidolon.rank)}
              width={64}
              height={64}
              className="aspect-square min-w-[64px] invert dark:invert-0"
            />
            <Badge className="w-fit sm:inline">E{eidolon.rank}</Badge>
          </div>

          <span className="md:text-lg">{eidolon.name}</span>
        </Toggle>
      ))}
    </div>
  );
};

function url(charID: number, eidolon: number) {
  let fmt = `rank${eidolon}`;
  if (eidolon == 3) fmt = "skill";
  if (eidolon == 5) fmt = "ultimate";
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/skill/${charID}_${fmt}.png`;
}

export { CharacterEidolon };
