import { Fragment } from "react";
import { IconMap } from "@/hooks/transform/usePropertyMap";

interface Props {
  backendData?: unknown;
}

interface Mock {
  property: keyof typeof IconMap;
  label: string;
  value: number;
  bonus: 100;
  merge: boolean;
  percent: boolean;
}

// value = default params, from character
// bonus = from relics + traces
const MOCK_DATA: Mock[] = [
  { property: "ATK_BASE", label: "Atk", value: 100, bonus: 100, merge: false, percent: false },
  { property: "HP_BASE", label: "HP", value: 100, bonus: 100, merge: false, percent: false },
  { property: "DEF_BASE", label: "DEF", value: 100, bonus: 100, merge: false, percent: false },
  { property: "SPD_BASE", label: "SPD", value: 100, bonus: 100, merge: false, percent: true },
  {
    property: "CRIT_CHANCE",
    label: "CRIT Rate",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  { property: "CRIT_DMG", label: "CRIT DMG", value: 100, bonus: 100, merge: true, percent: true },
  {
    property: "BREAK_EFFECT",
    label: "Break Effect",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  {
    property: "ENERGY_REGEN_CONVERT",
    label: "Energy Regeneration",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  {
    property: "EFFECT_HIT_RATE",
    label: "Effect Hit Rate",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  {
    property: "EFFECT_RES",
    label: "Effect Res",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  {
    property: "ICE_DMG_PERCENT",
    label: "Ice DMG Boost",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
  {
    property: "WIND_DMG_PERCENT",
    label: "Wind DMG Boost",
    value: 100,
    bonus: 100,
    merge: true,
    percent: true,
  },
];

// TODO: take above mock data as props
const CharacterStatTable = ({ backendData }: Props) => {
  if (backendData) console.log(backendData);
  return (
    <div className="grid grid-cols-2 gap-y-2">
      {MOCK_DATA.map(({ property, label, value, bonus, merge, percent }, index) => (
        <Fragment key={index}>
          <div className="flex">
            <img src={iconUrl(property)} alt={property} className="invert-0 dark:invert" />
            <div>{label}</div>
          </div>
          {merge ? (
            <span>
              {value + bonus} {percent && "%"}
            </span>
          ) : (
            <span className="whitespace-nowrap">
              {value} + <span className="text-wind">{bonus}</span> {percent && "%"}
            </span>
          )}
        </Fragment>
      ))}
    </div>
  );
};
export { CharacterStatTable };

function iconUrl(stat: keyof typeof IconMap) {
  return `/icons/${IconMap[stat]}.svg`;
}
