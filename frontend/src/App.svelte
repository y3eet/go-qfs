<script lang="ts">
    import { onMount } from "svelte";
    import { getFiles } from "./lib/utils/api";
    import Header from "./lib/components/Header.svelte";
    import type { FileType } from "./lib/types";

    let files: FileType[] = $state([]);
    let fetchLoading = $state(true);
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
    {#if fetchLoading}
        <div class="flex items-center justify-center h-64">
            <span class="loading loading-spinner loading-lg"></span>
        </div>
    {/if}
    {#if files.length > 0}
        <ul class="list bg-base-100 rounded-box shadow-md">
            {#each files as file}
                <li class="list-row">{file.name}</li>
            {/each}
        </ul>
    {/if}
</section>
