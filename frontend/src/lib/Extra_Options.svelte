<script>
  let disable_tooltips;
  export let output_height;
</script>

<div class="h-full w-full border-dotted border-zinc-700 border-2 relative p-3  gap-y-2 flex flex-col
dark:border-background">
  <h2
  class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
  dark:bg-black text-nowrap">Extra options:</h2>
  <div class="py-2 flex flex-col gap-y-4">
    <div class="flex items-center gap-x-4">
      <h3>Disable tooltips:</h3>
      <input 
      bind:checked={disable_tooltips}
      type="checkbox">
    </div>
    <div class="flex items-center gap-x-4">
      <h3>Output height:</h3>
      <input 
      bind:checked={output_height}
      type="checkbox">
    </div>
  </div>
</div>
<!-- Extra options:
  Convert: -->      
      <!-- 
        default output formats
        AVIF crf
        webp quality
        use output height instead
        custom width/heights
      -->