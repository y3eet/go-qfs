<script lang="ts">
    import { Plus, Upload } from "@lucide/svelte";
    import { uploadFile } from "../utils/api";
    import { directory } from "../state/directory.svelte";

    let modal: HTMLDialogElement;
    let files: FileList | undefined = $state();
    let isUploading: boolean = $state(false);

    function openModal() {
        modal.showModal();
    }

    function closeModal() {
        modal.close();
    }
    async function handleUpload() {
        if (!files || files.length === 0) return;
        isUploading = true;
        try {
            const promises = Array.from(files).map((file) =>
                uploadFile(file, directory.join("/")),
            );
            await Promise.all(promises);
            closeModal();
            isUploading = false;
        } catch (error) {
            console.error("Error uploading file:", error);
            isUploading = false;
            return;
        }
    }
</script>

<section>
    <button
        disabled={isUploading}
        class="btn btn-dash btn-primary"
        onclick={openModal}
    >
        {#if isUploading}
            <span class="loading loading-spinner loading-md"></span>
        {/if}
        <Plus /></button
    >

    <dialog bind:this={modal} class="modal">
        <div class="modal-box">
            <h3 class="text-lg font-bold">Upload File</h3>
            <input
                bind:files
                type="file"
                class="file-input file-input-bordered w-full mt-4"
                multiple
            />

            <div class="modal-action">
                <button class="btn" onclick={closeModal}> Close </button>
                <button class="btn btn-primary" onclick={handleUpload}>
                    <Upload size={16} />Upload
                </button>
            </div>
        </div>
    </dialog>
</section>
