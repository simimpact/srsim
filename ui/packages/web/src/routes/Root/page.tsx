import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/Primitives/HoverCard";
import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { CharacterLineup } from "./CharacterLineup";
import { SimActionBar } from "./SimActionBar";

const Root = () => {
  return (
    <div id="dev" className="flex h-full self-start grow">
      <div className="flex flex-col grow gap-4">
        <div className="flex gap-4 justify-center mx-8">
          <div>
            <CharacterLineup />
          </div>
          <HoverCard openDelay={300}>
            <HoverCardTrigger>
              <CharacterLineup isEnemy />
            </HoverCardTrigger>
            <HoverCardContent className="flex flex-col gap-2 duration-500">
              <CharacterLineup isEnemy header="wave1 (click on the icons)" />
              <CharacterLineup isEnemy header="wave2 (click on the icons)" />
              <CharacterLineup isEnemy header="wave3 (click on the icons)" />
            </HoverCardContent>
          </HoverCard>
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
