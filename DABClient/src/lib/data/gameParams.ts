// Defined here for easy use in game & backend api logic
export const gameDefaults = {
    players: ["user", "robot"],
    difficulty: "easy",
    gridSize: "small",
} as const;

export const gridSizes = ["small", "medium", "large"] as const;

export const difficulties = ["easy", "medium", "hard"] as const;
