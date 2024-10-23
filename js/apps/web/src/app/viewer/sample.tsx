"use client";
import React from "react";
import { ExecutorContext } from "../exec/provider";

type PropType = {};

export const Sample = ({}: PropType) => {
  const { supplier } = React.useContext(ExecutorContext);
  const exec = supplier();
  return <div>sample</div>;
};
