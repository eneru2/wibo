<script lang="ts">
  import Tooltip from "../Tooltip.svelte";
  let hovered;
  export let output_schema = "";  
</script>

<div class="relative">
  <div class="flex gap-x-2">
    <h2 class="">
    Output name schema:
    </h2> 
  <!-- svelte-ignore a11y-mouse-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div
    class="w-6 h-6"
    on:mouseover={() => (hovered = true)}
    on:mouseleave={() => (hovered = false)}>
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
      {#if hovered}
        <Tooltip/>
      {/if}
    </div>
  </div>
  <input
    spellcheck="false"
    bind:value={output_schema}
    class="w-fit border-2 border-black bg-transparent px-2"
    type="text">
</div>