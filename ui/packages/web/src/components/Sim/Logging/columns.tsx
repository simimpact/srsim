import * as Event from "@srsim/types/src/event";
import { createColumnHelper } from "@tanstack/react-table";
import { ChevronsDownUp, ChevronsUpDown, ExternalLink } from "lucide-react";
import { ReactNode } from "react";
import { Badge } from "@/components/Primitives/Badge";
import { Checkbox } from "@/components/Primitives/Checkbox";
import { Sheet, SheetContent, SheetTrigger } from "@/components/Primitives/Sheet";
import { Toggle } from "@/components/Primitives/Toggle";
import { SimLog } from "@/utils/fetchLog";

const columnHelper = createColumnHelper<SimLog>();

// https://ui.shadcn.com/docs/components/data-table
// https://tanstack.com/table/v8/docs/guide/column-defs
// https://ui.shadcn.com/docs/components/data-table#row-selection
export const columns = [
  columnHelper.display({
    id: "index",
    header: "#",
    cell: ({ row }) => row.index,
  }),
  columnHelper.display({
    id: "checkbox",
    header: ({ table }) => (
      <div className="flex items-center justify-center">
        <Checkbox
          checked={table.getIsAllPageRowsSelected()}
          onCheckedChange={value => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      </div>
    ),
    cell: ({ row }) => (
      <div className="text-center">
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={value => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      </div>
    ),
    enableSorting: false,
    enableHiding: false,
  }),
  columnHelper.display({
    id: "expander",
    header: () => null,
    cell: ({ row }) =>
      row.getCanExpand() && (
        <Toggle
          size="sm"
          // careful of the double callback
          onClick={row.getToggleExpandedHandler()}
        >
          {row.getIsExpanded() ? (
            <ChevronsDownUp className="w-4 h-4" />
          ) : (
            <ChevronsUpDown className="w-4 h-4" />
          )}
        </Toggle>
      ),
  }),
  columnHelper.accessor(data => data.name, {
    id: "name",
    filterFn: (row, _id, value: SimLog["name"][]) => {
      // NOTE: value is `any` by default, console log to double check type
      return value.includes(row.original.name);
    },
    cell: ({ row }) => (
      <Badge variant={row.getIsSelected() ? "destructive" : "default"}>
        {row.getValue("name")}
      </Badge>
    ),
  }),
  columnHelper.accessor(data => data, {
    id: "Important Key",
    cell: props => summarizeBy(props.getValue(), 0),
  }),
  columnHelper.accessor(data => data, {
    id: "event_key_1",
    cell: props => summarizeBy(props.getValue(), 1),
  }),
  columnHelper.accessor(data => data, {
    id: "event_key_2",
    cell: props => summarizeBy(props.getValue(), 2),
  }),
];

/**
 * this function takes in an event object and returns the appropriate element
 * that will be used to display in the table
 * @param data the log entry
 * @param tableIndex index of the column in the table (index of 0 means the
 * columns 2nd from the left, 1st column is the event name)
 * @returns Table cell
 */
function summarizeBy(data: SimLog, tableIndex: number): ReactNode {
  const { name, event } = data;

  function asDefault(index: number) {
    return (
      <Sheet>
        <SheetTrigger className="inline-flex items-center underline">
          {Object.keys(event)[index] && (
            <>
              {Object.keys(event)[index]} <ExternalLink className="ml-2 h-4 w-4" />
            </>
          )}
        </SheetTrigger>
        <SheetContent className="w-96 whitespace-pre-wrap text-muted-foreground overflow-y-auto">
          <p>{JSON.stringify(event[Object.keys(event)[index] as keyof typeof event], null, 4)}</p>
        </SheetContent>
      </Sheet>
    );
  }

  // here you can return a react component if you want something more complex
  // (e.g a dialog/popover/button etc for a table cell for big data)
  switch (name) {
    case "Initialize":
      return summarizeInitialize(event, tableIndex);
    case "SPChange":
    case "AttackStart":
    case "AttackEnd":
    case "HPChange":
    case "StanceChange":
    case "EnergyChange":
    case "HitEnd":
    case "StanceBreak":
    case "StanceReset":
    case "HealStart":
    case "HealEnd":
    case "InsertStart":
    case "InsertEnd":
    case "GaugeChange":
    case "CurrentGaugeCostChange":
      return [event.key, asDefault(1), asDefault(2)][tableIndex];
    case "HitStart":
      return [
        event.hit?.key ?? asDefault(0),
        `attacker: ${event.attacker}`,
        `defender: ${event.defender}`,
      ][tableIndex];
    case "ModifierAdded":
      return [event.modifier.name, `source: ${event.modifier.source}`, asDefault(2)][tableIndex];
    case "ModifierRemoved":
      return [event.modifier.name, asDefault(1), asDefault(2)][tableIndex];
    case "ActionEnd":
    case "ActionStart":
      return [event.attack_type, asDefault(1), asDefault(2)][tableIndex];
    default:
      return asDefault(tableIndex);
  }
}

function summarizeInitialize(event: Event.Initialize, tableIndex: number): ReactNode {
  if (tableIndex == 0) {
    return (
      <Sheet>
        <SheetTrigger>Config Schema</SheetTrigger>
        <SheetContent className="w-96 whitespace-pre-wrap text-muted-foreground overflow-y-auto">
          <p>{JSON.stringify(event.config, null, 4)}</p>
        </SheetContent>
      </Sheet>
    );
  } else if (tableIndex == 1) {
    return <span>Seed: {event.seed}</span>;
  }
  return <></>;
}
