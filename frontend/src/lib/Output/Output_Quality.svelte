<script lang="ts">
  import RangeLayout from "../Layouts/Range_Layout.svelte";
  import TooltipLayout from "../Layouts/Tooltip_Layout.svelte";  
  import { WriteOutputQualities } from "../../../wailsjs/go/main/App";

  export let disable_tooltips;
  let on_mount;
  let on_mount2;
  export function load_config(config) {
    avif_crf = config.saved_state.avif_crf;
    webp_quality = config.saved_state.webp_quality;    
    on_mount();
    on_mount2();
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
      AVIF:
      <TooltipLayout
            
            bind:disable_tooltips
            bind:description={avif_description}
            bind:default_behaviour={avif_default_behaviour}
            />
      <RangeLayout
        bind:change_input={on_mount}
        bind:value_change={write_config}
        bind:value={avif_crf}
        name={"avif"}
        container_width={160}
        maxlength={2}
        max={63}/>
    </div>
    <div class="flex items-center">
      WEBP:
      <TooltipLayout
        bind:disable_tooltips
        bind:description={webp_description}
        bind:default_behaviour={webp_default_behaviour}
        />
      <RangeLayout
        bind:change_input={on_mount2}
        bind:value_change={write_config}
        bind:value={webp_quality}
        name={"webp"}
        container_width={160}
        maxlength={3}
        max={100}/>
    </div>    
  </div>
</div>