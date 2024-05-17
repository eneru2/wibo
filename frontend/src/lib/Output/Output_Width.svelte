<script>
  export let defaultResolutions;
  export let output_width;
  export let output_height;

  // $: button_equals_output_width = output_width === 512 ? 'bg-black' : '';
</script>
  <h2>Output
    {#if output_height}
      height:
      {:else}
      width:
    {/if}
  </h2>
  <div class="flex items-center gap-x-2">
    <input
      bind:value={output_width}
      class="px-2 py-1"
      min="16"
      max="3840"
      type="number"
      on:keydown={(evt) => ["e", "E", "+", "-"].includes(evt.key) && evt.preventDefault()}>    
    <!-- <input
      bind:value={output_width}
      type="range"
      min="16"
      max="3840"/> -->
    <default
      class="grid grid-cols-6 gap-0.5 border-2 border-black bg-black w-full
      dark:border-background dark:bg-background">
      {#each defaultResolutions as res}
        <button
          class="flex justify-center items-center p-0.5 bg-background aspect-square text-sm
          dark:bg-black"
          on:click={() => (output_width = res)}>{res}</button>
      {/each}
    </default>      
  </div>