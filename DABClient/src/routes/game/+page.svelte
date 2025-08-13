<script lang="ts">
    import { page } from '$app/state';
    import { gameDefaults } from '$lib/data/gameParams';

    // Extracting players, difficulty, and gridSize from the URL search parameters
    let players = page.url.searchParams.getAll('player');
    if(!players.length) players = [...gameDefaults.players];

    const difficulty = page.url.searchParams.get('difficulty') || gameDefaults.difficulty;
    const gridSize = page.url.searchParams.get('gridSize') || gameDefaults.gridSize;
</script>

{#each players as player}
    <p>{player}</p>
{/each}
<br>
<p>{difficulty}</p>
<br>
<p>{gridSize}</p>