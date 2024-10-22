'use client';
import React from 'react';
import {model} from '@srsim/ts-types';
import {RollupCard} from 'ui';

type RollupCardProps = {
  data: model.Statistics;
};

const fmt = (val?: number) => {
  return val?.toLocaleString(navigator.language, {maximumFractionDigits: 0});
};

const DamagePerCycleRollup = (props: RollupCardProps) => {
  const d = props.data.total_damage_dealt_per_cycle;
  return (
    <RollupCard
      color="bg-red"
      title="Damage Per Cycle (DPC)"
      value={fmt(d?.mean)}
      auxStats={[
        {title: 'min', value: fmt(d?.min)},
        {title: 'max', value: fmt(d?.max)},
        {title: 'std', value: fmt(d?.sd)},
        {title: 'p25', value: fmt(d?.q1)},
        {title: 'p50', value: fmt(d?.q2)},
        {title: 'p75', value: fmt(d?.q3)},
      ]}
    />
  );
};

const TotalDamageDealt = (props: RollupCardProps) => {
  const d = props.data.total_damage_dealt;
  return (
    <RollupCard
      color="bg-blue"
      title="Total Damage Dealt (TDD)"
      value={fmt(d?.mean)}
      auxStats={[
        {title: 'min', value: fmt(d?.min)},
        {title: 'max', value: fmt(d?.max)},
        {title: 'std', value: fmt(d?.sd)},
        // {title: 'p25', value: fmt(d?.q1)},
        // {title: 'p50', value: fmt(d?.q2)},
        // {title: 'p75', value: fmt(d?.q3)},
      ]}
    />
  );
};

const TotalDamageTaken = (props: RollupCardProps) => {
  const d = props.data.total_damage_taken;
  return (
    <RollupCard
      color="bg-blue"
      title="Total Damage Taken (TDT)"
      value={fmt(d?.mean)}
      auxStats={[
        {title: 'min', value: fmt(d?.min)},
        {title: 'max', value: fmt(d?.max)},
        {title: 'std', value: fmt(d?.sd)},
        // {title: 'p25', value: fmt(d?.q1)},
        // {title: 'p50', value: fmt(d?.q2)},
        // {title: 'p75', value: fmt(d?.q3)},
      ]}
    />
  );
};

export const RollupGrid = (props: RollupCardProps) => {
  return (
    <div className="grid grid-cols-3">
      <DamagePerCycleRollup {...props} />
      <TotalDamageDealt {...props} />
      <TotalDamageTaken {...props} />
    </div>
  );
};
