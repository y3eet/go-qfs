<script lang="ts">
    import { Check, Copy, SquareArrowOutUpRight } from "@lucide/svelte";
    import QRCode from "qrcode";
    import { onMount } from "svelte";

    type Props = {
        link: string;
    };
    const { link }: Props = $props();

    let copied = $state(false);
    let qr = $state("");

    async function generate() {
        qr = await QRCode.toDataURL(link);
    }
    onMount(generate);
    async function copyLink() {
        try {
            await navigator.clipboard.writeText(link);
            copied = true;
            setTimeout(() => (copied = false), 1800);
        } catch {
            copied = false;
        }
    }
</script>

<section class="flex justify-center">
    <div class="card bg-base-100 w-96 shadow-md border border-base-200">
        <figure class="bg-base-200/40 pt-6">
            <div class="bg-white p-3 rounded-xl shadow-sm">
                <img
                    src={qr}
                    alt="QR code for share link"
                    class="size-44 object-contain"
                />
            </div>
        </figure>
        <div class="card-body gap-3">
            <h2 class="card-title text-base">Share this link</h2>

            <div class="flex items-center gap-2">
                <a
                    href={link}
                    target="_blank"
                    rel="noopener noreferrer"
                    class="link link-hover text-sm text-base-content/70 truncate flex-1"
                    title={link}
                >
                    {link}
                </a>
            </div>

            <div class="card-actions mt-1">
                <button
                    class="btn btn-primary btn-sm flex-1"
                    onclick={copyLink}
                    disabled={copied}
                >
                    {#if copied}
                        <Check size={16} />
                        Copied
                    {:else}
                        <Copy size={16} />
                        Copy link
                    {/if}
                </button>
                <a
                    href={link}
                    target="_blank"
                    rel="noopener noreferrer"
                    class="btn btn-ghost btn-sm"
                >
                    <SquareArrowOutUpRight size={16} />
                    Open
                </a>
            </div>
        </div>
    </div>
</section>
