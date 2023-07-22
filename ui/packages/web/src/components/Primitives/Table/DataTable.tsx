import { Row, Table as TableType, flexRender } from "@tanstack/react-table";
import { ForwardedRef, Fragment, HTMLAttributes, forwardRef } from "react";
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
  stickyHeader?: boolean;
  renderSubComponent: (props: { row: Row<TData> }) => React.ReactElement;
}

function DataTableInner<TData>(
  { table, renderSubComponent, stickyHeader = false, className, ...props }: Props<TData>,
  ref: ForwardedRef<HTMLDivElement>
) {
  return (
    <div
      ref={ref}
      className={cn("rounded-md border border-muted-foreground", className)}
      {...props}
    >
      <Table className="border-separate border-spacing-0">
        <TableHeader
          className={cn(stickyHeader ? "[&_th]:sticky [&_th]:top-0 [&_th]:bg-muted" : "")}
        >
          {table.getHeaderGroups().map(headerGroup => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map(header => {
                return (
                  <TableHead
                    key={header.id}
                    className="border-b-muted-foreground border-b"
                    style={{
                      width:
                        header.getSize() === Number.MAX_SAFE_INTEGER ? "auto" : header.getSize(),
                    }}
                  >
                    {header.isPlaceholder
                      ? null
                      : flexRender(header.column.columnDef.header, header.getContext())}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody className="[&_td]:border-b [&_td]:border-b-muted-foreground [&_tr:last-child_td]:border-0">
          {table.getRowModel().rows.length ? (
            table.getRowModel().rows.map(row => (
              <Fragment key={row.id}>
                <TableRow
                  data-state={row.getIsSelected() && "selected"}
                  // className="[&_td]:border-b [&_td]:border-b-muted-foreground"
                >
                  {row.getVisibleCells().map(cell => (
                    <TableCell
                      key={cell.id}
                      style={{
                        width:
                          cell.column.getSize() === Number.MAX_SAFE_INTEGER
                            ? "auto"
                            : cell.column.getSize(),
                      }}
                    >
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </TableCell>
                  ))}
                </TableRow>
                {row.getIsExpanded() && (
                  <TableCell colSpan={row.getVisibleCells().length}>
                    {renderSubComponent({ row })}
                  </TableCell>
                )}
              </Fragment>
            ))
          ) : (
            <TableRow className="[&_tr]:border-muted-foreground">
              <TableCell colSpan={table.getAllColumns().length} className="h-24 text-center">
                No results. Please run the simulation
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
export const DataTable = forwardRef(DataTableInner);
