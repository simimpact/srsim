import { HTMLAttributes, forwardRef } from "react";
import { Path } from "@/bindings/AvatarConfig";
import { Element } from "@/bindings/PatchBanner";
import useCardEffect from "@/hooks/animation/useCardEffect";
import { cn } from "@/utils/classname";
import { range } from "@/utils/helpers";

interface Props {
  rarity?: number;
  element?: Element;
  path?: Path;
  name: string;
  imgUrl: string;
}

const LightConeCard = ({ rarity, element, path, name, imgUrl }: Props) => {
  const { flowRef, glowRef, removeListener, rotateToMouse } = useCardEffect();

  if (element) console.log("unimplemented");
  if (path) console.log("unimplemented");

  return (
    <div>
      <div
        ref={flowRef}
        className="relative h-full w-full transition-all ease-out"
        onMouseLeave={removeListener}
        onMouseMove={rotateToMouse}
        style={{ perspective: "1500px" }}
      >
        <div className={cn("absolute left-[18%] top-[14%] h-[76%] w-[65%] rotate-[13deg]", "card")}>
          <div ref={glowRef} className="glow" />
        </div>
        <img src={imgUrl} alt={name} width={374} height={512} />
        {/* {element && <ElementIcon element={element} size="15%" className="absolute left-1 top-0" />} */}
      </div>
      {/* {path && (
        <PathIcon
          path={path}
          size="15%"
          className={cn("absolute left-1 text-white", element ? "top-[15%]" : "top-0")}
        />
      )} */}
      {rarity && <RarityIcon rarity={rarity} className="-my-6 h-6 w-full" />}
    </div>
  );
};
export { LightConeCard };

interface RarityIconProps extends HTMLAttributes<HTMLDivElement> {
  rarity: number;
}
const RarityIcon = forwardRef<HTMLDivElement, RarityIconProps>(
  ({ rarity, className, ...props }, ref) => (
    <div className={cn("absolute flex justify-center", className)} {...props} ref={ref}>
      {Array.from(range(1, rarity, 1)).map((_, index) => (
        <div key={index} className="aspect-square">
          <img
            src="/images/Star.png"
            height={128}
            width={128}
            alt={`${rarity} ✦`}
            className="pointer-events-none"
          />
        </div>
      ))}
    </div>
  )
);
