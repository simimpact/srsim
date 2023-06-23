import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { LogTab } from "./LogTab";
import { MvpTab } from "./MvpTab";

interface Props {
  placeholder: string;
}
const LogViewer = ({ placeholder }: Props) => {
  return (
    <div className="w-[95vw] h-[95vh]">
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
          <LogTab />
        </TabsContent>
      </Tabs>
    </div>
  );
};
export { LogViewer };
