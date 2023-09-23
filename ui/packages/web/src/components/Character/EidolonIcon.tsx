import { ButtonHTMLAttributes, Fragment, HTMLAttributes, forwardRef } from "react";
import { AvatarRankConfig } from "@/bindings/AvatarRankConfig";
import { cn } from "@/utils/classname";
import { sanitizeNewline } from "@/utils/helpers";
import { Popover, PopoverContent, PopoverTrigger } from "../Primitives/Popover";
import { Tooltip, TooltipContent, TooltipTrigger } from "../Primitives/Tooltip";

interface Props extends HTMLAttributes<HTMLButtonElement> {
  disablePopover?: boolean;
  disableTooltip?: boolean;
  data: AvatarRankConfig;
  characterId: number;
  disabled?: boolean;
}

const EidolonIcon = (props: Props) => {
  const { disabled = false, ...tooltipProps } = props;
  return (
    <Popover>
      <PopoverTrigger asChild>
        <IconWithTooltip
          {...tooltipProps}
          className={cn(
            "aspect-square min-w-[48px] invert dark:invert-0",
            disabled ? "brightness-[.25]" : ""
          )}
        />
      </PopoverTrigger>
      {!props.disablePopover && (
        <PopoverContent className="w-96">
          <span className="text-lg font-semibold text-accent-foreground">{props.data.name}</span>
          <EidolonDescription data={props.data} />
        </PopoverContent>
      )}
    </Popover>
  );
};

export { EidolonIcon };

const EidolonDescription = ({ data }: { data: AvatarRankConfig }) => {
  return (
    <div className="whitespace-pre-wrap">
      {data.desc.map((descPart, index) => (
        <Fragment key={index}>
          <span className="whitespace-pre-wrap">{sanitizeNewline(descPart)}</span>
          <span className="text-accent-foreground font-semibold">{data.param[index]}</span>
        </Fragment>
      ))}
    </div>
  );
};

interface IconProps extends Props, ButtonHTMLAttributes<HTMLButtonElement> {}
const IconWithTooltip = forwardRef<HTMLButtonElement, IconProps>(
  ({ disableTooltip = false, characterId, data, disablePopover, className, ...props }, ref) => {
    if (disableTooltip)
      return (
        <img
          src={url(characterId, data.rank)}
          alt={data.name}
          width={48}
          height={48}
          className="aspect-square min-w-[48px] invert dark:invert-0"
        />
      );
    return (
      <Tooltip disableHoverableContent>
        <TooltipTrigger asChild>
          <button ref={ref} className={className} {...props}>
            <img src={url(characterId, data.rank)} alt={data.name} width={48} height={48} />
          </button>
        </TooltipTrigger>
        <TooltipContent className="select-none">{data.name}</TooltipContent>
      </Tooltip>
    );
  }
);

function url(charID: number, eidolon: number) {
  let fmt = `rank${eidolon}`;
  if (eidolon == 3) fmt = "skill";
  if (eidolon == 5) fmt = "ultimate";
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/skill/${charID}_${fmt}.png`;
}
