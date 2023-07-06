import { useMutation } from "@tanstack/react-query";
import { Button } from "@/components/Primitives/Button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { ENDPOINT } from "@/utils/constants";
import { fetchLog } from "@/utils/fetchLog";
import { LogTab } from "./LogTab";
import { MvpTab } from "./MvpTab";

interface Props {
  placeholder: string;
}
const LogViewer = ({ placeholder }: Props) => {
  console.log(placeholder);
  const logger = useMutation({
    mutationKey: [ENDPOINT.logMock],
    mutationFn: async () => await fetchLog(),
    onSuccess: data => console.log(data),
  });

  return (
    <div>
      <Button className="my-4" onClick={() => logger.mutate()}>
        Generate Log (Make sure your CLI is running)
      </Button>
      <Tabs defaultValue="mvp">
        <TabsList className="w-full h-full">
          <TabsTrigger value="mvp" className="w-full">
            MVP tab
          </TabsTrigger>
          <TabsTrigger value="log" className="w-full">
            Logging/Debugging
          </TabsTrigger>
        </TabsList>
        <TabsContent value="mvp">
          <MvpTab name="test" />
        </TabsContent>
        <TabsContent value="log">
          <LogTab data={logger.data ?? []} />
        </TabsContent>
      </Tabs>
    </div>
  );
};
export { LogViewer };
