import { cva } from "class-variance-authority";
import { HTMLAttributes, forwardRef } from "react";
import SVG from "react-inlinesvg";
import { Path } from "@/bindings/AvatarConfig";
import { Element } from "@/bindings/PatchBanner";
import useCardEffect from "@/hooks/animation/useCardEffect";
import { cn } from "@/utils/classname";
import { range } from "@/utils/helpers";

interface Props extends HTMLAttributes<HTMLDivElement> {
  rarity: number;
  damage_type?: Element;
  avatar_base_type: Path;
  avatar_name: string;
  imgUrl: string;
}

const CharacterFloatCard = forwardRef<HTMLDivElement, Props>(
  ({ rarity, avatar_base_type, avatar_name, damage_type, imgUrl, className, ...props }, ref) => {
    const { glowRef, flowRef, rotateToMouse, removeListener } = useCardEffect();
    const elementVariants = cva("", {
      variants: {
        variant: {
          Fire: "text-fire",
          Ice: "text-ice",
          Wind: "text-wind",
          Physical: "text-physical",
          Lightning: "text-lightning",
          Quantum: "text-quantum",
          Imaginary: "text-imaginary",
        },
      },
    });

    return (
      <div
        className={cn("relative", className)}
        style={{ perspective: "1500px" }}
        ref={ref}
        {...props}
      >
        <div
          ref={flowRef}
          className={cn(
            "relative h-full w-full rounded-tr-3xl border-b-2 bg-gradient-to-b from-transparent from-80%  to-black/50",
            rarity === 5 ? "border-[#ffc870]" : "border-[#c199fd]",
            "card"
          )}
          onMouseLeave={removeListener}
          onMouseMove={rotateToMouse}
        >
          <img
            className={cn(
              "rounded-tr-3xl bg-gradient-to-b",
              rarity === 5 ? "bg-[#d0aa6e]" : "bg-[#9c65d7]"
            )}
            src={imgUrl}
            alt={avatar_name}
            width={374}
            height={512}
          />
          <div className="absolute left-1 top-0 w-[15%] h-[15%]">
            <SVG
              src={`/public/icons/${damage_type ?? "Physical"}.svg`}
              className={cn("w-full h-full", elementVariants({ variant: damage_type }))}
              style={{ filter: "drop-shadow(1px 1px 1px rgb(0 0 0 / 1))" }}
              viewBox="0 0 14 14"
            />
          </div>
          <div
            className={cn("absolute left-1 w-[15%] h-[15%]", damage_type ? "top-[15%]" : "top-0")}
          >
            <SVG
              src={`/public/icons/${avatar_base_type}.svg`}
              className="w-full h-full"
              style={{ filter: "drop-shadow(1px 1px 1px rgb(0 0 0 / 1))" }}
              viewBox="0 0 14 14"
            />
          </div>
          <RarityIcon rarity={rarity} className="top-[85%] h-6 w-full" />
          <div ref={glowRef} className={cn("rounded-tr-3xl", "glow")} />
        </div>
      </div>
    );
  }
);
export { CharacterFloatCard };

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
