import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { CharacterLineup } from "./CharacterLineup";
import { SimActionBar } from "./SimActionBar";

const Root = () => {
  return (
    <div id="dev" className="flex h-full self-start grow">
      <div className="flex flex-col grow gap-4">
        <div className="flex gap-4 justify-center mx-8">
          <CharacterLineup />
        </div>

        <div className="flex gap-4 px-4">
          <SimActionBar />
          <div className="grow bg-accent text-accent-foreground flex flex-col rounded-md p-10">
            <LogViewer />
          </div>
        </div>
      </div>
    </div>
  );
};
export { Root };
