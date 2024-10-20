'use client';
import React from 'react';
import {ViewerContext} from './provider';
import {OverviewStatsBarGraph} from 'ui';

export default function Viewer() {
  const {state} = React.useContext(ViewerContext);
  if (state.error !== null && state.error !== '') {
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
      <div>
        <p>Total damage</p>
        {JSON.stringify(state.data.statistics?.total_damage_dealt, null, 2)}
      </div>
      {state.data.statistics.damage_dealt_by_cycle === undefined ? null : (
        <div>
          <p>Avg Damage By Cycle</p>
          <OverviewStatsBarGraph
            dataKey="avg_dmg_by_cycle"
            data={state.data.statistics.damage_dealt_by_cycle}
          />
        </div>
      )}
    </div>
  );
}
