import { Column } from "@tanstack/react-table";
import { Input } from "../Input";

interface Props<TData> {
  column: Column<TData> | undefined;
}
function ColumnFieldFilter<TData>({ column }: Props<TData>) {
  return (
    <Input
      placeholder="filter eventIndex"
      // eslint-disable-next-line @typescript-eslint/no-unnecessary-condition
      value={(column?.getFilterValue() as string) ?? ""}
      onChange={event => column?.setFilterValue(event.target.value)}
      className="w-64 max-w-sm"
    />
  );
}
export { ColumnFieldFilter };
