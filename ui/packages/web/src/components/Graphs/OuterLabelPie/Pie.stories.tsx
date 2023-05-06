import { Meta, StoryObj } from "@storybook/react";
import { OuterLabelPie } from "./Pie";

interface TestData {
  index: number;
  value: number;
  label: string;
}

const TestablePie = OuterLabelPie<TestData>;

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

const meta = {
  title: "Components/Graphs/OuterLabelPie",
  component: TestablePie,
} satisfies Meta<typeof TestablePie>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Basic: Story = {
  args: {
    data: [
      { label: "A", index: 0, value: 0.25 },
      { label: "B", index: 1, value: 0.25 },
      { label: "C", index: 2, value: 0.25 },
      { label: "D", index: 3, value: 0.25 },
    ],
    pieValue: d => d.value,
    width: 500,
    height: 500,
    color: d => colors[d.index],
    labelText: d => d.label,
    labelValue: d => d.value.toLocaleString("en", { style: "percent" }),
  },
};

export const ThinSlices: Story = {
  args: {
    data: [
      { label: "A", index: 0, value: 0.01 },
      { label: "B", index: 1, value: 0.01 },
      { label: "C", index: 2, value: 0.08 },
      { label: "D", index: 3, value: 0.05 },
      { label: "E", index: 4, value: 0.05 },
      { label: "F", index: 5, value: 0.8 },
    ],
    pieValue: d => d.value,
    width: 500,
    height: 500,
    color: d => colors[d.index],
    labelText: d => d.label,
    labelValue: d => d.value.toLocaleString("en", { style: "percent" }),
  },
};
