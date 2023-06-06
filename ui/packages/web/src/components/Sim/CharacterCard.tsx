import { VariantProps } from "class-variance-authority";
import { cn } from "@/utils/classname";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
  cardVariants,
} from "../Primitives/Card";

interface Props extends VariantProps<typeof cardVariants> {
  name: string;
  rarity: number;
}
const CharacterCard = ({ name, rarity, variant }: Props) => {
  return (
    <div className={cn(cardVariants({ variant }), "pl-4 border-0")}>
      <Card>
        <CardHeader>
          <div className="flex gap-4 justify-center">
            <CardTitle>{name}</CardTitle>
            <CardDescription>{rarity} *</CardDescription>
          </div>
        </CardHeader>
        <CardContent className="flex justify-center">
          <p className="font-bold">Disappear among the sea of butterflies</p>
        </CardContent>
        <CardFooter className="justify-center">
          <p>Illusion of the past</p>
        </CardFooter>
      </Card>
    </div>
  );
};
export { CharacterCard };
