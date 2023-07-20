import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

/**
 * helper function to combine classes
 * @param inputs
 * @returns
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
