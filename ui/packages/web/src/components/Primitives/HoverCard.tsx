"use client";

import * as HoverCardPrimitive from "@radix-ui/react-hover-card";
import { cva } from "class-variance-authority";
import * as React from "react";
import { cn } from "@/utils/classname";

const hoverCardVariants = cva(
  "z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none animate-in zoom-in-90"
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
