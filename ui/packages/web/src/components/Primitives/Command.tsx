"use client";

import { DialogProps } from "@radix-ui/react-dialog";
import { VariantProps, cva } from "class-variance-authority";
import { Command as CommandPrimitive } from "cmdk";
import { Search } from "lucide-react";
import * as React from "react";
import { cn } from "@/utils/classname";
import { Dialog, DialogContent } from "./Dialog";

const commandVariants = cva(
  "flex h-full w-full flex-col overflow-hidden rounded-md bg-popover text-popover-foreground"
);

const inputVariants = cva("flex items-center border-b px-3", {
  variants: {
    variant: {
      default:
        "flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});

const itemVariants = cva(
  "relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none aria-selected:bg-accent aria-selected:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50"
);

const listVariants = cva("max-h-[300px] overflow-y-auto overflow-x-hidden");

const contentVariants = cva("overflow-hidden p-0 shadow-2xl");

const groupVariants = cva(
  "[&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:py-1.5 [&_[cmdk-group-heading]]:text-xs [&_[cmdk-group-heading]]:font-medium [&_[cmdk-group-heading]]:text-muted-foreground",
  {
    variants: {
      variant: { default: "overflow-hidden p-1 text-foreground" },
    },
    defaultVariants: { variant: "default" },
  }
);

const separatorVariants = cva("-mx-1 h-px bg-border");

const shortcutVariants = cva("ml-auto text-xs tracking-widest text-muted-foreground");

const Command = React.forwardRef<React.ElementRef<typeof CommandPrimitive>, CommandProps>(
  ({ className, ...props }, ref) => (
    <CommandPrimitive ref={ref} className={cn(commandVariants({ className }))} {...props} />
  )
);
Command.displayName = CommandPrimitive.displayName;

interface CommandDialogProps extends DialogProps {}

const CommandDialog = ({ children, ...props }: CommandDialogProps) => {
  return (
    <Dialog {...props}>
      <DialogContent className={cn(contentVariants())}>
        <Command className="[&_[cmdk-group-heading]]:text-muted-foreground [&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:font-medium [&_[cmdk-group]:not([hidden])_~[cmdk-group]]:pt-0 [&_[cmdk-group]]:px-2 [&_[cmdk-input-wrapper]_svg]:h-5 [&_[cmdk-input-wrapper]_svg]:w-5 [&_[cmdk-input]]:h-12 [&_[cmdk-item]]:px-2 [&_[cmdk-item]]:py-3 [&_[cmdk-item]_svg]:h-5 [&_[cmdk-item]_svg]:w-5">
          {children}
        </Command>
      </DialogContent>
    </Dialog>
  );
};

const CommandInput = React.forwardRef<React.ElementRef<typeof CommandPrimitive.Input>, InputProps>(
  ({ className, variant, ...props }, ref) => (
    // eslint-disable-next-line react/no-unknown-property
    <div className="flex items-center border-b px-3" cmdk-input-wrapper="">
      <Search className="mr-2 h-4 w-4 shrink-0 opacity-50" />
      <CommandPrimitive.Input
        ref={ref}
        className={cn(inputVariants({ variant, className }))}
        {...props}
      />
    </div>
  )
);

CommandInput.displayName = CommandPrimitive.Input.displayName;

const CommandList = React.forwardRef<React.ElementRef<typeof CommandPrimitive.List>, ListProps>(
  ({ className, ...props }, ref) => (
    <CommandPrimitive.List ref={ref} className={cn(listVariants({ className }))} {...props} />
  )
);

CommandList.displayName = CommandPrimitive.List.displayName;

const CommandEmpty = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Empty>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Empty>
>((props, ref) => (
  <CommandPrimitive.Empty ref={ref} className="py-6 text-center text-sm" {...props} />
));

CommandEmpty.displayName = CommandPrimitive.Empty.displayName;

const CommandGroup = React.forwardRef<React.ElementRef<typeof CommandPrimitive.Group>, GroupProps>(
  ({ className, variant, ...props }, ref) => (
    <CommandPrimitive.Group
      ref={ref}
      className={cn(groupVariants({ variant, className }))}
      {...props}
    />
  )
);

CommandGroup.displayName = CommandPrimitive.Group.displayName;

const CommandSeparator = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Separator>,
  SeparatorProps
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Separator
    ref={ref}
    className={cn(separatorVariants({ className }))}
    {...props}
  />
));
CommandSeparator.displayName = CommandPrimitive.Separator.displayName;

const CommandItem = React.forwardRef<React.ElementRef<typeof CommandPrimitive.Item>, ItemProps>(
  ({ className, ...props }, ref) => (
    <CommandPrimitive.Item ref={ref} className={cn(itemVariants({ className }))} {...props} />
  )
);

CommandItem.displayName = CommandPrimitive.Item.displayName;

const CommandShortcut = ({ className, ...props }: ShortcutProps) => {
  return <span className={cn(shortcutVariants({ className }))} {...props} />;
};
CommandShortcut.displayName = "CommandShortcut";

interface CommandProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive>,
    VariantProps<typeof commandVariants> {}

interface InputProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive.Input>,
    VariantProps<typeof inputVariants> {}

interface ListProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive.List>,
    VariantProps<typeof listVariants> {}

interface GroupProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive.Group>,
    VariantProps<typeof groupVariants> {}

interface SeparatorProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive.Separator>,
    VariantProps<typeof separatorVariants> {}

interface ItemProps
  extends React.ComponentPropsWithoutRef<typeof CommandPrimitive.Item>,
    VariantProps<typeof itemVariants> {}

interface ShortcutProps
  extends React.HTMLAttributes<HTMLSpanElement>,
    VariantProps<typeof shortcutVariants> {}

export {
  Command,
  CommandDialog,
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandShortcut,
  CommandSeparator,
};
