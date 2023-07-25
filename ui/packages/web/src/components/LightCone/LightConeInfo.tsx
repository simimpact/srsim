import { HTMLAttributes, forwardRef } from "react";

interface Props extends HTMLAttributes<HTMLDivElement> {
  name: string;
  imposition: number;
  currentLevel: number;
  maxLevel: number;
}
const LightConeInfo = forwardRef<HTMLDivElement, Props>(
  ({ name, imposition, currentLevel, maxLevel, className, ...props }, ref) => {
    return (
      <div id="lc-info" className={className} ref={ref} {...props}>
        <span>{name}</span>
        <div className="flex gap-2">
          <div className="bg-background flex aspect-square items-center justify-center rounded-full p-1 text-sm">
            {asRoman(imposition)}
          </div>
          <div>
            Lv. {currentLevel} / {maxLevel}
          </div>
        </div>
      </div>
    );
  }
);

export { LightConeInfo };

function asRoman(value: number) {
  switch (value) {
    case 1:
      return "I";
    case 2:
      return "II";
    case 3:
      return "III";
    case 4:
      return "IV";
    case 5:
      return "V";
    case 6:
      return "VI";
  }
}
