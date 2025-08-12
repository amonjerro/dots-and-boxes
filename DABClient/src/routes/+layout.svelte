<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
    const title = "Dots and Boxes";
	
	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/7.0.0/css/all.min.css" integrity="sha512-DxV+EoADOkOygM4IR9yXP8Sb2qwgidEmeqAEmDKIOfPRQZOWbXCzLC6vjbZyy0vPisbH2SyW27+ddLVCN+OMzQ==" crossorigin="anonymous" referrerpolicy="no-referrer" />
</svelte:head>

<h1 class="text-4xl font-bold text-center p-4 bg-gradient-to-r from-primary-500 via-secondary-700 to-tertiary-500">{title}</h1>

{@render children?.()}
