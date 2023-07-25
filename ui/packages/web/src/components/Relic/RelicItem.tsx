import { useRelicSearch } from "@/hooks/queries/useRelic";
import { IconMap, getPropertyMap } from "@/hooks/transform/usePropertyMap";
import { Relic } from "@/providers/temporarySimControlTypes";
import { asPercentage } from "@/utils/helpers";

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
    <div className="flex">
      <img src={url(relicSetConfig.set_id, mockIndex, asSet)} width={64} height={64} />
      <div id="stats" className="grid grow grid-cols-2">
        {data.sub_stats.map(({ stat, amount }, index) => (
          <span key={index} className="flex">
            {/* TODO: color resolution */}
            <img
              src={iconUrl(stat as keyof typeof IconMap)}
              alt={stat}
              className="invert-0 dark:invert"
            />
            {getPropertyMap(stat as keyof typeof IconMap).isPercent ? asPercentage(amount) : amount}
          </span>
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
