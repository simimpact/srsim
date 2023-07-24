import { AvatarSkillConfig } from "@/bindings/AvatarSkillConfig";

interface Props {
  disableTooltip?: boolean;
  disablePopover?: boolean;
  data: AvatarSkillConfig;
  characterId: number;
}
const SkillIcon = ({
  disableTooltip = false,
  disablePopover = false,
  data,
  characterId,
}: Props) => {
  return (
    <img
      src={getImagePath(characterId, data) ?? ""}
      alt={data.skill_name}
      className="invert dark:invert-0"
      width={64}
      height={64}
    />
  );
};

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
