import { useMutation } from "@tanstack/react-query";
import { Character } from "@/bindings/MihoResponse";
import { Button } from "@/components/Primitives/Button";
import { ENDPOINT, typedFetch } from "@/utils/constants";

interface Props {
  name: string;
}
const MvpTab = ({ name }: Props) => {
  // TODO: mutation
  const statMock = useMutation({
    mutationKey: [ENDPOINT.statMock],
    // NOTE: NOT AN ACTUAL FN
    mutationFn: async () => await typedFetch<undefined, Character>(ENDPOINT.statMock),
  });

  return (
    <div className="w-[70vw] h-[70vh]">
      <Button onClick={() => statMock.mutate()}>Button</Button>
    </div>
  );
};
export { MvpTab };
