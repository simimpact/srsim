"use client";

import * as DialogPrimitive from "@radix-ui/react-dialog";
import { VariantProps, cva } from "class-variance-authority";
import { X } from "lucide-react";
import * as React from "react";
import { cn } from "@/utils/classname";

const overlayVariants = cva(
  "fixed inset-0 z-50 bg-background/80 backdrop-blur-sm data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0"
);

const contentVariants = cva(
  "fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg md:w-full"
);

const headerVariants = cva("flex flex-col space-y-1.5 text-center sm:text-left");

const footerVariants = cva("flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2");

const titleVariants = cva("text-lg font-semibold leading-none tracking-tight");

const descriptionVariants = cva("text-sm text-muted-foreground");

const Dialog = DialogPrimitive.Root;

const DialogTrigger = DialogPrimitive.Trigger;

const portalVariants = cva("fixed inset-0 z-50 flex items-start justify-center sm:items-center");

const closeVariants = cva(
  "absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
);

const DialogPortal = ({ className, ...props }: DialogPrimitive.DialogPortalProps) => (
  <DialogPrimitive.Portal className={cn(portalVariants({ className }))} {...props} />
);
DialogPortal.displayName = DialogPrimitive.Portal.displayName;

const DialogOverlay = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Overlay>,
  OverlayProps
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Overlay ref={ref} className={cn(overlayVariants({ className }))} {...props} />
));
DialogOverlay.displayName = DialogPrimitive.Overlay.displayName;

const DialogContent = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <DialogPortal>
    <DialogOverlay />
    <DialogPrimitive.Content ref={ref} className={cn(contentVariants({ className }))} {...props}>
      {children}
      <DialogPrimitive.Close className={closeVariants()}>
        <X className="h-4 w-4" />
        <span className="sr-only">Close</span>
      </DialogPrimitive.Close>
    </DialogPrimitive.Content>
  </DialogPortal>
));
DialogContent.displayName = DialogPrimitive.Content.displayName;

const DialogHeader = ({ className, ...props }: HeaderFooterProps) => (
  <div className={cn(headerVariants({ className }))} {...props} />
);
DialogHeader.displayName = "DialogHeader";

const DialogFooter = ({ className, ...props }: HeaderFooterProps) => (
  <div className={cn(footerVariants({ className }))} {...props} />
);
DialogFooter.displayName = "DialogFooter";

const DialogTitle = React.forwardRef<React.ElementRef<typeof DialogPrimitive.Title>, TitleProps>(
  ({ className, ...props }, ref) => (
    <DialogPrimitive.Title ref={ref} className={cn(titleVariants({ className }))} {...props} />
  )
);
DialogTitle.displayName = DialogPrimitive.Title.displayName;

const DialogDescription = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Description>,
  DescriptionProps
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Description
    ref={ref}
    className={cn(descriptionVariants({ className }))}
    {...props}
  />
));
DialogDescription.displayName = DialogPrimitive.Description.displayName;

interface OverlayProps
  extends React.ComponentPropsWithoutRef<typeof DialogPrimitive.Overlay>,
    VariantProps<typeof overlayVariants> {}
interface HeaderFooterProps extends React.HTMLAttributes<HTMLDivElement> {}
interface TitleProps extends React.ComponentPropsWithoutRef<typeof DialogPrimitive.Title> {}
interface DescriptionProps
  extends React.ComponentPropsWithoutRef<typeof DialogPrimitive.Description> {}

export {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
};
