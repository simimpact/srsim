import { HTMLAttributes, forwardRef } from "react";
import { cn } from "@/utils/classname";

interface Props extends HTMLAttributes<HTMLDivElement> {
  imposition: number;
}
const ImpositionIcon = forwardRef<HTMLDivElement, Props>(
  ({ imposition, className, ...props }, ref) => (
    <div
      className={cn(
        "bg-background rounded-full aspect-square text-center w-6 h-6 font-medium font-[Cinzel]",
        imposition >= 5 ? "text-[#191919] bg-[#F9CC71]" : "text-[#F9CC71] bg-[#191919]",
        className
      )}
      ref={ref}
      {...props}
    >
      {asRoman(imposition)}
    </div>
  )
);
export { ImpositionIcon };

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
