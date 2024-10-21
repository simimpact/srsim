'use client';
import React from 'react';
import Simulator from './simulator';
import Navigation from './navigation';
import {ViewerProvider} from './viewer/provider';

export default function Home() {
  return (
    <>
      <Simulator />
    </>
  );
}
