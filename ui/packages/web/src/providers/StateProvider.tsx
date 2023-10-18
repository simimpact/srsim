import { Provider } from "jotai";
import { DevTools } from "jotai-devtools";
import { ReactNode } from "react";

interface Props {
  devTools?: boolean;
  children: ReactNode;
}
export function StateProvider({ children, devTools = false }: Props) {
  return (
    <Provider>
      {children}
      {devTools && <DevTools />}
    </Provider>
  );
}
