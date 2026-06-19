<script lang="ts">
    import { EllipsisVertical, FolderInput } from "@lucide/svelte";
    import UploadModal from "./UploadModal.svelte";
</script>

<div class="navbar bg-base-100 shadow-sm">
    <div class="flex-1">
        <p class="btn btn-ghost text-xl"><FolderInput />Go QFS</p>
    </div>
    <div class="flex-none">
        {#if window.go}
            <button class="btn btn-square btn-ghost"
                ><EllipsisVertical /></button
            >
        {:else}
            <UploadModal />
        {/if}
    </div>
</div>
