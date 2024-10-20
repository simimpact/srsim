'use client';
import React, {createContext, ReactNode, useEffect, useReducer} from 'react';
import {model} from '@srsim/ts-types';
import {produce} from 'immer';

export interface State {
  data: model.SimResult | null;
  recoveryConfig: string | null;
  error: string | null;
}

type SetResult = {
  type: 'SET_RESULT';
  payload: model.SimResult;
};

type SetError = {
  type: 'SET_ERROR';
  payload: {
    error: string;
    config: string;
  };
};

type Clear = {
  type: 'CLEAR';
};

export type Actions = SetResult | SetError | Clear;

export const initialState: State = {
  data: null,
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

export const ViewerProvider = ({children}: ContextProviderProps) => {
  const [state, dispatch] = useReducer(
    produce((state: State, action: Actions) => {
      switch (action.type) {
        case 'SET_RESULT':
          state.data = action.payload;
          state.error = null;
          state.recoveryConfig = null;
          break;
        case 'SET_ERROR':
          state.recoveryConfig = action.payload.config;
          state.error = action.payload.error;
          break;
        case 'CLEAR':
          state.data = null;
          state.recoveryConfig = null;
          state.error = null;
          break;
        default:
          break;
      }
    }),
    initialState,
  );

  return (
    <ViewerContext.Provider value={{state, dispatch}}>
      {children}
    </ViewerContext.Provider>
  );
};
