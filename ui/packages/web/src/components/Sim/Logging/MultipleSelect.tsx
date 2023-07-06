import { Table } from "@tanstack/react-table";
import { Check } from "lucide-react";
import { Button } from "@/components/Primitives/Button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "@/components/Primitives/Command";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/Primitives/Popover";
import { cn } from "@/utils/classname";
import { SimLog } from "@/utils/fetchLog";

// TODO: move event mapping to parent, making this component generic
interface MultipleSelectProps {
  data: SimLog[];
  table: Table<SimLog>;
}
const MultipleSelect = ({ data, table }: MultipleSelectProps) => {
  // removes duplications
  const options = Array.from(new Set(data.map(event => event.name)));
  const selectedEvents = table.getIsSomePageRowsSelected();

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button size="sm">Select Events</Button>
      </PopoverTrigger>
      <PopoverContent>
        <Command>
          <CommandInput placeholder="Enter event" className="border-none" />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup>
              {options.map(eventName => {
                const { rows } = table.getRowModel();

                const rowsWithEventName = rows.filter(row => row.getValue("name") == eventName);

                const isSelected = rowsWithEventName.every(row => row.getIsSelected());

                return (
                  <CommandItem
                    key={eventName}
                    onSelect={() => {
                      const filtered = rows.filter(row => row.getValue("name") == eventName);
                      if (filtered.every(row => row.getIsSelected())) {
                        filtered.forEach(row => row.toggleSelected(false));
                      } else filtered.forEach(row => row.toggleSelected(true));
                    }}
                  >
                    <div
                      className={cn(
                        "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
                        isSelected
                          ? "bg-primary text-primary-foreground"
                          : "opacity-50 [&_svg]:invisible"
                      )}
                    >
                      <Check className={cn("h-4 w-4")} />
                    </div>
                    {/* {option.icon && <option.icon className="mr-2 h-4 w-4 text-muted-foreground" />} */}
                    <span>{eventName}</span>
                    {/* {facets?.get(option.value) && (
                      <span className="ml-auto flex h-4 w-4 items-center justify-center font-mono text-xs">
                        {facets.get(option.value)}
                      </span>
                    )} */}
                  </CommandItem>
                );
              })}
            </CommandGroup>
            {selectedEvents && (
              <>
                <CommandSeparator />
                <CommandGroup>
                  <CommandItem
                    onSelect={() => table.toggleAllPageRowsSelected(false)}
                    className="justify-center text-center"
                  >
                    Clear selection
                  </CommandItem>
                </CommandGroup>
              </>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
};
export { MultipleSelect };
