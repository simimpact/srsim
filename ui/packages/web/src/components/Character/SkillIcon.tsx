import { HTMLAttributes, forwardRef } from "react";
import { AvatarSkillConfig } from "@/bindings/AvatarSkillConfig";
import { Badge } from "../Primitives/Badge";
import { Popover, PopoverContent, PopoverTrigger } from "../Primitives/Popover";
import { Tooltip, TooltipContent, TooltipTrigger } from "../Primitives/Tooltip";
import { SkillDescription } from "./SkillDescription";

interface Props extends HTMLAttributes<HTMLButtonElement> {
  disableTooltip?: boolean;
  disablePopover?: boolean;
  data: AvatarSkillConfig;
  characterId: number;
  slv?: number;
}
const SkillIcon = ({ disableTooltip = false, disablePopover = false, slv, ...props }: Props) => {
  if (disablePopover)
    return (
      <IconWithTooltip disableTooltip={disableTooltip} disablePopover={disablePopover} {...props} />
    );
  const { skill_desc, param_list } = props.data;
  return (
    <Popover>
      <PopoverTrigger asChild>
        <IconWithTooltip
          disableTooltip={disableTooltip}
          disablePopover={disablePopover}
          {...props}
        />
      </PopoverTrigger>
      <PopoverContent className="w-[600px]">
        <div className="text-accent-foreground flex items-center gap-2 text-lg font-semibold">
          <Badge className="h-fit p-1">{props.data.skill_type_desc}</Badge>
          <span>{props.data.skill_name}</span>
          {props.data.spneed?.value && <span>({props.data.spneed.value} Energy)</span>}
        </div>
        <SkillDescription skillDesc={skill_desc} paramList={param_list} slv={slv ?? 0} />
      </PopoverContent>
    </Popover>
  );
};

const IconWithTooltip = forwardRef<HTMLButtonElement, Props>(
  ({ disableTooltip, disablePopover, data, characterId, className, ...props }, ref) => {
    if (disableTooltip)
      return (
        <img
          src={getImagePath(characterId, data) ?? ""}
          alt={data.skill_name}
          className="invert dark:invert-0"
          width={64}
          height={64}
        />
      );
    return (
      <Tooltip disableHoverableContent>
        <TooltipTrigger asChild>
          <button ref={ref} className={className} {...props}>
            <img
              src={getImagePath(characterId, data) ?? ""}
              alt={data.skill_name}
              className="invert dark:invert-0"
              width={64}
              height={64}
            />
          </button>
        </TooltipTrigger>
        <TooltipContent className="select-none">{data.skill_name}</TooltipContent>
      </Tooltip>
    );
  }
);

function getImagePath(
  characterId: number | null | undefined,
  skill: AvatarSkillConfig
): string | undefined {
  let ttype = "";
  if (skill.attack_type) {
    switch (skill.attack_type) {
      case "Normal":
        ttype = "basic_atk";
        break;
      case "BPSkill":
        ttype = "skill";
        break;
      case "Ultra":
        ttype = "ultimate";
        break;
      case "Talent":
        ttype = "talent";
        break;
      case "Maze":
        ttype = "technique";
        break;
    }
  } else {
    switch (skill.skill_type_desc) {
      case "Basic ATK":
        ttype = "basic_atk";
        break;
      case "Skill":
        ttype = "skill";
        break;
      case "Ultra":
        ttype = "ultimate";
        break;
      case "Talent":
        ttype = "talent";
        break;
      case "Technique":
        ttype = "technique";
        break;
    }
  }
  if (!characterId) return undefined;
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/skill/${characterId}_${ttype}.png`;
}

export { SkillIcon };
