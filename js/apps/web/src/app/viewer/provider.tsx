"use client";
import React, { createContext, ReactNode, useEffect, useReducer } from "react";
import { model } from "@srsim/ts-types";
import { produce } from "immer";

export interface State {
  data: model.SimResult | null;
  progress: number | null;
  recoveryConfig: string | null;
  error: string | null;
}

type SetResult = {
  type: "SET_RESULT";
  payload: {
    result: model.SimResult;
    progress: number;
  };
};

type SetError = {
  type: "SET_ERROR";
  payload: {
    error: string;
    config: string;
  };
};

type Clear = {
  type: "CLEAR";
};

export type Actions = SetResult | SetError | Clear;

export const initialState: State = {
  data: null,
  progress: null,
  recoveryConfig: null,
  error: null,
};

type ContextProviderProps = {
  children: ReactNode;
};

type ViewerContextType = {
  state: State;
  dispatch: React.Dispatch<Actions>;
};

export const ViewerContext = createContext<ViewerContextType>({
  state: initialState,
  dispatch: () => {},
});

export const ViewerProvider = ({ children }: ContextProviderProps) => {
  const [state, dispatch] = useReducer(
    produce((state: State, action: Actions) => {
      switch (action.type) {
        case "SET_RESULT":
          state.data = action.payload.result;
          state.progress = action.payload.progress;
          state.error = null;
          state.recoveryConfig = null;
          break;
        case "SET_ERROR":
          state.recoveryConfig = action.payload.config;
          state.error = action.payload.error;
          break;
        case "CLEAR":
          state.data = null;
          state.progress = null;
          state.recoveryConfig = null;
          state.error = null;
          break;
        default:
          break;
      }
    }),
    initialState
  );

  return <ViewerContext.Provider value={{ state, dispatch }}>{children}</ViewerContext.Provider>;
};
