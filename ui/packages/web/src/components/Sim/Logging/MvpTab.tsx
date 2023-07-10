import { useMutation } from "@tanstack/react-query";
import { ComponentProps } from "react";
import { MvpWrapper } from "@/bindings/MvpWrapper";
import { OuterLabelPie } from "@/components/Graphs";
import { Button } from "@/components/Primitives/Button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { ENDPOINT, typedFetch } from "@/utils/constants";
import { TeamXY } from "./TeamXY";

interface Props {
  name: string;
}
const MvpTab = ({ name }: Props) => {
  console.log(name);
  // TODO: mutation
  const statMock = useMutation({
    mutationKey: [ENDPOINT.statMock],
    // NOTE: NOT AN ACTUAL FN
    mutationFn: async () => await typedFetch<undefined, MvpWrapper>(ENDPOINT.statMock),
    onSuccess: data => console.log(data),
  });

  interface TestData {
    index: number;
    value: number;
    label: string;
  }
  const colors: string[] = [
    "#147EB3",
    "#29A634",
    "#D1980B",
    "#D33D17",
    "#9D3F9D",
    "#00A396",
    "#DB2C6F",
    "#8EB125",
    "#946638",
    "#7961DB",
  ];
  const pieMockData: ComponentProps<typeof OuterLabelPie<TestData>> = {
    data: [
      { label: "basic", index: 0, value: 0.03 },
      { label: "ult", index: 1, value: 0.2 },
      { label: "skill", index: 2, value: 0.69 },
      { label: "followup", index: 3, value: 0.08 },
    ],
    pieValue: d => d.value,
    width: 500,
    height: 500,
    color: d => colors[d.index],
    labelText: d => d.label,
    labelValue: d => d.value.toLocaleString("en", { style: "percent" }),
  };

  return (
    <>
      <Button onClick={() => statMock.mutate()}>Generate</Button>
      <div className="flex gap-2">
        <div id="left-container" className="flex flex-col gap-2 grow  max-w-[45vw]">
          <div id="portrait" className="bg-background rounded-md p-4 h-64">
            portrait
          </div>
          <div id="summary-distribution" className="bg-background rounded-md p-4 grow">
            {statMock.data && (
              <>
                <Tabs defaultValue="self">
                  <TabsList>
                    <TabsTrigger value="self">Self dist.</TabsTrigger>
                    <TabsTrigger value="occurence">Self dist. (count)</TabsTrigger>
                    <TabsTrigger value="team">team distribution</TabsTrigger>
                  </TabsList>
                  <TabsContent value="self">
                    <OuterLabelPie {...pieMockData} />
                  </TabsContent>
                  <TabsContent value="occurence">
                    <img
                      src="https://media.discordapp.net/attachments/1114188946721742898/1121109046964007003/image.png"
                      alt="example"
                    />
                  </TabsContent>
                  <TabsContent value="team">
                    gcsim{"'"}s {"cumulative contribution"} option
                    <br />
                    <a href="https://simimpact.app/sh/b2673849-b8ef-47ae-a4c0-90ec9582dce1#">
                      damage timeline(click)
                    </a>
                    <TeamXY data={statMock.data.team_distribution} />
                  </TabsContent>
                </Tabs>
              </>
            )}
          </div>
        </div>
        <div id="right-data-propagation" className="bg-background rounded-md p-4 grow">
          right: data propagation
        </div>
      </div>
    </>
  );
};
export { MvpTab };
