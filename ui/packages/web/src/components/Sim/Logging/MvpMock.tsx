import { useMutation } from "@tanstack/react-query";
import { Button } from "@/components/Primitives/Button";

interface Props {
  name: string;
}
const MvpMock = ({ name }: Props) => {
  // TODO: mutation
  const { mutate } = useMutation({
    mutationKey: ["aaaarst"],
    // NOTE: NOT AN ACTUAL FN
    mutationFn: async (data: string) => await fetch("url", { body: data }),
  });

  return (
    <div className="flex">
      <div>{name}</div>
      <Button onClick={() => mutate(name)}>Button</Button>
    </div>
  );
};
export { MvpMock };
