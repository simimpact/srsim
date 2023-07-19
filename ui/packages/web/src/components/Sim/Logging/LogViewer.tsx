import { useContext } from "react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { SimControlContext } from "@/providers/SimControl";
import { LogTab } from "./LogTab";
import { MvpTab } from "./MvpTab";
import { ResultTab } from "./ResultTab";

interface Props {
  placeholder?: string;
}
const LogViewer = ({ placeholder }: Props) => {
  if (placeholder) console.log(placeholder);

  const { simulationData, simulationResult } = useContext(SimControlContext);

  return (
    <div>
      <Tabs defaultValue="log">
        <TabsList className="w-full h-full">
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
        <TabsContent value="mvp">
          <MvpTab name="test" />
        </TabsContent>
        <TabsContent value="log">
          <LogTab data={simulationData} />
        </TabsContent>
        <TabsContent value="result">
          <ResultTab data={simulationResult} />
        </TabsContent>
      </Tabs>
    </div>
  );
};
export { LogViewer };
