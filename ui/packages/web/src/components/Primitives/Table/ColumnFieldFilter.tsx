import { Column } from "@tanstack/react-table";
import { Input } from "../Input";

interface Props<TData> {
  column: Column<TData> | undefined;
}
function ColumnFieldFilter<TData>({ column }: Props<TData>) {
  return (
    <Input
      placeholder="filter eventIndex"
      value={(column?.getFilterValue() as string) ?? ""}
      onChange={event => column?.setFilterValue(event.target.value)}
      className="max-w-sm w-64"
    />
  );
}
export { ColumnFieldFilter };
