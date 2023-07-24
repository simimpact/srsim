import { EquipmentConfig } from "@/bindings/EquipmentConfig";
import useCardEffect from "@/hooks/animation/useCardEffect";
import { cn } from "@/utils/classname";

interface Props {
  data: EquipmentConfig;
}
export function LightConePortrait({ data }: Props) {
  const { flowRef, glowRef, removeListener, rotateToMouse } = useCardEffect();

  return (
    <div
      ref={flowRef}
      className={cn("relative h-fit w-full", "card")}
      onMouseLeave={removeListener}
      onMouseMove={rotateToMouse}
    >
      <img
        src={url(data.equipment_id)}
        width={902}
        height={1260}
        className="place-self-start object-contain"
        alt={data.equipment_name}
      />
      <div ref={glowRef} className="glow" />
    </div>
  );
}

function url(id: string | number): string {
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/image/light_cone_portrait/${id}.png`;
}
