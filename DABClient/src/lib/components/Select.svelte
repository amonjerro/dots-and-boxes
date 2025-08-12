<script lang="ts">
    import { Select } from "melt/builders";

    // declare options prop
    let {label, options}: {label: string, options: string[]} = $props();
    type Option = (typeof options)[number];
        
    const select = new Select<Option>();
    select.value = options[0];
</script>

<label {...select.label}>{label}</label>
<button type="button"  {...select.trigger} class="btn btn-lg preset-filled min-w-64">
    {select.value}
</button>

<div {...select.content} class="p-4 bg-surface-50-950 outline-1 outline-surface-950-50 rounded-base text-xl min-w-64">
    {#each options as option}
        <div {...select.getOption(option)}>
            <span>{option}</span>
        </div>
    {/each}
</div>

<!-- Hidden input to track selected option -->
<input type="hidden" name={label} value={select.value} />

<style>
    [data-melt-select-option] {
        padding: 8px 16px;
        border-radius: var(--radius-base);
    }

    [data-highlighted] {
        background-color: var(--color-surface-800);
        color: white;
    }

    [data-melt-select-option][aria-selected="true"] {
        background-color: var(--color-surface-950-50);
        color: var(--color-surface-50-950);
    }

    [data-melt-select-option][aria-selected="true"][data-highlighted] {
        filter: brightness(75%);
    }
</style>