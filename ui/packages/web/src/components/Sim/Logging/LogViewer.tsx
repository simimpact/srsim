import { useContext } from "react";
import { Button } from "@/components/Primitives/Button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { useTabRouteHelper } from "@/hooks/useTabRouteHelper";
import { SimControlContext } from "@/providers/SimControl";
import { LogTab } from "./LogTab";
// import { MvpTab } from "./MvpTab";
import { ResultTab } from "./ResultTab";

interface Props {
  placeholder?: string;
}
const LogViewer = ({ placeholder }: Props) => {
  if (placeholder) console.log(placeholder);
  const { tab, setTab } = useTabRouteHelper();

  const { simulationData, simulationResult, getResult } = useContext(SimControlContext);

  return (
    <div>
      <Tabs value={tab ?? "log"} onValueChange={setTab}>
        <TabsList className="h-full w-full">
          <TabsTrigger value="log" className="w-full">
            Logging/Debugging
          </TabsTrigger>
          <TabsTrigger value="result" className="w-full">
            Result tab
          </TabsTrigger>
          <TabsTrigger value="mvp" className="w-full">
            MVP tab
          </TabsTrigger>
        </TabsList>

        <TabsContent value="log">
          <LogTab data={simulationData} />
        </TabsContent>
        <TabsContent value="result">
          <Button size="sm" onClick={() => getResult()}>
            Get Results
          </Button>
          <ResultTab data={simulationResult} />
        </TabsContent>
        <TabsContent value="mvp">{/* <MvpTab name="test" /> */}</TabsContent>
      </Tabs>
    </div>
  );
};
export { LogViewer };
