import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarTrigger,
} from "@/components/Primitives/Menubar";

const SimActionBar = () => {
  return (
    <Menubar orientation="vertical" className="gap-4">
      <MenubarMenu>
        <MenubarTrigger className="bg-green-500 hover:bg-green-500/90">Run</MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger className="bg-red-500 hover:bg-red-500/90">Debug</MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>APL</MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>click me</MenubarTrigger>
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
