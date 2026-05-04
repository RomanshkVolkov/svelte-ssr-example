<script lang="ts">
    import CheckIcon from "@lucide/svelte/icons/check";
    import ChevronsUpDownIcon from "@lucide/svelte/icons/chevrons-up-down";
    import { tick } from "svelte";
    import Label from "@/components/ui/label/label.svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Command from "$lib/components/ui/command/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { cn } from "$lib/utils.js";

    const {
        name,
        defaultValue = "",
        options,
    }: {
        name: string;
        defaultValue?: string;
        options: Record<"label" | "value", string>[];
    } = $props();

    let open = $state(false);
    let value = $state(defaultValue);
    let triggerRef = $state<HTMLButtonElement>(null!);

    const selectedItem = $derived(
        options.find((f) => {
            console.debug("option", f, "value", value);
            return f.value === value;
        }),
    );

    // We want to refocus the trigger button when the user selects
    // an item from the list so users can continue navigating the
    // rest of the form with the keyboard.
    function closeAndFocusTrigger() {
        open = false;
        tick().then(() => {
            triggerRef.focus();
        });
    }
</script>

<Popover.Root bind:open>
    <Popover.Trigger bind:ref={triggerRef}>
        {#snippet child({ props })}
            <div class="flex flex-col gap-2">
                <Button
                    {...props}
                    variant="outline"
                    class="w-full justify-between capitalize"
                    role="combobox"
                    aria-expanded={open}
                >
                    {selectedItem?.label || "Select a option..."}
                    <ChevronsUpDownIcon class="opacity-50" />
                </Button>
                <input hidden {name} value={selectedItem?.value} />
            </div>
        {/snippet}
    </Popover.Trigger>
    <Popover.Content class="w-[200px] p-0">
        <Command.Root>
            <Command.Input placeholder="Search option..." />
            <Command.List>
                <Command.Empty>Not found option.</Command.Empty>
                <Command.Group value="frameworks">
                    {#each options as option (option.value)}
                        <Command.Item
                            value={option.value}
                            onSelect={() => {
                                console.debug(option);
                                value = option.value;
                                closeAndFocusTrigger();
                            }}
                        >
                            <CheckIcon
                                class={cn(
                                    value !== option.value &&
                                        "text-transparent",
                                )}
                            />
                            {option.label}
                        </Command.Item>
                    {/each}
                </Command.Group>
            </Command.List>
        </Command.Root>
    </Popover.Content>
</Popover.Root>
