import { XYChart, Axis, Grid, Tooltip, darkTheme, AreaStack, AreaSeries } from "@visx/xychart";
import { InTeamDistribution } from "@/bindings/MvpWrapper";

interface Props {
  data: [InTeamDistribution, InTeamDistribution, InTeamDistribution, InTeamDistribution][];
}
const TeamXY = ({ data }: Props) => {
  // [{name: 'a', rate: 0.0} ... ]
  // [{name: 'a', rate: 0.0} ... ]
  // [{name: 'a', rate: 0.0} ... ]
  // [{name: 'a', rate: 0.0} ... ]
  interface ProcessedDataByChar {
    turn: number;
    value: InTeamDistribution;
  }
  const processedData = (charIndex: number): ProcessedDataByChar[] =>
    data.map((e, index) => ({ turn: index, value: e[charIndex] }));

  return (
    <XYChart
      theme={darkTheme}
      xScale={{ type: "band" }}
      yScale={{ type: "linear" }}
      height={600}
      width={600}
    >
      <Axis orientation={"bottom"} />
      <Grid columns={false} numTicks={4} />
      <AreaStack>
        {[0, 1, 2, 3].map(charIndex => (
          <AreaSeries
            key={charIndex}
            dataKey={`Line-${charIndex}`}
            data={processedData(charIndex)}
            xAccessor={e => e.turn}
            yAccessor={e => e.value.rate}
            fillOpacity={0.4}
          />
        ))}
      </AreaStack>
      <Tooltip<ProcessedDataByChar>
        renderTooltip={({ tooltipData }) => (
          <>
            <span>{tooltipData?.nearestDatum?.index}</span>
            <br />
            {tooltipData?.nearestDatum?.datum && (
              <p>
                {tooltipData.nearestDatum.datum.value.name}:{" "}
                {(tooltipData.nearestDatum.datum.value.rate * 100).toFixed(2)} %
              </p>
            )}
          </>
        )}
      />
    </XYChart>
  );
};
export { TeamXY };
