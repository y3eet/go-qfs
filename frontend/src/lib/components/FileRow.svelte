<script lang="ts">
    import { Download, EllipsisVertical } from "@lucide/svelte";
    import type { FileType } from "../types";
    import { downloadFile } from "../utils/api";
    import { MIMES } from "../constants";
    import { directory } from "../state/directory.svelte";
    interface Props {
        file: FileType;
    }
    let { file }: Props = $props();

    async function handleDowndloadFile(file: string) {
        try {
            const formattedDirectory =
                directory.length > 0 ? directory.join("/") + "/" : "";
            await downloadFile(formattedDirectory + file);
        } catch (e) {
            console.error(e);
        }
    }
    const mime = $derived(
        MIMES[file.is_dir ? "dir" : file.ext] ?? MIMES["default"],
    );
    const Icon = $derived(mime?.icon);
</script>

<section>
    <li
        class="list-row relative flex items-center justify-between px-2 md:px-4 lg:px-6 hover:bg-base-200 rounded-lg"
    >
        {#if file.is_dir}
            <button
                aria-labelledby="file-row-{file.name}"
                type="button"
                onclick={() => directory.push(file.name)}
                class="absolute inset-0 rounded-lg"
            ></button>
        {/if}

        <div class="relative z-10 flex gap-2 pointer-events-none">
            <Icon />
            <div class="flex flex-col items-start">
                <span class="font-bold">{file.name}</span>
                <span class="text-xs text-gray-500">{file.size} bytes</span>
            </div>
        </div>

        {#if !file.is_dir}
            <div class="relative z-10 dropdown dropdown-bottom dropdown-end">
                <button tabindex="0" class="btn btn-ghost hover:cursor-pointer">
                    <EllipsisVertical size={18} />
                </button>
                <ul
                    tabindex="-1"
                    class="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"
                >
                    <button
                        onclick={() => handleDowndloadFile(file.name)}
                        class="btn btn-ghost"
                    >
                        <Download size={18} />Download
                    </button>
                </ul>
            </div>
        {/if}
    </li>
</section>
