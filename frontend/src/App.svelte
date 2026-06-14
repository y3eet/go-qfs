<script lang="ts">
    import { onMount } from "svelte";
    import { getFiles } from "./lib/utils/api";
    import Header from "./lib/components/Header.svelte";
    import type { FileType } from "./lib/types";
    import FilesList from "./lib/components/FilesList.svelte";

    let files: FileType[] = $state([]);
    let fetchLoading = $state(true);
    let dir = "";

    async function fetchFiles() {
        fetchLoading = true;
        const f = await getFiles();
        files = f;
        fetchLoading = false;
    }
    $effect(() => {
        fetchFiles();
    });
</script>

<section>
    <Header />
    <div class="breadcrumbs text-sm mx-4 my-3">
        <ul></ul>
    </div>
    {#if fetchLoading}
        <div class="flex items-center justify-center h-64">
            <span class="loading loading-spinner loading-lg"></span>
        </div>
    {/if}
    <FilesList {files} />
</section>
