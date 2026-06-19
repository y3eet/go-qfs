<script lang="ts">
    import Header from "./lib/components/Header.svelte";
    import FilesPage from "./lib/pages/FilesPage.svelte";
    import AppPage from "./lib/pages/AppPage.svelte";
</script>

<section>
    <Header />
    {#if window.go}
        <AppPage />
    {:else}
        <FilesPage />
    {/if}
</section>
