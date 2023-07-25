import { ButtonHTMLAttributes, Fragment, forwardRef } from "react";
import { AvatarRankConfig } from "@/bindings/AvatarRankConfig";
import { cn } from "@/utils/classname";
import { sanitizeNewline } from "@/utils/helpers";
import { Popover, PopoverContent, PopoverTrigger } from "../Primitives/Popover";
import { Tooltip, TooltipContent, TooltipTrigger } from "../Primitives/Tooltip";

interface Props {
  disablePopover?: boolean;
  disableTooltip?: boolean;
  data: AvatarRankConfig;
  characterId: number;
  disabled?: boolean;
}

const EidolonIcon = (props: Props) => {
  const { disabled = false, ...innerProps } = props;
  return (
    <Popover>
      <PopoverTrigger asChild>
        <IconWithTooltip
          {...innerProps}
          className={cn(
            "aspect-square min-w-[64px] invert dark:invert-0",
            disabled ? "brightness-[.25]" : ""
          )}
        />
      </PopoverTrigger>
      {!props.disablePopover && (
        <PopoverContent>
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
          width={64}
          height={64}
          className="aspect-square min-w-[64px] invert dark:invert-0"
        />
      );
    return (
      <Tooltip>
        <TooltipTrigger asChild>
          <button ref={ref} className={className} {...props}>
            <img src={url(characterId, data.rank)} alt={data.name} width={64} height={64} />
          </button>
        </TooltipTrigger>
        <TooltipContent>{data.name}</TooltipContent>
      </Tooltip>
    );
  }
);
IconWithTooltip.displayName = "IconWithTooltip";

function url(charID: number, eidolon: number) {
  let fmt = `rank${eidolon}`;
  if (eidolon == 3) fmt = "skill";
  if (eidolon == 5) fmt = "ultimate";
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/skill/${charID}_${fmt}.png`;
}
