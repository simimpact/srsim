import { cva } from "class-variance-authority";
import { HTMLAttributes, forwardRef } from "react";
import { Path } from "@/bindings/AvatarConfig";
import { SkillTreeConfig } from "@/bindings/SkillTreeConfig";
import { cn } from "@/utils/classname";
import { Popover, PopoverContent, PopoverTrigger } from "../Primitives/Popover";
import { Tooltip, TooltipContent, TooltipTrigger } from "../Primitives/Tooltip";
import { SkillDescription } from "./SkillDescription";

interface Props {
  traces: SkillTreeConfig[];
  bigTraceAscension: number;
  path: Path;
  charTraces: number[];
  emptyBigTrace?: boolean;
}

const iconWrapVariant = cva("flex flex-col items-center", {
  variants: {
    variant: {
      SMALL: "",
      BIG: "bg-zinc-300",
      CORE: "",
    },
  },
});

const TraceTree = ({
  bigTraceAscension,
  traces,
  path,
  charTraces,
  emptyBigTrace = false,
}: Props) => {
  const toRenderTraces = pointTable(path)
    .find(e => e.ascension === bigTraceAscension)
    ?.points.map(
      point => traces.find(e => e.anchor == "Point" + String(point).padStart(2, "0")) ?? traces[0]
    );

  return (
    <div className={cn("flex flex-col items-center gap-2", emptyBigTrace ? "pt-[56px]" : "")}>
      {toRenderTraces?.map((traceNode, index) =>
        getNodeType(traceNode) !== "SMALL" ? (
          <Popover key={index}>
            <PopoverTrigger asChild>
              <IconWithTooltip
                node={traceNode}
                className={cn(
                  "rounded-full invert dark:invert-0",
                  !charTraces.includes(traceNode.point_id) ? "brightness-[.25]" : ""
                )}
              />
            </PopoverTrigger>
            <PopoverContent className="w-96">
              <span className="text-lg font-semibold text-accent-foreground">
                {traceNode.point_name}
              </span>
              <SkillDescription
                skillDesc={traceNode.point_desc}
                paramList={traceNode.param_list}
                slv={0}
              />
            </PopoverContent>
          </Popover>
        ) : (
          <IconWithTooltip
            key={index}
            node={traceNode}
            className={cn(
              "rounded-full invert dark:invert-0",
              !charTraces.includes(traceNode.point_id) ? "brightness-[.25]" : ""
            )}
          />
        )
      )}
    </div>
  );
};

interface IconProps extends HTMLAttributes<HTMLButtonElement> {
  node: SkillTreeConfig;
}
const IconWithTooltip = forwardRef<HTMLButtonElement, IconProps>(
  ({ node, className, ...props }, ref) => (
    <Tooltip disableHoverableContent>
      <TooltipTrigger asChild>
        <button
          ref={ref}
          className={cn(iconWrapVariant({ variant: getNodeType(node) }), className)}
          {...props}
        >
          <img
            src={traceIconUrl(node)}
            alt={String(node.point_id)}
            width={getNodeType(node) === "BIG" ? 48 : 32}
            height={getNodeType(node) === "BIG" ? 48 : 32}
            className={getNodeType(node) === "BIG" ? "invert" : ""}
          />
          {/* percentage */}
          {asPercentage(node)}
        </button>
      </TooltipTrigger>
      <TooltipContent className="select-none">{node.point_name}</TooltipContent>
    </Tooltip>
  )
);

function getNodeType(node: SkillTreeConfig): "CORE" | "SMALL" | "BIG" {
  if (node.icon_path.includes("_SkillTree")) return "BIG";
  if (
    ["Normal.png", "BP.png", "Maze.png", "Passive.png", "Ultra.png"].some(ends =>
      node.icon_path.endsWith(ends)
    )
  )
    return "CORE";
  return "SMALL";
}

function asPercentage(node: SkillTreeConfig): string | undefined {
  if (!node.status_add_list[0]) return undefined;
  const num = node.status_add_list[0].value.value;
  return Number(`${(num * 100).toFixed(2)}`).toString() + " %";
}

function traceIconUrl(node: SkillTreeConfig) {
  const base = `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon`;
  switch (getNodeType(node)) {
    case "CORE": {
      let path = "";
      const mapper = [
        { left: "Normal.png", right: "_basic_atk.png" },
        { left: "Passive.png", right: "_talent.png" },
        { left: "BP.png", right: "_skill.png" },
        { left: "Maze.png", right: "_technique.png" },
        { left: "Ultra.png", right: "_ultimate.png" },
      ];

      mapper.forEach(({ left, right }) => {
        if (node.icon_path.endsWith(left)) path = `/skill/${node.avatar_id}${right}`;
      });
      return base + path;
    }
    case "SMALL": {
      const lastSlash = node.icon_path.lastIndexOf("/");
      const name = node.icon_path.slice(lastSlash);
      return `${base}/property${name}`;
    }
    case "BIG": {
      // SkillTree1.png
      return `${base}/skill/${node.avatar_id}_${node.icon_path.slice(-14).toLowerCase()}`;
    }
  }
}

interface TraceTree {
  ascension: number;
  points: number[];
}

function pointTable(path: Path): TraceTree[] {
  switch (path) {
    case "Destruction":
      return [
        { ascension: 0, points: [9] },
        { ascension: 2, points: [6, 10, 11, 12] },
        { ascension: 4, points: [7, 13, 14, 15] },
        { ascension: 6, points: [8, 16, 17, 18] },
      ];
    case "Hunt":
      return [
        { ascension: 0, points: [9, 12, 15] },
        { ascension: 2, points: [6, 10, 11] },
        { ascension: 4, points: [7, 13, 14] },
        { ascension: 6, points: [8, 16, 17, 18] },
      ];
    case "Erudition":
      return [
        { ascension: 0, points: [9, 18] },
        { ascension: 2, points: [6, 10, 11, 12] },
        { ascension: 4, points: [7, 13, 14, 15] },
        { ascension: 6, points: [8, 16, 17] },
      ];
    case "Harmony":
      return [
        { ascension: 0, points: [9, 12, 15] },
        { ascension: 2, points: [6, 10, 11] },
        { ascension: 4, points: [7, 13, 14] },
        { ascension: 6, points: [8, 16, 17, 18] },
      ];
    case "Nihility":
      return [
        { ascension: 0, points: [9, 18] },
        { ascension: 2, points: [6, 10, 11, 12] },
        { ascension: 4, points: [7, 13, 14, 15] },
        { ascension: 6, points: [8, 16, 17] },
      ];
    case "Preservation":
      return [
        { ascension: 0, points: [9, 12, 15] },
        { ascension: 2, points: [6, 10, 11] },
        { ascension: 4, points: [7, 13, 14] },
        { ascension: 6, points: [8, 16, 17, 18] },
      ];
    case "Abundance":
      return [
        { ascension: 0, points: [9, 18] },
        { ascension: 2, points: [6, 10, 11, 12] },
        { ascension: 4, points: [7, 13, 14, 15] },
        { ascension: 6, points: [8, 16, 17] },
      ];
  }
}
export { TraceTree };
