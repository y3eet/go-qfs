<script lang="ts">
    import { ExecCommand, GetServerInfo } from "../../../wailsjs/go/main/App";
    type Props = {
        info: Awaited<ReturnType<typeof GetServerInfo>>;
    };
    const { info }: Props = $props();
    const commands: Record<
        string,
        Record<string, { exposePort: string; closePort: string }>
    > = $derived({
        linux: {
            ufw: {
                exposePort: `sudo ufw allow ${info.port}`,
                closePort: `sudo ufw deny ${info.port}`,
            },
            firewalld: {
                exposePort: `sudo firewall-cmd --zone=public --add-port=${info.port}/tcp --permanent && sudo firewall-cmd --reload`,
                closePort: `sudo firewall-cmd --zone=public --remove-port=${info.port}/tcp --permanent && sudo firewall-cmd --reload`,
            },
            iptables: {
                exposePort: `sudo iptables -A INPUT -p tcp --dport ${info.port} -j ACCEPT`,
                closePort: `sudo iptables -D INPUT -p tcp --dport ${info.port} -j ACCEPT`,
            },
        },
        windows: {
            powershell: {
                exposePort: `Start-Process powershell -Verb RunAs -ArgumentList '-Command New-NetFirewallRule -DisplayName "Allow Port ${info.port} Private" -Direction Inbound -LocalPort ${info.port} -Protocol TCP -Action Allow -Profile Private'`,
                closePort: `Start-Process powershell -Verb RunAs -ArgumentList '-Command Remove-NetFirewallRule -DisplayName "Allow Port ${info.port} Private"'`,
            },
        },
    });
    const options = $derived(Object.keys(commands[info.os]));

    let openIndex = $state(0);
</script>

<section>
    <div class="card bg-base-100 w-full mt-2 shadow-md border border-base-200">
        <div class="card-body p-3 gap-2">
            <div class="flex items-center justify-between">
                <h2 class="card-title text-base">Firewall Commands</h2>
                <span class="badge badge-outline badge-sm">{info.os}</span>
            </div>

            <div class="pr-1 flex flex-col gap-1.5">
                {#each options as option, i}
                    <div
                        class="collapse collapse-arrow bg-base-200/50 border border-base-300 rounded-md"
                    >
                        <input
                            type="radio"
                            name="firewall-tool-accordion"
                            checked={openIndex === i}
                            onchange={() => (openIndex = i)}
                        />
                        <div class="collapse-title font-medium py-2 min-h-0">
                            {option}
                        </div>
                        <div class="collapse-content">
                            <div class="flex flex-col gap-2 pt-1">
                                <div class="flex items-start gap-2">
                                    <code
                                        class="flex-1 bg-base-300/60 rounded px-2 py-1.5 overflow-x-auto whitespace-pre"
                                        >{commands[info.os][option]
                                            .exposePort}</code
                                    >
                                    <button
                                        class="btn btn-primary shrink-0"
                                        onclick={() =>
                                            ExecCommand(
                                                commands[info.os][option]
                                                    .exposePort,
                                            )}
                                    >
                                        Expose Port
                                    </button>
                                </div>
                                <div class="flex items-start gap-2">
                                    <code
                                        class="flex-1 bg-base-300/60 rounded px-2 py-1.5 overflow-x-auto whitespace-pre"
                                        >{commands[info.os][option]
                                            .closePort}</code
                                    >
                                    <button
                                        class="btn btn-secondary shrink-0"
                                        onclick={() =>
                                            ExecCommand(
                                                commands[info.os][option]
                                                    .closePort,
                                            )}
                                    >
                                        Close Port
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</section>
