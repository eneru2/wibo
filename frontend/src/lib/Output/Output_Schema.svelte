<script lang="ts">
    import TooltipLayout from "../Layouts/Tooltip_Layout.svelte";

  
  let hovered;
  export let output_schema = "";
  export let disable_tooltips;
  let description = "Allows to create custom file names for the output files.";
  let default_behaviour = "If this field is empty the output will default to the input's file name."
  let options = `
  [res]<br/>
  <p class="pl-4">
    If you want to use the target res in the name of your output, you can do
    it with this argument, as an example, <b>flying_bird-[res]x</b> would
    output:</p>
      <p class="pl-8">
        - flying_bird-512x.jpg<br />
        - flying_bird-512x.webp<br />
        - flying_bird-512x.avif</p>
  `
</script>

<div class="relative">
  <div class="flex items-center mb-2">
    <h2>Output name schema:</h2>
    <TooltipLayout
      bind:disable_tooltips
      bind:options
      bind:description={description}
      bind:default_behaviour={default_behaviour}/>
  </div>
  <input
    spellcheck="false"
    bind:value={output_schema}
    class="w-fit border-2 border-black bg-transparent px-2"
    type="text">
</div>