<script lang="ts">
    import { onMount } from "svelte";

  export let disable_tooltips;
  export let description = "";
  export let default_behaviour = "";
  export let options = "";
  export let example = "";
  let tooltip_on_screen: boolean;

  onMount(() => {
    
  })

</script>

<!-- svelte-ignore a11y-mouse-events-have-key-events -->
{#if !disable_tooltips}
  <div
    class="w-5 h-5 ml-1 relative"
    on:mouseover={() => (tooltip_on_screen = true)}
    on:mouseleave={() => (tooltip_on_screen = false)}>
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
    {#if tooltip_on_screen}
      <tooltip class="absolute left-5 top-0 max-w-96 w-96 px-6 py-3 text-black bg-background border border-black z-10">
        <h3>DESCRIPTION</h3>
        <p>{description}</p>
        {#if default_behaviour !== ""}    
          <h3>DEFAULT</h3>
          <p>{default_behaviour}</p>
        {/if}
        {#if options !== ""}
          <h3>OPTIONS</h3>
          <div
            contenteditable="false"
            bind:innerHTML={options}></div>
        {/if}
        {#if example !== ""}
          <h3>EXAMPLE</h3>
          <div
            contenteditable="false"
            bind:innerHTML={example}></div>
        {/if}
      </tooltip>   
    {/if}
  </div>
{/if}


<style lang="postcss">
  h3 {
    text-decoration: underline;
    font-weight: bold;
  }
  p {
    @apply mb-1;
  }
</style>
