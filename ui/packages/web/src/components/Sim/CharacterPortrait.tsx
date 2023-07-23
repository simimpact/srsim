import { AvatarConfig } from "@/routes/Root/CharacterLineup";
import { cn } from "@/utils/classname";
import { elementVariants, rarityVariants } from "@/utils/variants";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../Primitives/Dialog";

interface Props {
  data: AvatarConfig;
}

const CharacterPortrait = ({ data }: Props) => {
  const { damage_type: element, avatar_base_type: path } = data;

  return (
    <Dialog>
      <DialogTrigger className="flex items-center">
        <img
          src={url(data.avatar_id)}
          alt={data.avatar_name}
          className={cn(
            "max-h-12 rounded-full box-content p-1 border-2",
            elementVariants({ border: element }),
            rarityVariants({ rarity: data.rarity as 1 | 2 | 3 | 4 | 5 | null })
          )}
        />
      </DialogTrigger>
      <DialogContent className="text-foreground">
        <DialogHeader>
          <DialogTitle className="text-foreground">{data.avatar_name}</DialogTitle>
          <DialogDescription>{`${data.rarity} ✦ ${path}`}</DialogDescription>
        </DialogHeader>
        <div className="flex gap-2.5">
          <img
            src={url(data.avatar_id)}
            alt={data.avatar_name}
            className={cn(
              "h-32 w-32 rounded-full box-content p-1 border-2",
              elementVariants({ border: element }),
              rarityVariants({ rarity: data.rarity as 1 | 2 | 3 | 4 | 5 | null })
            )}
          />
          <p>
            Id: {data.avatar_id} <br />
            Path: {path} <br />
            Element: {element}
          </p>
        </div>
      </DialogContent>
    </Dialog>
  );
};

function url(characterCode: number): string {
  return `https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/icon/character/${characterCode}.png`;
}
export { CharacterPortrait };
