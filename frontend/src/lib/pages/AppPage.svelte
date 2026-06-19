<script lang="ts">
    import { onMount } from "svelte";
    import { GetServerInfo } from "../../../wailsjs/go/main/App";
    import ShareLinkCard from "../components/ShareLinkCard.svelte";

    let info: Awaited<ReturnType<typeof GetServerInfo>> | undefined = $state();

    onMount(async () => {
        const res = await GetServerInfo();
        info = res;
    });
</script>

<section>
    {#if info}
        <ShareLinkCard link={info.host} />
    {/if}
</section>
