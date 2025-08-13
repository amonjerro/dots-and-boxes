<!-- A utility component for displaying FontAwesome icons -->
<script lang="ts">
    import type { Icon, IconSize } from "$lib/types/uiTypes";

    let {name, size}: {name: Icon, size?: IconSize} = $props();
</script>

<i class={`fa-solid fa-${name} fa-${size}`}></i>