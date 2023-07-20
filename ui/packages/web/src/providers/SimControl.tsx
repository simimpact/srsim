import { useMutation } from "@tanstack/react-query";
import { ReactNode, createContext, useState } from "react";
import { ENDPOINT } from "@/utils/constants";
import { SimLog, SimResult, fetchLog, fetchResult } from "@/utils/fetchLog";

interface SimControlContextPayload {
  runSimulation: () => void;
  simulationData: SimLog[];

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

  const logMutation = useMutation({
    mutationKey: [ENDPOINT.logMock],
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

    getResult,
    simulationResult: resultMutation.data,

    reset,
  };
}

export const defaultSimControl: SimControlContextPayload = {
  runSimulation: () => {},
  simulationData: [],

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
