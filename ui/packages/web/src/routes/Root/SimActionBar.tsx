import { cva } from "class-variance-authority";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarTrigger,
} from "@/components/Primitives/Menubar";
import { SimControlContext } from "@/providers/SimControl";
import { cn } from "@/utils/classname";

const SimActionBar = () => {
  const verticalHelper = cva("w-full justify-center", {
    variants: {
      variant: {
        run: "bg-green-500 data-[state=open]:bg-green-500/90 focus:bg-green-500/90",
        destructive: "bg-red-500 data-[state=open]:bg-red-500/90 focus:bg-red-500/90",
      },
    },
  });
  const { runSimulation, reset } = useContext(SimControlContext);

  function onRun() {
    runSimulation();
  }
  const route = useNavigate();

  return (
    <Menubar orientation="vertical" className="gap-2 min-w-max">
      <MenubarMenu>
        <MenubarTrigger className={verticalHelper({ variant: "run" })} onClick={onRun}>
          Run
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger className={verticalHelper({ variant: "destructive" })} onClick={reset}>
          Reset
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger className={cn(verticalHelper())} onClick={() => route("/config")}>
          Config
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger disabled className={cn(verticalHelper(), "cursor-not-allowed")}>
          APL
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>(don{"'"}t)click me</MenubarTrigger>
        <MenubarContent side="right">
          <MenubarItem>Import</MenubarItem>
          <MenubarItem>Export</MenubarItem>
          <MenubarSeparator />
          <MenubarItem>Share</MenubarItem>
        </MenubarContent>
      </MenubarMenu>
    </Menubar>
  );
};

export { SimActionBar };
