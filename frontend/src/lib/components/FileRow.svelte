<script lang="ts">
    import { Download, EllipsisVertical } from "@lucide/svelte";
    import type { FileType } from "../types";
    import { downloadFile } from "../utils/api";
    interface Props {
        file: FileType;
    }
    let { file }: Props = $props();

    async function handleDowndloadFile(file: string) {
        try {
            await downloadFile(file);
        } catch (e) {
            console.error(e);
        }
    }
</script>

<section>
    <li class="list-row flex items-center justify-between px-6 hover:bg-base-200 rounded-lg">
        <button onclick={() => {}}>
            {file.name}
        </button>
        {#if !file.is_dir}
            <div class="dropdown dropdown-bottom dropdown-end">
                <button
                    tabindex="0"
                    class="btn btn-ghost hover:cursor-pointer"
                >
                    <EllipsisVertical size={18} />
                </button>
                <ul
                    tabindex="-1"
                    class="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"
                >
                    <button
                        onclick={() => handleDowndloadFile(file.name)}
                        class="btn btn-ghost"
                        ><Download size={18} />Download</button
                    >
                </ul>
            </div>
        {/if}
    </li>
</section>
