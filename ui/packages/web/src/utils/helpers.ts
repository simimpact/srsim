import { SkillType } from "@/bindings/AvatarSkillConfig";

export function* range(start: number, end: number, step: number) {
  while (start <= end) {
    yield start;
    start += step;
  }
}

export function parseSkillType(skillType: SkillType | undefined | null, fallbackSkillDesc: string) {
  if (skillType)
    switch (skillType) {
      case "Normal":
        return "Attack";
      case "BPSkill":
        return "Skill";
      case "Ultra":
        return "Ultimate";
      case "Talent":
        return "Talent";
      case "MazeNormal":
        return "Attack";
      case "Maze":
        return "Technique";
    }
  else
    switch (fallbackSkillDesc) {
      case "Basic ATK":
        return "Attack";
      case "Skill":
        return "BPSkill";
      case "Ultra":
        return "Ultimate";
      case "Talent":
        return "Talent";
      case "Technique":
        return "Technique";
    }
  return "";
}

export function sanitizeNewline(data?: string) {
  if (!data) return "";
  return data.replaceAll("\\n", "\n");
}

/**
 * This function removes trailing zeroes if it's a whole number (eg. 18.00)
 * Otherwise a float percent with 2 decimals is returned
 * undefined number will return '0 %'
 */
export function asPercentage(data: number | undefined): string {
  if (!data) return "0 %";
  else {
    return Number(`${(data * 100).toFixed(2)}`).toString() + " %";
  }
}
