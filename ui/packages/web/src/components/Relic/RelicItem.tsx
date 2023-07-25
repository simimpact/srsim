import { useRelicSearch } from "@/hooks/queries/useRelic";
import { IconMap, getPropertyMap } from "@/hooks/transform/usePropertyMap";
import { Relic } from "@/providers/temporarySimControlTypes";
import { asPercentage } from "@/utils/helpers";
import { Badge } from "../Primitives/Badge";
import { Separator } from "../Primitives/Separator";

interface Props {
  data: Relic;
  mockIndex: number;
  asSet?: boolean;
}
const RelicItem = ({ data, mockIndex, asSet = false }: Props) => {
  const { relicSetConfig } = useRelicSearch(data.key);

  if (!relicSetConfig) return null;

  // TODO: render case for isSet
  return (
    <div className="flex rounded-md border p-2">
      <img src={url(relicSetConfig.set_id, mockIndex, asSet)} width={96} height={96} />

      <Badge className="-ml-4 h-min rounded-full px-1">+15</Badge>

      <div id="mainstat" className="grid grid-cols-1 grid-rows-2 pl-3">
        <img
          src={iconUrl("HP_BASE")}
          alt={"HP_BASE"}
          height={32}
          width={32}
          className="self-center invert-0 dark:invert"
        />
        <span className="self-center text-center text-lg font-bold">290</span>
      </div>

      <Separator className="mx-3" orientation="vertical" />

      <div id="stats" className="grid grow grid-cols-2">
        {data.sub_stats.map(({ stat, amount }, index) => (
          <div key={index} className="flex items-center">
            <img
              src={iconUrl(stat as keyof typeof IconMap)}
              alt={stat}
              height={32}
              width={32}
              className="invert-0 dark:invert"
            />

            {getPropertyMap(stat as keyof typeof IconMap).isPercent ? asPercentage(amount) : amount}
          </div>
        ))}
      </div>
    </div>
  );
};

function url(setId: number, mockIndex: number, isSet: boolean) {
  const value = isSet ? setId : `${setId}_${mockIndex}`;
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/relic/${value}.png`;
}

function iconUrl(stat: keyof typeof IconMap) {
  return `/icons/${getPropertyMap(stat).dmValue}.svg`;
}

export { RelicItem };
