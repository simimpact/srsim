import { useRelicSearch } from "@/hooks/queries/useRelic";
import { IconMap, getPropertyMap } from "@/hooks/transform/usePropertyMap";
import { Relic } from "@/providers/temporarySimControlTypes";
import { asPercentage } from "@/utils/helpers";
import { Badge } from "../Primitives/Badge";
import { Separator } from "../Primitives/Separator";

type RelicType = "HEAD" | "HAND" | "BODY" | "FOOT" | "OBJECT" | "NECK" | undefined;
interface Props {
  data: Relic;
  type: RelicType;
}
const RelicItem = ({ data, type }: Props) => {
  const { relicSetConfig } = useRelicSearch(data.key);

  if (!relicSetConfig) return null;

  return (
    <div className="flex rounded-md border p-2">
      <img src={url(relicSetConfig.set_id, type)} width={96} height={96} />

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

function url(setId: number, type: RelicType | undefined) {
  let index: number | undefined = undefined;
  switch (type) {
    case "HEAD":
      index = 0;
      break;
    case "HAND":
      index = 1;
      break;
    case "BODY":
      index = 2;
      break;
    case "FOOT":
      index = 3;
      break;
    case "OBJECT":
      index = 0;
      break;
    case "NECK":
      index = 1;
      break;
  }

  const value = !index ? setId : `${setId}_${index}`;
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/relic/${value}.png`;
}

function iconUrl(stat: keyof typeof IconMap) {
  return `/icons/${getPropertyMap(stat).dmValue}.svg`;
}

export { RelicItem };
