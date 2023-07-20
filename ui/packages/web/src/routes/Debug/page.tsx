import { useContext, useEffect } from "react";
import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { SimControlContext } from "@/providers/SimControl";

const Debug = () => {
  const { runSimulation } = useContext(SimControlContext);

  useEffect(() => {
    runSimulation();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div id="dev" className="flex h-full self-start grow">
      <div className="flex flex-col grow gap-4">
        <div className="flex gap-4 px-4">
          <div className="grow bg-accent text-accent-foreground flex flex-col rounded-md p-10">
            <LogViewer />
          </div>
        </div>
      </div>
    </div>
  );
};

export { Debug };
