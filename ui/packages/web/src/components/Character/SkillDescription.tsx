import { Fragment } from "react";
import { ParameterizedDescription } from "@/bindings/SkillTreeConfig";
import { sanitizeNewline } from "@/utils/helpers";

interface SkillDescriptionProps {
  skillDesc: ParameterizedDescription;
  paramList: string[][] | string[];
  slv: number;
}

export const SkillDescription = ({ skillDesc, paramList, slv }: SkillDescriptionProps) => {
  // for depth of 2, flatten once
  const currentParam = Array.isArray(paramList[0]) ? paramList[slv] : paramList;

  return (
    <p className="text-justify">
      {skillDesc.map((descPart, index) => (
        <Fragment key={index}>
          <span className="whitespace-pre-wrap">{sanitizeNewline(descPart)}</span>
          <span className="text-accent-foreground font-semibold">{currentParam[index]}</span>
        </Fragment>
      ))}
    </p>
  );
};
