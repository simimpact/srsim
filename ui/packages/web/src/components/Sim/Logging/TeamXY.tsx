import { curveCardinal } from "@visx/curve";
import { ParentSize } from "@visx/responsive";
import { XYChart, Axis, Grid, Tooltip, AreaStack, AreaSeries, lightTheme } from "@visx/xychart";
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
    <ParentSize>
      {parent => (
        <XYChart
          theme={lightTheme}
          xScale={{ type: "band" }}
          yScale={{ type: "linear" }}
          width={parent.width}
          height={400}
          resizeObserverPolyfill={ResizeObserver}
          margin={{ left: 0, right: 0, top: 16, bottom: 16 }}
        >
          <Axis orientation={"bottom"} />
          <Grid columns={false} numTicks={4} />
          <AreaStack offset="expand" curve={curveCardinal}>
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
            showVerticalCrosshair
            showSeriesGlyphs
            renderTooltip={({ tooltipData, colorScale }) => (
              <>
                <p>Turn {tooltipData?.nearestDatum?.index}</p>
                {tooltipData?.nearestDatum?.datum && (
                  <p>
                    {Object.values(tooltipData.datumByKey)
                      .reverse()
                      .map((dist, index) => (
                        <p
                          key={index}
                          style={{
                            color: colorScale?.(dist.key),
                            textDecoration:
                              tooltipData.nearestDatum?.key === dist.key ? "underline" : undefined,
                          }}
                        >
                          {dist.datum.value.name}: {(dist.datum.value.rate * 100).toFixed(2)} %
                        </p>
                      ))}
                  </p>
                )}
              </>
            )}
          />
        </XYChart>
      )}
    </ParentSize>
  );
};
export { TeamXY };
