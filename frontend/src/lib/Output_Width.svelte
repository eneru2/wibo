<script>
  export let defaultResolutions;
  export let output_width;

  // $: button_equals_output_width = output_width === 512 ? 'bg-black' : '';
</script>
  <h2>Output width:</h2>
  <div class="flex items-center gap-x-2">
    <input
      bind:value={output_width}
      class="border-2 border-black bg-transparent px-2 py-1
      dark:border-background"
      min="16"
      max="3840"
      type="number"/>
    <!-- <input
      bind:value={output_width}
      type="range"
      min="16"
      max="3840"/> -->
    <default
      class="grid grid-cols-6  gap-0.5 border-2 border-black bg-black
      dark:border-background dark:bg-background">
      {#each defaultResolutions as res}
        <button
          class="flex justify-center items-center p-0.5 bg-background aspect-square text-sm
          dark:bg-black"
          on:click={() => (output_width = res)}>{res}</button>
      {/each}
    </default>      
  </div>