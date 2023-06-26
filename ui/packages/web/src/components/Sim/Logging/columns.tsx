import { ChevronDownIcon, ChevronRightIcon } from "@radix-ui/react-icons";
import { ColumnDef } from "@tanstack/react-table";
import { Log } from "@/bindings/Log";
import { Badge } from "@/components/Primitives/Badge";
import { Toggle } from "@/components/Primitives/Toggle";

// NOTE: halt, needs backend data before cooking,
// table might still not be the best for skimming through logs
// likely plausible with extra styling with minimal padding/spacing + make
// rows work as text lines
// using filter, no pagination
// https://ui.shadcn.com/docs/components/data-table
// https://tanstack.com/table/v8/docs/guide/column-defs
export const columns: ColumnDef<Log>[] = [
  {
    accessorKey: "eventIndex",
    header: "Event Index",
    cell: ({ row }) => (
      <div style={{ paddingLeft: `${row.depth * 2}rem` }} className="flex items-center gap-2">
        <Toggle
          {...{ onClick: row.getToggleExpandedHandler() }}
          variant={"outline"}
          size={"sm"}
          disabled={!row.getCanExpand()}
        >
          {row.getIsExpanded() ? <ChevronDownIcon /> : <ChevronRightIcon />}
        </Toggle>
        {row.getValue("eventIndex")}
      </div>
    ),
    filterFn: (row, id, value: number) => {
      return String(row.getValue(id)).includes(String(value));
    },
  },
  {
    accessorKey: "eventName",
    header: "Event Name",
    cell: ({ row }) => (
      <Badge variant={row.getValue("eventName") === "SPChange" ? "destructive" : "default"}>
        {row.getValue("eventName")}
      </Badge>
    ),
    filterFn: (row, columnId, filterValue: string) => {
      return filterValue.includes(row.getValue(columnId));
    },
  },
  { accessorKey: "bar", header: "Bar" },
  { accessorKey: "bazz", header: "Bazz" },
  { accessorKey: "fooo", header: "Foo" },
  { accessorKey: "abc", header: "Column1" },
  { accessorKey: "sss", header: "Column2" },
];
