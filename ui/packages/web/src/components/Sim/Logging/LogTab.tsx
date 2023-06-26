import { useMutation } from "@tanstack/react-query";
import { LucideIcon } from "lucide-react";
import { Log } from "@/bindings/Log";
import { Button } from "@/components/Primitives/Button";
import {
  ColumnFieldFilter,
  ColumnSelectFilter,
  ColumnToggle,
  DataTable,
  useTable,
} from "@/components/Primitives/Table/index";
import { ENDPOINT, typedFetch } from "@/utils/constants";
import { columns } from "./columns";

const LogTab = () => {
  const logger = useMutation({
    mutationKey: [ENDPOINT.logMock],
    mutationFn: async () => await typedFetch<undefined, { list: Log[] }>(ENDPOINT.logMock),
    onSuccess: data => console.log(data),
  });

  const { table } = useTable<Log>({
    columns,
    data: logger.data?.list ?? [],
    childKey: "children",
  });

  const options: {
    label: string;
    value: string;
    icon?: LucideIcon;
  }[] = [
    { label: "Turn End", value: "TurnEnd" },
    { label: "Turn Reset", value: "TurnReset" },
    { label: "SP Change", value: "SPChange" },
  ];

  return (
    <>
      <Button onClick={() => logger.mutate()}>Generate Log</Button>
      <div className="flex items-center">
        <div className="flex items-center py-4 gap-4">
          <ColumnFieldFilter column={table.getColumn("eventIndex")} />

          <ColumnSelectFilter
            placeholder={"Select state"}
            options={options}
            column={table.getColumn("eventName")}
            buttonPlaceholder="Event Name"
          />

          <div className="grow" />
        </div>
        <div className="ml-auto">
          <ColumnToggle table={table} />
        </div>
      </div>
      <DataTable table={table} />
    </>
  );
};

export { LogTab };
