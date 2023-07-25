import { event } from "@srsim/types";
import { useMutation } from "@tanstack/react-query";
import { ReactNode, createContext, useState } from "react";
import { SimLog, SimResult, fetchLog, fetchResult } from "@/utils/fetchLog";
import { SimConfig } from "./temporarySimControlTypes";

interface SimControlContextPayload {
  runSimulation: () => void;
  simulationData: SimLog[];
  simulationConfig: SimConfig | undefined;

  getResult: () => void;
  simulationResult: SimResult | undefined;
  reset: () => void;
}

/**
 * initializes an instance of SimControl, this should not be called more than once other than in AppProvider.tsx
 *
 * Inside components use useContext to access the data
 * @returns
 */
function useSimControl(): SimControlContextPayload {
  const [simulationData, setSimulationData] = useState<SimLog[]>([]);
  const [simulationConfig, setSimulationConfig] = useState<SimConfig | undefined>(undefined);

  const logMutation = useMutation({
    mutationKey: ["simulation"],
    mutationFn: async () => await fetchLog(),
    onSuccess: onLogMutate,
  });

  const resultMutation = useMutation({
    mutationKey: ["result"],
    mutationFn: async () => await fetchResult(),
  });

  function onLogMutate(data: SimLog[]) {
    console.log(data);
    setSimulationData(data);

    const config = data.find(e => e.name == "Initialize");
    if (config) {
      // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
      const { event }: { event: { config: SimConfig } } = config as event.Initialize["config"];
      setSimulationConfig(event.config);
    }
  }

  function runSimulation() {
    logMutation.mutate();
  }

  function getResult() {
    resultMutation.mutate();
  }

  function reset() {
    setSimulationData([]);
  }

  return {
    runSimulation,
    simulationData,
    simulationConfig,

    getResult,
    simulationResult: resultMutation.data,

    reset,
  };
}

export const defaultSimControl: SimControlContextPayload = {
  runSimulation: () => {},
  simulationData: [],
  simulationConfig: undefined,

  getResult: () => {},
  simulationResult: undefined,

  reset: () => {},
};

export const SimControlContext = createContext<SimControlContextPayload>(defaultSimControl);

/**
 * wrapper provider so we can use react query
 */
export const SimControl = ({ children }: { children: ReactNode }) => {
  const simControl = useSimControl();
  return <SimControlContext.Provider value={simControl}>{children}</SimControlContext.Provider>;
};
