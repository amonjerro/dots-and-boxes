<!-- Adapted from https://next.melt-ui.com/components/select/ -->
<script lang="ts">
    import { Select } from "melt/builders";

    let {label, options, name, disabled = false}: {label: string, options: string[], name: string, disabled?: boolean} = $props();
    type Option = (typeof options)[number];
        
    const select = new Select<Option>();
    select.value = options[0];
</script>

<!-- Select Label -->
<label {...select.label} class:pointer-events-none={disabled} class={disabled ? "filter brightness-75" : ""}>{label}</label>
<!-- Select Trigger -->
<button {disabled} type="button" {...select.trigger} class="btn btn-lg preset-filled min-w-64 capitalize" class:pointer-events-none={disabled}>
    {select.value}
</button>

<!-- Select Content (Options) -->
<div {...select.content} class="p-4 bg-surface-50-950 outline-1 outline-surface-950-50 rounded-base text-xl min-w-64 capitalize">
    {#each options as option}
        <div {...select.getOption(option)}>
            <span>{option}</span>
        </div>
    {/each}
</div>

<!-- Hidden input to track selected option -->
<input type="hidden" {disabled} {name} value={select.value} />

<!-- Select styling using Melt UI data attributes -->
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