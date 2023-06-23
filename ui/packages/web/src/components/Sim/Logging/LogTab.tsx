import { useMutation } from "@tanstack/react-query";
import { Fragment } from "react";
import { Log } from "@/bindings/Log";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/Primitives/Accordion";
import { Button } from "@/components/Primitives/Button";
import { ScrollArea } from "@/components/Primitives/ScrollArea";
import { ENDPOINT, typedFetch } from "@/utils/constants";
import { columns } from "./columns";
import { DataTable } from "./LogTable";

const LogTab = () => {
  const logger = useMutation({
    mutationKey: [ENDPOINT.logMock],
    mutationFn: async () => await typedFetch<undefined, { list: Log[] }>(ENDPOINT.logMock),
    onSuccess: data => console.log(data),
  });
  // eslint-disable-next-line no-constant-condition, @typescript-eslint/no-unnecessary-condition
  if (1 === 2) {
    return (
      <>
        <Button onClick={() => logger.mutate()}>Generate Log</Button>
        {logger.data && (
          <ScrollArea className="h-96 p-4 rounded-md border bg-background">
            <Accordion type="single" collapsible>
              {logger.data.list.map((log, index) => (
                <Fragment key={index}>
                  <AccordionItem value={`log-${index}`}>
                    <AccordionTrigger>{log.fooo}</AccordionTrigger>
                    <AccordionContent>
                      <ul>
                        {Object.keys(log).map(key => (
                          <li key={key}>
                            {key}: {log[key] as string}
                          </li>
                        ))}
                      </ul>
                    </AccordionContent>
                  </AccordionItem>
                </Fragment>
              ))}
            </Accordion>
          </ScrollArea>
        )}
      </>
    );
  }
  return (
    <div className="bg-background">
      <Button onClick={() => logger.mutate()}>Generate Log</Button>
      <DataTable columns={columns} data={logger.data?.list ?? []} />
    </div>
  );
};

export { LogTab };
