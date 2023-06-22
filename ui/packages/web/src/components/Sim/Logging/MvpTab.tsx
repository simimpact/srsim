import { useMutation } from "@tanstack/react-query";
import { MvpWrapper } from "@/bindings/MvpWrapper";
import { Button } from "@/components/Primitives/Button";
import { ENDPOINT, typedFetch } from "@/utils/constants";
import { TeamXY } from "./TeamXY";

interface Props {
  name: string;
}
const MvpTab = ({ name }: Props) => {
  // TODO: mutation
  const statMock = useMutation({
    mutationKey: [ENDPOINT.statMock],
    // NOTE: NOT AN ACTUAL FN
    mutationFn: async () => await typedFetch<undefined, MvpWrapper>(ENDPOINT.statMock),
    onSuccess: data => console.log(data),
  });

  return (
    <>
      <Button onClick={() => statMock.mutate()}>Generate</Button>
      <div className="w-[95vw] h-[95vh] flex gap-2">
        <div id="left-container" className="flex flex-col gap-2 grow">
          <div id="portrait" className="bg-background rounded-md p-4 h-64">
            portrait
          </div>
          <div id="summary-distribution" className="bg-background rounded-md p-4 grow">
            {statMock.data && (
              <>
                <p>self distribution</p>
                <p>team distribution</p>
                <TeamXY data={statMock.data.team_distribution} />
              </>
            )}
          </div>
        </div>
        <div id="right-data-propagation" className="bg-background rounded-md p-4 grow">
          right: data propagation
        </div>
      </div>
    </>
  );
};
export { MvpTab };
