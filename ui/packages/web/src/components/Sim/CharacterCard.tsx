import { VariantProps } from "class-variance-authority";
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
  // this will eventually be expanded when character has more info like id,
  // rarity etc
  name: string;
  rarity: number;
}
const CharacterCard = ({ name, variant }: Props) => {
  return (
    <div className="flex">
      {/* TODO: inject this with the variant */}
      <div className="bg-fire pl-4 rounded-lg">
        <Card>
          <CardHeader>
            <CardTitle>{name}</CardTitle>
            <CardDescription>Card Description</CardDescription>
          </CardHeader>
          <CardContent>
            <p>Card Content</p>
          </CardContent>
          <CardFooter>
            <p>Card Footer</p>
          </CardFooter>
        </Card>
      </div>
    </div>
  );
};
export { CharacterCard };
