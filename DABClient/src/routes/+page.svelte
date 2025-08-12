<script lang="ts">
    import Select from "$lib/components/Select.svelte";
    import type { PlayerType } from "$lib/types/gameTypes";
    import { Avatar } from "@skeletonlabs/skeleton-svelte"
    import { flip } from "svelte/animate";
    import { scale } from "svelte/transition";

    const players: PlayerType[] = $state([ "human", "computer" ]);
    const playersContainComputer = $derived(players.includes("computer"));
    const playersContainHuman = $derived(players.includes("human"));

    const addPlayer = () => {
        if (players.length < 4) {
            players.push("computer");
        }
        else {
            players.splice(2, 2);
        }
    };

    const removePlayer = () => {
        if (players.length > 2) {
            players.pop();
        }
        else{
            players.push("computer", "computer");
        }
    };

    const togglePlayer = (index: number) => {
        players[index] = players[index] === "human" ? "computer" : "human";
    }

    const computerDifficultyOptions = ["Easy", "Medium", "Hard"];
    const gridSizeOptions = ["Small (4x3)", "Medium (7x6)", "Large (10x9)"];

    const iconMap = {
        human: "fa-solid fa-user fa-xl",
        computer: "fa-solid fa-robot fa-xl"
    };

    $inspect(players, playersContainComputer, playersContainHuman);
</script>

<form class="flex flex-col gap-8 items-center [&_label]:h4" action="">
    <h2 class="h2 mt-8">Start a Game</h2>

    <div class="flex flex-col gap-4 items-center">
        <label for="playerCount">Player Count</label>
        <div class="flex gap-4 items-center">
            <button
                type="button"
                class="btn-icon preset-filled"
                onclick={removePlayer}
                aria-label="remove player"
            >
                <i class="fa-solid fa-chevron-left"></i>
            </button>
            
            <p>{players.length} Players</p>
    
            <button
                type="button"
                class="btn-icon preset-filled"
                onclick={addPlayer}
                aria-label="add player"
            >
                <i class="fa-solid fa-chevron-right"></i>
            </button>
        </div>
    
        <div class="flex gap-4">
            {#each players as player, index (index)}
                <button
                    type="button"
                    transition:scale={{ duration: 100 }} 
                    animate:flip={{ duration: 100 }} 
                    onclick={() => togglePlayer(index)}
                >
                    <Avatar name="icon" background="preset-filled">
                        <i class={iconMap[player]}></i>
                    </Avatar>
                    <!-- Hidden input to track player type -->
                    <input type="hidden" name="player" value={player} />
                </button>
            {/each}
        </div>
    
        <div>
            {#if !playersContainHuman}
                <p class="text-error-500">At least one human player is required.</p>
            {:else}
                <p class="text-surface-200">Click a player to change their type.</p>
            {/if}
        </div>
    </div>

    <!-- Computer Difficulty -->
    {#if playersContainComputer}
        <div class="flex flex-col items-center">
            <Select label="Computer Difficulty" options={computerDifficultyOptions} />
        </div>
    {/if}

    <!-- Grid Size -->
    <div class="flex flex-col items-center">
        <Select label="Grid Size" options={gridSizeOptions} />
    </div>

    <button type="submit" disabled={!playersContainHuman} class="btn preset-filled">START</button>
</form>

<div class="bg-[url('https://placehold.co/450x300?text=Dots+And+Boxes')] bg-contain flex justify-center mt-10">
    <article class="max-w-[750px] space-y-4 bg-surface-500 p-8">
        <h2 class="h3">How to Play "Dots and Boxes"</h2>
        <p>In Dots and Boxes, players take turns drawing lines between dots on a grid. The objective is to complete more boxes than your opponent.</p>
        <p>On your turn, click between two dots to create a line between them. If you complete a box, you earn a point and get to take another turn.</p>
        <p>The game ends when all boxes are completed. The player with the most points (boxes) wins!</p>
    </article>
</div>