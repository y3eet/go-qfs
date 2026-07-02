<script lang="ts">
    import { House } from "@lucide/svelte";
    import FilesList from "../components/FilesList.svelte";
    import { directory } from "../state/directory.svelte";
    import { fileStore, fetchFiles } from "../state/files.svelte";

    $effect(() => {
        fetchFiles(directory);
    });
</script>

<section>
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
    {#if fileStore.fetchLoading}
        <div class="flex items-center justify-center h-64">
            <span class="loading loading-spinner loading-lg"></span>
        </div>
    {/if}
    <FilesList files={fileStore.files} />
</section>
