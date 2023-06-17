import { cva } from "class-variance-authority";

export const elementVariants = cva("", {
  variants: {
    element: {
      none: "",
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
    element: "none",
  },
});
