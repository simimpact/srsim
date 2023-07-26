import { RelicSetConfig } from "@/bindings/RelicSetConfig";
import { RelicSetSkillConfig } from "@/bindings/RelicSetSkillConfig";
import { cn } from "@/utils/classname";
import { SkillDescription } from "../Character/SkillDescription";
import { Badge } from "../Primitives/Badge";

interface Props {
  pieceActivated: number;
  bonusCfg: RelicSetSkillConfig | undefined;
  setCfg: RelicSetConfig | undefined;
}
const RelicSetBonus = ({ setCfg, pieceActivated, bonusCfg }: Props) => {
  if (!bonusCfg || !setCfg) return null;

  return (
    <div className="flex flex-col gap-2">
      <div className="mb-1 mt-2 text-xl font-semibold">{setCfg.set_name}</div>

      {bonusCfg.require_num.map((requireNum, index) => (
        <div key={requireNum} className={cn(pieceActivated >= requireNum ? "text-wind-500" : "")}>
          <Badge>{requireNum} Pcs</Badge>

          <SkillDescription
            skillDesc={bonusCfg.skill_desc[index]}
            paramList={bonusCfg.ability_param_list[index]}
            slv={0}
          />
        </div>
      ))}
    </div>
  );
};
export { RelicSetBonus };
