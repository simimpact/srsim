import {
  ContextMenu,
  ContextMenuCheckboxItem,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuLabel,
  ContextMenuRadioGroup,
  ContextMenuRadioItem,
  ContextMenuSeparator,
  ContextMenuShortcut,
  ContextMenuSub,
  ContextMenuSubContent,
  ContextMenuSubTrigger,
  ContextMenuTrigger,
} from "@/components/Primitives/ContextMenu";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/Primitives/HoverCard";
import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { CharacterLineup } from "./CharacterLineup";
import { SimActionBar } from "./SimActionBar";

const Root = () => {
  return (
    <div className="flex h-full self-start grow">
      <div className="ml-4 self-center">
        <SimActionBar />
      </div>
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
        <ContextMenu>
          <ContextMenuTrigger className="bg-accent text-accent-foreground flex h-full mx-8 rounded-md p-10">
            <LogViewer placeholder="test" />
          </ContextMenuTrigger>
          <ContextMenuContent>
            <ContextMenuItem inset>
              Back
              <ContextMenuShortcut>⌘[</ContextMenuShortcut>
            </ContextMenuItem>
            <ContextMenuItem inset disabled>
              Forward
              <ContextMenuShortcut>⌘]</ContextMenuShortcut>
            </ContextMenuItem>
            <ContextMenuItem inset>
              Reload
              <ContextMenuShortcut>⌘R</ContextMenuShortcut>
            </ContextMenuItem>
            <ContextMenuSub>
              <ContextMenuSubTrigger inset>More Tools</ContextMenuSubTrigger>
              <ContextMenuSubContent className="w-48">
                <ContextMenuItem>
                  Save Page As...
                  <ContextMenuShortcut>⇧⌘S</ContextMenuShortcut>
                </ContextMenuItem>
                <ContextMenuItem>Create Shortcut...</ContextMenuItem>
                <ContextMenuItem>Name Window...</ContextMenuItem>
                <ContextMenuSeparator />
                <ContextMenuItem>Developer Tools</ContextMenuItem>
              </ContextMenuSubContent>
            </ContextMenuSub>
            <ContextMenuSeparator />
            <ContextMenuCheckboxItem checked>
              Show Bookmarks Bar
              <ContextMenuShortcut>⌘⇧B</ContextMenuShortcut>
            </ContextMenuCheckboxItem>
            <ContextMenuCheckboxItem>Show Full URLs</ContextMenuCheckboxItem>
            <ContextMenuSeparator />
            <ContextMenuRadioGroup value="pedro">
              <ContextMenuLabel inset>People</ContextMenuLabel>
              <ContextMenuSeparator />
              <ContextMenuRadioItem value="pedro">Pedro Duarte</ContextMenuRadioItem>
              <ContextMenuRadioItem value="colm">Colm Tuite</ContextMenuRadioItem>
            </ContextMenuRadioGroup>
          </ContextMenuContent>
        </ContextMenu>
      </div>
    </div>
  );
};
export { Root };
