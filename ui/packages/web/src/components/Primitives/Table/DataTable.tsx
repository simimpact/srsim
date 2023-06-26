import { Table as TableType, flexRender } from "@tanstack/react-table";
import { ForwardedRef, HTMLAttributes, forwardRef } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/Primitives/Table";
import { cn } from "@/utils/classname";

// Redecalare forwardRef to accept generic types
// INFO: https://fettblog.eu/typescript-react-generic-forward-refs/
declare module "react" {
  function forwardRef<T, P = NonNullable<unknown>>(
    render: (props: P, ref: React.Ref<T>) => React.ReactElement | null
  ): (props: P & React.RefAttributes<T>) => React.ReactElement | null;
}

interface Props<TData> extends HTMLAttributes<HTMLDivElement> {
  table: TableType<TData>;
}

function DataTableInner<TData>(
  { table, className, ...props }: Props<TData>,
  ref: ForwardedRef<HTMLDivElement>
) {
  return (
    <div ref={ref} className={cn("rounded-md border", className)} {...props}>
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map(headerGroup => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map(header => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(header.column.columnDef.header, header.getContext())}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.length ? (
            table.getRowModel().rows.map(row => (
              <TableRow key={row.id} data-state={row.getIsSelected() && "selected"}>
                {row.getVisibleCells().map(cell => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={table.getAllColumns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
export const DataTable = forwardRef(DataTableInner);
