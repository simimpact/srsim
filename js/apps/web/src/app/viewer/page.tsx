'use client';
import React from 'react';
import {ViewerContext} from './provider';

export default function Viewer() {
  const {state} = React.useContext(ViewerContext);
  if (state.error !== null && state.error !== '') {
    return <pre>{state.error}</pre>;
  }
  if (state.data === null) {
    return <div>No results yet...</div>;
  }
  return <pre>{JSON.stringify(state.data.statistics, null, 2)}</pre>;
}
