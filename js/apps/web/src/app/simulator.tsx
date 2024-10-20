'use client';
import React from 'react';
import {Button, Editor} from '@ui/components';

export default function Simulator() {
  const [cfg, setCfg] = React.useState<string>('');
  return (
    <div className="m-3">
      <Editor cfg={cfg} onChange={(v) => setCfg(v)} className="mb-2"></Editor>
      <div className="sticky bottom-0 flex flex-col gap-y-1 z-10">
        <Button variant="secondary">Run</Button>
      </div>
    </div>
  );
}
