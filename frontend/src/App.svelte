<script lang="ts">
    import { onMount } from "svelte";
    import { getFiles } from "./lib/utils/api";
    import Header from "./lib/components/Header.svelte";
    import type { FileType } from "./lib/types";
    import FilesList from "./lib/components/FilesList.svelte";
    import { directory } from "./lib/state/directory.svelte";
    let files: FileType[] = $state([]);
    let fetchLoading = $state(true);

    async function fetchFiles(dir: string[]) {
        fetchLoading = true;
        const f = await getFiles(dir.join("/"));
        files = f;
        fetchLoading = false;
    }
    $effect(() => {
        fetchFiles(directory);
    });
</script>

<section>
    <Header />
    <div class="breadcrumbs text-sm mx-4 my-3">
        <ul>
            {#each directory as dir}
                <li>{dir}</li>
            {/each}
        </ul>
    </div>
    {#if fetchLoading}
        <div class="flex items-center justify-center h-64">
            <span class="loading loading-spinner loading-lg"></span>
        </div>
    {/if}
    <FilesList {files} />
</section>
