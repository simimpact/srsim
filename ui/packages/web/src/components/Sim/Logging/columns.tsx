import { createColumnHelper } from "@tanstack/react-table";
import { ChevronsDownUp, ChevronsUpDown } from "lucide-react";
import { Badge } from "@/components/Primitives/Badge";
import { Checkbox } from "@/components/Primitives/Checkbox";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/Primitives/Popover";
import { Toggle } from "@/components/Primitives/Toggle";
import { SimLog } from "@/utils/fetchLog";

const columnHelper = createColumnHelper<SimLog>();

// https://ui.shadcn.com/docs/components/data-table
// https://tanstack.com/table/v8/docs/guide/column-defs
// TODO: add row selection, letting user mark which kind of events they want
// highlighted
// https://ui.shadcn.com/docs/components/data-table#row-selection
export const columns = [
  columnHelper.display({
    id: "index",
    cell: ({ row }) => row.index,
  }),
  columnHelper.display({
    id: "checkbox",
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={value => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={value => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
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
          {row.getIsExpanded() ? <ChevronsDownUp /> : <ChevronsUpDown />}
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
      <Badge variant={row.original.name === "TurnStart" ? "destructive" : "default"}>
        {row.getValue("name")}
      </Badge>
    ),
  }),
  columnHelper.accessor(data => data, {
    id: "event_0",
    cell: props => summarizeBy(props.getValue(), 0),
  }),
  columnHelper.accessor(data => data, {
    id: "event_1",
    cell: props => summarizeBy(props.getValue(), 1),
  }),
  columnHelper.accessor(data => data, {
    id: "event_2",
    cell: props => summarizeBy(props.getValue(), 2),
  }),
  columnHelper.accessor(data => data, {
    id: "event_3",
    cell: props => summarizeBy(props.getValue(), 3),
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
const summarizeBy = (data: SimLog, tableIndex: number): JSX.Element => {
  const { name, event } = data;
  switch (name) {
    case "Initialize": {
      if (tableIndex == 0) {
        return (
          <Popover>
            <PopoverTrigger>Config Schema</PopoverTrigger>
            <PopoverContent className="w-96 whitespace-pre-wrap">
              <p>{JSON.stringify(event.config, null, 4)}</p>
            </PopoverContent>
          </Popover>
        );
      } else if (tableIndex == 1) {
        return <span>Seed: {event.seed}</span>;
      }
      return <></>;
    }
    default:
      return (
        <Popover>
          <PopoverTrigger>{Object.keys(event)[tableIndex]}</PopoverTrigger>
          <PopoverContent className="w-96 whitespace-pre-wrap">
            <p>
              {JSON.stringify(event[Object.keys(event)[tableIndex] as keyof typeof event], null, 4)}
            </p>
          </PopoverContent>
        </Popover>
      );
  }
};
