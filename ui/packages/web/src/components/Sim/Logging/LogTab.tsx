import { useMutation } from "@tanstack/react-query";
import { Log } from "@/bindings/Log";
import { Button } from "@/components/Primitives/Button";
import { ENDPOINT, typedFetch } from "@/utils/constants";
import { columns } from "./columns";
import { DataTable } from "./LogTable";

const LogTab = () => {
  const logger = useMutation({
    mutationKey: [ENDPOINT.logMock],
    mutationFn: async () => await typedFetch<undefined, { list: Log[] }>(ENDPOINT.logMock),
    onSuccess: data => console.log(data),
  });

  return (
    <div className="bg-background">
      <Button onClick={() => logger.mutate()}>Generate Log</Button>
      <DataTable columns={columns} data={logger.data?.list ?? []} />
    </div>
  );
};

export { LogTab };
