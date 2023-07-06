import {
  ColumnDef,
  ColumnFiltersState,
  ExpandedState,
  getCoreRowModel,
  getExpandedRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { useState } from "react";

type KeysOfType<T, V> = keyof {
  [P in keyof T as T[P] extends V ? P : never]: any;
};
interface Props<TData> {
  columns: ColumnDef<TData, string>[]; // TODO: check if string is correct
  data: TData[];
  childKey?: KeysOfType<TData, any[]>;
}

/**
 * Wrapper around tanstack's useReactTable() hook
 */
function useTable<TData>({ columns, data, childKey }: Props<TData>) {
  const [expanded, setExpanded] = useState<ExpandedState>({});
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);

  const table = useReactTable({
    data,
    columns,
    // WARN: untested against cases that's not TData[]
    getSubRows: !childKey ? undefined : row => row[childKey] as TData[],

    onExpandedChange: setExpanded,
    getExpandedRowModel: getExpandedRowModel(),

    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),

    getCoreRowModel: getCoreRowModel(),

    getPaginationRowModel: getPaginationRowModel(),

    state: { expanded, columnFilters },
  });
  return { table };
}

export { useTable };
