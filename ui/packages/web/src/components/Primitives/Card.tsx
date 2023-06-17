import { type VariantProps, cva } from "class-variance-authority";
import * as React from "react";
import { cn } from "@/utils/classname";

const cardVariants = cva("rounded-lg border shadow-sm", {
  variants: {
    variant: {
      default: "bg-card text-card-foreground",
      fire: "bg-fire text-fire-foreground hover:bg-fire/90",
      ice: "bg-ice text-ice-foreground hover:bg-ice/90",
      wind: "bg-wind text-wind-foreground hover:bg-wind/90",
      physical: "bg-physical text-physical-foreground hover:bg-physical/90",
      lightning: "bg-lightning text-lightning-foreground hover:bg-lightning/90",
      quantum: "bg-quantum text-quantum-foreground hover:bg-quantum/90",
      imaginary: "bg-imaginary text-imaginary-foreground hover:bg-imaginary/90",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});
const cardHeaderVariant = cva("flex flex-col space-y-1.5 p-6");
const cardTitleVariant = cva("text-lg font-semibold leading-none tracking-tight");
const cardDescriptionVariant = cva("text-sm text-muted-foreground");
const cardContentVariant = cva("p-6 pt-0");
const cardFooterVariant = cva("flex items-center p-6 pt-0");

/** Anatomy: https://ui.shadcn.com/docs/components/card#usage */
const Card = React.forwardRef<HTMLDivElement, CardProps>(
  ({ className, variant, ...props }, ref) => (
    <div ref={ref} className={cn(cardVariants({ variant, className }))} {...props} />
  )
);
Card.displayName = "Card";

const CardHeader = React.forwardRef<HTMLDivElement, CardProps>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn(cardHeaderVariant({ className }))} {...props} />
));
CardHeader.displayName = "CardHeader";

const CardTitle = React.forwardRef<HTMLParagraphElement, CardHeadingProps>(
  ({ className, ...props }, ref) => (
    <h3 ref={ref} className={cn(cardTitleVariant({ className }))} {...props} />
  )
);
CardTitle.displayName = "CardTitle";

const CardDescription = React.forwardRef<HTMLParagraphElement, CardParagraphProps>(
  ({ className, ...props }, ref) => (
    <p ref={ref} className={cn(cardDescriptionVariant({ className }))} {...props} />
  )
);
CardDescription.displayName = "CardDescription";

const CardContent = React.forwardRef<HTMLDivElement, CardProps>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn(cardContentVariant({ className }))} {...props} />
));
CardContent.displayName = "CardContent";

const CardFooter = React.forwardRef<HTMLDivElement, CardProps>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn(cardFooterVariant({ className }))} {...props} />
));
CardFooter.displayName = "CardFooter";

export interface CardProps
  extends React.HTMLAttributes<HTMLDivElement>,
    VariantProps<typeof cardVariants> {}
interface CardHeadingProps
  extends React.HTMLAttributes<HTMLHeadingElement>,
    VariantProps<typeof cardVariants> {}
interface CardParagraphProps
  extends React.HTMLAttributes<HTMLParagraphElement>,
    VariantProps<typeof cardVariants> {}

export { Card, CardHeader, CardFooter, CardTitle, CardDescription, CardContent, cardVariants };
