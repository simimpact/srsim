"use client";
import React from "react";
import { ViewerContext } from "./provider";
import { Card, InProgress, Tabs, TabsContent, TabsList, TabsTrigger } from "ui";
import { ExecutorContext } from "../exec/provider";
import { Summary } from "./summary";
import { Sample } from "./sample";

export default function Viewer() {
  const { state } = React.useContext(ViewerContext);
  const { supplier } = React.useContext(ExecutorContext);
  const exec = supplier();
  if (state.error !== null && state.error !== "") {
    return <pre>{state.error}</pre>;
  }
  if (state.data === null) {
    return <div>No results yet...</div>;
  }
  if (state.data.statistics === undefined) {
    return <div>No stats available yet...</div>;
  }
  return (
    <div className="p-3 flex flex-col">
      {exec.running() ? (
        <InProgress
          val={state.progress ?? 0}
          className="mb-2"
          cancel={() => {
            exec.cancel();
          }}
        />
      ) : null}
      <Tabs defaultValue="summary" className="w-full">
        <TabsList>
          <TabsTrigger value="summary">Result</TabsTrigger>
          <TabsTrigger value="sample">Sample</TabsTrigger>
        </TabsList>
        <TabsContent value="summary">
          <Card className="p-4">
            <Summary data={state.data} />
          </Card>
        </TabsContent>
        <TabsContent value="sample">
          <Card className="p-4">
            <Sample data={state.data} />
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
