<script lang="ts">
    import { onMount } from "svelte";
    import { getFiles } from "./lib/utils/api";
    import Header from "./lib/components/Header.svelte";
    import type { FileType } from "./lib/types";
    import FilesList from "./lib/components/FilesList.svelte";
    import { directory } from "./lib/state/directory.svelte";
    import { House } from "@lucide/svelte";

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
        <ul class="flex items-center">
            <li>
                <button onclick={() => directory.splice(0, directory.length)}>
                    <House size={18} />
                </button>
            </li>
            {#each directory as dir, i}
                <li>
                    <button
                        onclick={() =>
                            directory.splice(i + 1, directory.length - i)}
                        class="btn btn-link btn-sm text-sm text-primary p-0"
                        >{dir}</button
                    >
                </li>
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
