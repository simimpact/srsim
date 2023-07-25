import { cva } from "class-variance-authority";

export const elementVariants = cva("", {
  variants: {
    element: {
      none: "",
      Fire: "bg-fire text-fire-foreground hover:bg-fire/90",
      Ice: "bg-ice text-ice-foreground hover:bg-ice/90",
      Wind: "bg-wind text-wind-foreground hover:bg-wind/90",
      Physical: "bg-physical text-physical-foreground hover:bg-physical/90",
      Lightning: "bg-lightning text-lightning-foreground hover:bg-lightning/90",
      Quantum: "bg-quantum text-quantum-foreground hover:bg-quantum/90",
      Imaginary: "bg-imaginary text-imaginary-foreground hover:bg-imaginary/90",
    },
    border: {
      none: "",
      Fire: "border-fire",
      Ice: "border-ice",
      Wind: "border-wind",
      Physical: "border-physical",
      Lightning: "border-lightning",
      Quantum: "border-quantum",
      Imaginary: "border-imaginary",
    },
  },
  defaultVariants: {
    element: "none",
    border: "none",
  },
});

// TODO: update colors
export const rarityVariants = cva("", {
  variants: {
    rarity: {
      1: "bg-physical text-physical-foreground hover:bg-physical/90",
      2: "bg-wind text-wind-foreground hover:bg-wind/90",
      3: "bg-ice text-ice-foreground hover:bg-ice/90",
      4: "bg-quantum text-quantum-foreground hover:bg-quantum/90",
      5: "bg-imaginary text-imaginary-foreground hover:bg-imaginary/90",
    },
  },
  defaultVariants: { rarity: 1 },
});
