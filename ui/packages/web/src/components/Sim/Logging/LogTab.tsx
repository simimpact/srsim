import {
  ColumnFiltersState,
  PaginationState,
  Row,
  RowSelectionState,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { LucideIcon } from "lucide-react";
import { useState } from "react";
import { Button } from "@/components/Primitives/Button";
import {
  ColumnSelectFilter,
  ColumnToggle,
  DataTable,
  DataTablePagination,
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
  const [pagination, setPagination] = useState<PaginationState>({ pageIndex: 0, pageSize: 50 });

  const table = useReactTable({
    // INFO: data and column data of the table
    data,
    columns: columns,
    defaultColumn: {
      minSize: 0,
      size: Number.MAX_SAFE_INTEGER,
      maxSize: Number.MAX_SAFE_INTEGER,
    },

    // NOTE: all rows can be expanded, this might be the desired behaviour
    // since we want full JSON data on all events
    getRowCanExpand: () => true,
    getFilteredRowModel: getFilteredRowModel(),
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),

    // pass state to let the hook manage
    onColumnFiltersChange: setColumnFilters,
    onRowSelectionChange: setRowSelection,
    onPaginationChange: setPagination,

    state: { columnFilters, rowSelection, pagination },
  });

  const options: {
    label: string;
    value: SimLog["name"];
    icon?: LucideIcon;
  }[] = [
    { label: "HPChange", value: "HPChange" },
    { label: "LimboWaitHeal", value: "LimboWaitHeal" },
    { label: "EnergyChange", value: "EnergyChange" },
    { label: "StanceChange", value: "StanceChange" },
    { label: "StanceBreak", value: "StanceBreak" },
    { label: "StanceReset", value: "StanceReset" },
    { label: "SpChange", value: "SPChange" },
    { label: "AttackStart", value: "AttackStart" },
    { label: "AttackEnd", value: "AttackEnd" },
    { label: "HitStart", value: "HitStart" },
    { label: "HitEnd", value: "HitEnd" },
    { label: "HealStart", value: "HealStart" },
    { label: "HealEnd", value: "HealEnd" },
    { label: "ModifierAdded", value: "ModifierAdded" },
    { label: "ModifierResisted", value: "ModifierResisted" },
    { label: "ModifierRemoved", value: "ModifierRemoved" },
    { label: "ModifierExtendedDuration", value: "ModifierExtendedDuration" },
    { label: "ModifierExtendedCount", value: "ModifierExtendedCount" },
    { label: "ShieldAdded", value: "ShieldAdded" },
    { label: "ShieldRemoved", value: "ShieldRemoved" },
    { label: "ShieldChange", value: "ShieldChange" },
    { label: "Initialize", value: "Initialize" },
    { label: "CharactersAdded", value: "CharactersAdded" },
    { label: "EnemiesAdded", value: "EnemiesAdded" },
    { label: "BattleStart", value: "BattleStart" },
    { label: "Phase1Start", value: "Phase1Start" },
    { label: "Phase1End", value: "Phase1End" },
    { label: "Phase2Start", value: "Phase2Start" },
    { label: "Phase2End", value: "Phase2End" },
    { label: "TurnStart", value: "TurnStart" },
    { label: "TurnEnd", value: "TurnEnd" },
    { label: "Termination", value: "Termination" },
    { label: "ActionStart", value: "ActionStart" },
    { label: "ActionEnd", value: "ActionEnd" },
    { label: "InsertStart", value: "InsertStart" },
    { label: "InsertEnd", value: "InsertEnd" },
    { label: "TargetDeath", value: "TargetDeath" },
    { label: "TurnTargetsAdded", value: "TurnTargetsAdded" },
    { label: "TurnReset", value: "TurnReset" },
    { label: "GaugeChange", value: "GaugeChange" },
    { label: "CurrentGaugeCostChange", value: "CurrentGaugeCostChange" },
  ];
  const eventList = Array.from(new Set(data.map(event => event.name)));

  return (
    <>
      <div className="flex flex-col gap-2">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <ColumnToggle
              table={table}
              variant="outline"
              size="sm"
              className="border-muted-foreground h-8 lg:flex"
            />

            <ColumnSelectFilter
              placeholder="Select Event"
              options={options}
              column={table.getColumn("name")}
              buttonPlaceholder="Filter Event"
              variant="outline"
              className="border-muted-foreground"
              size="sm"
            />

            <MultipleSelect
              table={table}
              options={eventList}
              columnKey="name"
              size="sm"
              variant="outline"
              className="border-muted-foreground"
            />

            <Button
              size="sm"
              onClick={() =>
                table.getSelectedRowModel().rows.forEach(row => row.toggleExpanded(true))
              }
            >
              Expand Selected
            </Button>

            <Button
              size="sm"
              onClick={() =>
                table.getSelectedRowModel().rows.forEach(row => row.toggleExpanded(false))
              }
            >
              Collapse Selected
            </Button>
          </div>

          <Button
            size="sm"
            variant="destructive"
            onClick={() => {
              table.toggleAllRowsExpanded(false);
              table.toggleAllRowsSelected(false);
            }}
          >
            Reset Selection & Expand
          </Button>
        </div>

        <DataTable
          stickyHeader
          table={table}
          className="h-[1000px] overflow-auto"
          renderSubComponent={ExpandComponent}
        />

        <DataTablePagination table={table} rowOptions={[50, 75, 100, 150, 200]} />
      </div>
    </>
  );
};

const ExpandComponent = ({ row }: { row: Row<SimLog> }) => {
  return (
    <pre className="whitespace-pre-wrap text-[16px]">
      <code>{JSON.stringify(row.original, null, 2)}</code>
    </pre>
  );
};

export { LogTab };
