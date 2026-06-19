<script lang="ts">
    import { onMount } from "svelte";
    import Header from "./lib/components/Header.svelte";
    import { GetServerInfo } from "../wailsjs/go/main/App";
    import FilesPage from "./lib/pages/FilesPage.svelte";

    let status = $state("");

    onMount(() => {
        if (window.go) {
            GetServerInfo().then((info) => {
                status = JSON.stringify(info);
            });
        } else {
            console.warn(
                "Wails bindings not available (not running in Wails webview)",
            );
        }
    });
</script>

<section>
    <Header />
    {#if window.go}
        <div class="text-sm text-gray-500 mx-4 my-2">Server Info: {status}</div>
    {:else}
        <FilesPage />
    {/if}
</section>
