'use client';
import React, {ReactNode} from 'react';
import {cn} from 'ui/src/lib/utils';

type GroupProps = {
  children: ReactNode;
  className?: string;
};

const Group = ({children, className}: GroupProps) => {
  const cls = cn(
    className,
    'grid overflow-hidden',
    'grid-cols-2 sm:grid-cols-6',
    'gap-y-2',
    'sm:gap-2',
  );

  return <div className={cls}>{children}</div>;
};
