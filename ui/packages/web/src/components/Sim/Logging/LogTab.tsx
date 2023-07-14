import {
  ColumnFiltersState,
  getCoreRowModel,
  getFilteredRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { LucideIcon } from "lucide-react";
import { useState } from "react";
import { ColumnSelectFilter, ColumnToggle, DataTable } from "@/components/Primitives/Table/index";
import { SimLog } from "@/utils/fetchLog";
import { columns } from "./columns";

interface Props {
  data: SimLog[];
}
const LogTab = ({ data }: Props) => {
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);

  const table = useReactTable({
    data,
    columns: columns,

    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),

    getCoreRowModel: getCoreRowModel(),

    // getPaginationRowModel: getPaginationRowModel(),

    state: { columnFilters },
  });

  const options: {
    label: string;
    value: SimLog["name"];
    icon?: LucideIcon;
  }[] = [
    { label: "Turn End", value: "TurnEnd" },
    { label: "Turn Reset", value: "TurnReset" },
    { label: "Battle Start", value: "BattleStart" },
  ];

  return (
    <>
      <div className="flex flex-col gap-4">
        <div className="flex items-center">
          <div className="flex items-center gap-4">
            <ColumnSelectFilter
              placeholder="Select Event"
              options={options}
              column={table.getColumn("name")}
              buttonPlaceholder="Filter Event"
            />

            <div className="grow" />
          </div>
          <div className="ml-auto">
            <ColumnToggle table={table} />
          </div>
        </div>

        <DataTable table={table} className="bg-background" />

        {/* <DataTablePagination table={table} /> */}
      </div>
    </>
  );
};

export { LogTab };
