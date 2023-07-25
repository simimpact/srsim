import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { CharacterLineup } from "./CharacterLineup";
import { SimActionBar } from "./SimActionBar";

const Root = () => {
  return (
    <div id="dev" className="flex h-full grow self-start">
      <div className="flex grow flex-col gap-4">
        <div className="mx-8 flex justify-center gap-4">
          <CharacterLineup />
        </div>

        <div className="flex gap-4 px-4">
          <SimActionBar />
          <div className="bg-accent text-accent-foreground flex grow flex-col rounded-md p-10">
            <LogViewer />
          </div>
        </div>
      </div>
    </div>
  );
};
export { Root };
