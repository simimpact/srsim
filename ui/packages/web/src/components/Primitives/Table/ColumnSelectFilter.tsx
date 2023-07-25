import { Column } from "@tanstack/react-table";
import { Check, LucideIcon } from "lucide-react";
import { ComponentPropsWithoutRef } from "react";
import { cn } from "@/utils/classname";
import { Button } from "../Button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "../Command";
import { Popover, PopoverContent, PopoverTrigger } from "../Popover";

interface OptionItem {
  label: string;
  value: string;
  icon?: LucideIcon;
}
interface Props<TData> extends ComponentPropsWithoutRef<typeof Button> {
  placeholder: string;
  buttonPlaceholder?: string;
  options: OptionItem[];
  column: Column<TData> | undefined;
}

function ColumnSelectFilter<TData>({
  column,
  options,
  placeholder,
  buttonPlaceholder,
  ...props
}: Props<TData>) {
  const selectedValues = new Set(column?.getFilterValue() as string[]);
  const facets = column?.getFacetedUniqueValues();

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button {...props}>{buttonPlaceholder}</Button>
      </PopoverTrigger>
      <PopoverContent>
        <Command>
          <CommandInput placeholder={placeholder} className="border-none" />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup>
              {options.map(option => {
                const isSelected = selectedValues.has(option.value);
                return (
                  <CommandItem
                    key={option.value}
                    onSelect={() => {
                      if (isSelected) selectedValues.delete(option.value);
                      else selectedValues.add(option.value);

                      const filterValues = Array.from(selectedValues);
                      column?.setFilterValue(filterValues.length ? filterValues : undefined);
                    }}
                  >
                    <div
                      className={cn(
                        "border-primary mr-2 flex h-4 w-4 items-center justify-center rounded-sm border",
                        isSelected
                          ? "bg-primary text-primary-foreground"
                          : "opacity-50 [&_svg]:invisible"
                      )}
                    >
                      <Check className={cn("h-4 w-4")} />
                    </div>
                    {option.icon && <option.icon className="text-muted-foreground mr-2 h-4 w-4" />}
                    <span>{option.label}</span>
                    {facets?.get(option.value) && (
                      <span className="ml-auto flex h-4 w-4 items-center justify-center font-mono text-xs">
                        {facets.get(option.value)}
                      </span>
                    )}
                  </CommandItem>
                );
              })}
            </CommandGroup>
            {selectedValues.size > 0 && (
              <>
                <CommandSeparator />
                <CommandGroup>
                  <CommandItem
                    onSelect={() => column?.setFilterValue(undefined)}
                    className="justify-center text-center"
                  >
                    Clear filters
                  </CommandItem>
                </CommandGroup>
              </>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
export { ColumnSelectFilter };
