<script lang="ts">
  import { onMount } from "svelte";
  import RangeLayout from "../Layouts/Range_Layout.svelte";
  import TooltipLayout from "../Layouts/Tooltip_Layout.svelte";  
  import { WriteOutputQualities } from "../../../wailsjs/go/main/App";

  export function load_config(config) {
    avif_crf = config.saved_state.avif_crf;
    webp_quality = config.saved_state.webp_quality;
  }

  let write_config = (format:string, value:number) => {
    switch(format){
      case "avif":
        WriteOutputQualities("avif_crf", value);
        break;
      case "webp":
        WriteOutputQualities("webp_quality", value)
        break;
    };
  };
  
  export let avif_crf = 25
  export let webp_quality = 83

  let avif_tooltip = false
  let webp_tooltip = false

  let avif_description = `This option will allow you
  to customize the quality of the resulting AVIF image.`
  let avif_default_behaviour = `The default value is set to 25,
    which should heavily reduce file size while maintaining good quality.
    Lesser values improve quality at the expense of file size. The value can range from
    0 to 63`

  let webp_description = `This option will allow you
  to customize the quality of the resulting WEBP image.`
  let webp_default_behaviour = `The default value is set to 83,
    which should heavily reduce file size while maintaining good quality.
    Greater values improve quality at the expense of file size. The value can range from
    0 to 100`

</script>

<div>
  <h2>Output quality:</h2>
  <div class="flex flex-col gap-y-2 pt-1">
    <div class="flex items-center">
      AVIF
      <div class="w-6 h-6 relative ml-2"
      on:mouseenter={() => avif_tooltip = true}
      on:mouseleave={() => avif_tooltip = false}>
        <svg
          xmlns="http://www.w3.org/2000/svg"          
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round" >
          <circle cx="12" cy="12" r="10"/>
          <path d="M12 16v-4"/><path d="M12 8h.01"/>
        </svg>
        {#if avif_tooltip}
          <TooltipLayout
            bind:description={avif_description}
            bind:default_behaviour={avif_default_behaviour}
            />
        {/if}
      </div>:
      <RangeLayout
        bind:value_change={write_config}
        bind:value={avif_crf}
        name={"avif"}
        container_width={160}
        maxlength={2}
        max={63}/>
    </div>
      <!-- <input
        bind:value={avif_value}
        type="range" min="0" max="63" /> -->
    <div class="flex">
      WEBP
      <div class="w-6 h-6 relative ml-2"
      on:mouseenter={() => webp_tooltip = true}
      on:mouseleave={() => webp_tooltip = false}>
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
      {#if webp_tooltip}
        <TooltipLayout
          bind:description={webp_description}
          bind:default_behaviour={webp_default_behaviour}
          />
      {/if}
    </div>:
      <RangeLayout
        bind:value_change={write_config}
        bind:value={webp_quality}
        name={"webp"}
        container_width={160}
        maxlength={3}
        max={100}/>
    </div>    
  </div>
</div>