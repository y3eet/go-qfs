<script lang="ts">
    import { onMount } from "svelte";
    import { ExecCommand, GetServerInfo } from "../../../wailsjs/go/main/App";
    import ShareLinkCard from "../components/ShareLinkCard.svelte";
    import ManageFirewallCard from "../components/ManageFirewallCard.svelte";

    let info: Awaited<ReturnType<typeof GetServerInfo>> | undefined = $state();

    onMount(async () => {
        const res = await GetServerInfo();
        info = res;
    });
    function handleExec() {
        ExecCommand("sudo ufw status");
    }
</script>

<section>
    {#if info}
        <ShareLinkCard link={info.host} />
        <ManageFirewallCard {info} />
    {/if}
</section>
