<script lang="ts">
    import { EllipsisVertical, FolderInput } from "@lucide/svelte";
    import type { FileType } from "../types";
    import FileRow from "./FileRow.svelte";
    interface Props {
        files: FileType[];
    }
    let { files }: Props = $props();
</script>

<section>
    {#if files.length > 0}
        <ul class="list bg-base-100 rounded-box shadow-md">
            {#each files as file}
                <!-- <button
                    onclick={() => {
                        console.log(file);
                    }}
                    class="list-row">{file.name}</button
                > -->
                <FileRow {file} />
            {/each}
        </ul>
    {/if}
</section>
