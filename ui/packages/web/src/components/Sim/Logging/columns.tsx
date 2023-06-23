import { ColumnDef } from "@tanstack/react-table";
import { Log } from "@/bindings/Log";

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
  },
  {
    accessorKey: "eventName",
    header: "eventName",
  },
  {
    accessorKey: "bar",
    header: "Bar",
  },
  {
    accessorKey: "bazz",
    header: "Bazz",
  },
  {
    accessorKey: "foo",
    header: "Foo",
  },
];
