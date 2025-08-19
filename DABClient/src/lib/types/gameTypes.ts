import type { difficulties, gridSizes } from "$lib/data/gameParams";

export type PlayerType = "user" | "robot";

// Defined here for easy use in game & backend api logic
export type GridSize = (typeof gridSizes)[number];
export type Difficulty = (typeof difficulties)[number];
