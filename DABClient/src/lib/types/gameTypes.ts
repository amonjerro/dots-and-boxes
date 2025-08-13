export type PlayerType = "user" | "robot";

// Defined here for easy use in game & backend api logic
export const gridSizes = ["small", "medium", "large"] as const;
export type GridSize = (typeof gridSizes)[number];

export const difficulties = ["easy", "medium", "hard"] as const;
export type Difficulty = (typeof difficulties)[number];

/**
 * Icon Types - Used to help keep track of icon usage from FontAwesome
 * It's very easy to overdo it.
 */
export type Icon = "user" | "robot" | "chevron-left" | "chevron-right";
export type IconSize = "xs" | "sm" | "lg" | "xl";
