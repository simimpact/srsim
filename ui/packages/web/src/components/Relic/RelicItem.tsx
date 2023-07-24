import { useQuery } from "@tanstack/react-query";
import { IconMap } from "@/hooks/transform/usePropertyMap";
import { Relic } from "@/providers/temporarySimControlTypes";
import API from "@/utils/constants";

interface Props {
  data: Relic;
  mockIndex: number;
}
const RelicItem = ({ data, mockIndex }: Props) => {
  const { data: relicSetConfig } = useQuery({
    queryKey: ["relicSet", data.key],
    queryFn: async () => await API.relicSet.get(data.key),
  });

  if (!relicSetConfig) return null;

  return (
    <div className="flex">
      <img src={url(relicSetConfig.set_id, mockIndex)} width={64} height={64} />
      <div id="stats" className="grid grid-cols-2 grow">
        {data.sub_stats.map(({ stat, amount }, index) => (
          <span key={index} className="flex">
            {/* TODO: color resolution */}
            <img src={iconUrl(stat)} alt={stat} className="invert-0 dark:invert" />
            {amount}
          </span>
        ))}
      </div>
    </div>
  );
};

function url(setId: number, mockIndex: number) {
  const value = mockIndex == 0 ? setId : `${setId}_${mockIndex}`;
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/relic/${value}.png`;
}

function iconUrl(stat: keyof typeof IconMap) {
  return `/icons/${IconMap[stat]}.svg`;
}

export { RelicItem };
