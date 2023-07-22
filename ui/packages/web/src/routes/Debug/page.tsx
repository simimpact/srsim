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
    <div id="dev" className="w-full flex flex-col self-start grow gap-4">
      <div className="grow bg-accent text-accent-foreground flex flex-col rounded-md p-2 mx-2">
        <LogViewer />
      </div>
    </div>
  );
};

export { Debug };
