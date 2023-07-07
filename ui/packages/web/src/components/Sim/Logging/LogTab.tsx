import {
  ColumnFiltersState,
  Row,
  RowSelectionState,
  getCoreRowModel,
  getFilteredRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { LucideIcon } from "lucide-react";
import { useState } from "react";
import { Button } from "@/components/Primitives/Button";
import {
  ColumnSelectFilter,
  ColumnToggle,
  DataTable,
  MultipleSelect,
} from "@/components/Primitives/Table/index";
import { SimLog } from "@/utils/fetchLog";
import { columns } from "./columns";

interface Props {
  data: SimLog[];
}
const LogTab = ({ data }: Props) => {
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [rowSelection, setRowSelection] = useState<RowSelectionState>({});

  const table = useReactTable({
    // INFO: data and column data of the table
    data,
    columns: columns,

    // NOTE: all rows can be expanded, this might be the desired behaviour
    // since we want full JSON data on all events
    getRowCanExpand: () => true,
    getFilteredRowModel: getFilteredRowModel(),
    getCoreRowModel: getCoreRowModel(),
    // getPaginationRowModel: getPaginationRowModel(),

    // pass state to let the hook manage
    onColumnFiltersChange: setColumnFilters,
    onRowSelectionChange: setRowSelection,

    state: { columnFilters, rowSelection },
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
  const eventList = Array.from(new Set(data.map(event => event.name)));

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
            <MultipleSelect table={table} options={eventList} columnKey="name" />

            <Button
              onClick={() =>
                table.getSelectedRowModel().rows.forEach(row => row.toggleExpanded(true))
              }
            >
              Expand Selected
            </Button>
            <Button
              onClick={() =>
                table.getSelectedRowModel().rows.forEach(row => row.toggleExpanded(false))
              }
            >
              Collapse Selected
            </Button>
            <Button
              onClick={() => {
                table.toggleAllRowsExpanded(false);
                table.toggleAllRowsSelected(false);
              }}
            >
              Reset Selection & Expand
            </Button>

            <div className="grow" />
          </div>
          <div className="ml-auto">
            <ColumnToggle table={table} />
          </div>
        </div>

        <DataTable table={table} className="bg-background" renderSubComponent={ExpandComponent} />

        {/* <DataTablePagination table={table} /> */}
      </div>
    </>
  );
};

const ExpandComponent = ({ row }: { row: Row<SimLog> }) => {
  return (
    <pre className="text-[16px] whitespace-pre-wrap">
      <code>{JSON.stringify(row.original, null, 2)}</code>
    </pre>
  );
};

export { LogTab };
