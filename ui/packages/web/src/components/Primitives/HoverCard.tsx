"use client";

import * as HoverCardPrimitive from "@radix-ui/react-hover-card";
import { cva } from "class-variance-authority";
import * as React from "react";
import { cn } from "@/utils/classname";

const hoverCardVariants = cva(
  "z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2"
);

const HoverCard = HoverCardPrimitive.Root;

const HoverCardTrigger = HoverCardPrimitive.Trigger;

const HoverCardContent = React.forwardRef<
  React.ElementRef<typeof HoverCardPrimitive.Content>,
  ContentProps
>(({ className, align = "center", sideOffset = 4, ...props }, ref) => (
  <HoverCardPrimitive.Content
    ref={ref}
    align={align}
    sideOffset={sideOffset}
    className={cn(hoverCardVariants({ className }))}
    {...props}
  />
));
HoverCardContent.displayName = HoverCardPrimitive.Content.displayName;

interface ContentProps extends React.ComponentPropsWithoutRef<typeof HoverCardPrimitive.Content> {}

export { HoverCard, HoverCardTrigger, HoverCardContent };
