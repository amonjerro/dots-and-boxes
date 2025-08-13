<script lang="ts">
    import Icon from "$lib/components/Icon.svelte";
    import Select from "$lib/components/Select.svelte";
    import { difficulties, gridSizes, type PlayerType } from "$lib/types/gameTypes";
    import { Avatar } from "@skeletonlabs/skeleton-svelte"
    import { flip } from "svelte/animate";
    import { scale } from "svelte/transition";
    
    const players: PlayerType[] = $state([ "user", "robot" ]);
    const playersContainRobot = $derived(players.includes("robot"));
    const playersContainUser = $derived(players.includes("user"));
    const moreThanOneUser = $derived(players.filter(p => p === "user").length > 1);

    /**
     * If there are fewer than four players, add a robot.
     * Otherwise, overflow to first 2 players.
     */
    const addPlayer = () => {
        if (players.length < 4) {
            players.push("robot");
        }
        else {
            players.splice(2, 2);
        }
    };

    /**
     * If there are more than two players, remove the last one.
     * Otherwise, underflow to 4 players.
     */
    const removePlayer = () => {
        if (players.length > 2) {
            players.pop();
        }
        else{
            players.push("robot", "robot");
        }
    };

    /**
     * Toggle the player type between "user" and "robot".
     * @param index The index of the player to toggle.
     */
    const togglePlayer = (index: number) => {
        players[index] = players[index] === "user" ? "robot" : "user";
    }

    $inspect(players, playersContainRobot, playersContainUser, moreThanOneUser);
</script>

<!-- Game Parameters Form -->
<form class="flex flex-col gap-8 items-center [&_label]:h3" action="">
    <h2 class="h2 mt-8">Start a Game</h2>

    <!-- Player Settings -->
    <div class="flex flex-col gap-4 items-center">
        <label for="playerCount">Player Count</label>
        <!-- Player Incrementer -->
        <div class="flex gap-4 items-center">
            <button
                type="button"
                class="btn-icon btn-icon-lg preset-filled"
                onclick={removePlayer}
                aria-label="remove player"
            >
                <Icon name="chevron-left" />
            </button>
            
            <p class="text-2xl">{players.length} Players</p>
    
            <button
                type="button"
                class="btn-icon btn-icon-lg preset-filled"
                onclick={addPlayer}
                aria-label="add player"
            >
                <Icon name="chevron-right" />
            </button>
        </div>
    
        <!-- Player Editor -->
        <div class="flex gap-4">
            {#each players as player, index (index)}
                <button
                    type="button"
                    transition:scale={{ duration: 100 }} 
                    animate:flip={{ duration: 100 }} 
                    onclick={() => togglePlayer(index)}
                >
                    <Avatar name="icon" background="preset-filled">
                        <Icon name={player} />
                    </Avatar>
                    <!-- Hidden input to track player type -->
                    <input type="hidden" name="player" value={player} />
                </button>
            {/each}
        </div>
    
        <!-- Player Editor Instructions -->
        <div class="text-center">
            {#if !playersContainUser}
                <p class="text-error-500">At least one human is required.</p>
            {:else}
                <p class="text-surface-200">Click a player to change their type.</p>
            {/if}
        </div>
    </div>

    <!-- Computer Difficulty -->
    <div class="flex flex-col items-center">
        <Select label="Difficulty" options={[...difficulties]} disabled={!playersContainRobot}/>
    </div>

    <!-- Grid Size -->
    <div class="flex flex-col items-center">
        <Select label="Grid Size" options={[...gridSizes]} />
    </div>

    <!-- Submit Button -->
    <button type="submit" disabled={!playersContainUser} class:pointer-events-none={!playersContainUser} class="btn preset-filled-success-500 btn-lg mt-2 min-w-64 uppercase">
        {moreThanOneUser ? "Create Lobby" : "Start Game"}
    </button>
</form>

<!-- Game Instructions -->
<div class="bg-[url('https://placehold.co/450x300?text=Dots+And+Boxes')] bg-contain flex justify-center mt-10">
    <article class="max-w-[750px] space-y-4 bg-surface-500 p-8">
        <h2 class="h3">How to Play "Dots and Boxes"</h2>
        <p>In Dots and Boxes, players take turns drawing lines between dots on a grid. The objective is to complete more boxes than your opponent.</p>
        <p>On your turn, click between two dots to create a line between them. If you complete a box, you earn a point and get to take another turn.</p>
        <p>The game ends when all boxes are completed. The player with the most points (boxes) wins!</p>
    </article>
</div>