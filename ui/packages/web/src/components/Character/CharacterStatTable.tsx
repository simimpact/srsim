import { Fragment, HTMLAttributes, forwardRef } from "react";
import { IconMap, getPropertyMap } from "@/hooks/transform/usePropertyMap";
import { asPercentage } from "@/utils/helpers";

interface Props extends HTMLAttributes<HTMLDivElement> {
  backendData?: unknown;
}

interface Mock {
  prop: keyof typeof IconMap;
  label: string;
  val: number;
  bonus: number;
  merge: boolean;
}

// value = default params, from character
// bonus = from relics + traces
const MOCK_DATA: Mock[] = [
  { prop: "ATK_BASE", label: "Atk", val: 100, bonus: 100, merge: false },
  { prop: "HP_BASE", label: "HP", val: 100, bonus: 100, merge: false },
  { prop: "DEF_BASE", label: "DEF", val: 100, bonus: 100, merge: false },
  { prop: "SPD_BASE", label: "SPD", val: 100, bonus: 100, merge: false },
  { prop: "CRIT_CHANCE", label: "CRIT Rate", val: 0.05, bonus: 0.5, merge: true },
  { prop: "CRIT_DMG", label: "CRIT DMG", val: 0.5, bonus: 0.7, merge: true },
  { prop: "BREAK_EFFECT", label: "Break Effect", val: 0.05, bonus: 1, merge: true },
  { prop: "ENERGY_REGEN_CONVERT", label: "Energy Limit", val: 420, bonus: 0, merge: true },
  { prop: "ENERGY_REGEN", label: "Energy Regeneration", val: 0.5, bonus: 0.38, merge: true },
  { prop: "EFFECT_HIT_RATE", label: "Effect Hit Rate", val: 0.5, bonus: 0.38, merge: true },
  { prop: "EFFECT_RES", label: "Effect Res", val: 0.5, bonus: 0.38, merge: true },
  { prop: "ICE_DMG_PERCENT", label: "Ice DMG Boost", val: 0.5, bonus: 0.38, merge: true },
  { prop: "WIND_DMG_PERCENT", label: "Wind DMG Boost", val: 0.5, bonus: 0.38, merge: true },
];

// TODO: take above mock data as props
const CharacterStatTable = forwardRef<HTMLDivElement, Props>(
  ({ backendData, className, ...props }, ref) => {
    if (backendData) console.log(backendData);
    return (
      <div className={className} ref={ref} {...props}>
        {MOCK_DATA.map(({ prop: property, label, val: value, bonus, merge }) => (
          <Fragment key={property}>
            <div className="flex gap-2 items-center">
              <img
                src={iconUrl(property)}
                alt={property}
                width={24}
                height={24}
                className="h-6 w-6 invert-0 dark:invert aspect-square"
              />
              <div>{label}</div>
            </div>
            {merge ? (
              <span>
                {getPropertyMap(property).isPercent ? asPercentage(value + bonus) : value + bonus}
              </span>
            ) : (
              <span className="whitespace-nowrap">
                {value} + <span className="text-wind">{bonus}</span>{" "}
                {getPropertyMap(property).isPercent && "%"}
              </span>
            )}
          </Fragment>
        ))}
      </div>
    );
  }
);
export { CharacterStatTable };

function iconUrl(stat: keyof typeof IconMap) {
  return `/icons/${IconMap[stat]}.svg`;
}
