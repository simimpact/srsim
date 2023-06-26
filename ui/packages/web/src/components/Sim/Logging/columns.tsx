import { ColumnDef } from "@tanstack/react-table";
import { Log } from "@/bindings/Log";
import { Button } from "@/components/Primitives/Button";
import { cn } from "@/utils/classname";

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
    header: "eventIndex",
    cell: ({ row }) => (
      <div style={{ paddingLeft: `${row.depth * 2}rem` }}>
        {row.getCanExpand() && (
          <Button {...{ onClick: row.getToggleExpandedHandler() }} variant={"outline"} size={"sm"}>
            {row.getIsExpanded() ? "v" : ">"}
          </Button>
        )}
        {row.getValue("eventIndex")}
      </div>
    ),
    filterFn: (row, id, value: number) => {
      return String(row.getValue(id)).includes(String(value));
    },
  },
  {
    accessorKey: "eventName",
    header: "eventName",
    cell: ({ row }) => (
      <div className={cn(row.getValue("eventName") == "SPChange" ? "bg-red-400" : "bg-green-400")}>
        {row.getValue("eventName")}
      </div>
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
