<script lang="ts">
  import { onMount } from 'svelte'
  import { ReadConfig, WriteOutputFormats } from '../../wailsjs/go/main/App';
  
  interface config {
    saved_state: {
      indenting_amount: number
      include_mime_types: boolean
      use_absolute_paths: boolean
      use_new_lines_for_properties: boolean
      
      output_avif: boolean
      output_webp: boolean
      output_jpg: boolean
  
      avif_crf: number
      webp_quality: number
    },
    options: {
      output_resolutions: Array<string>
      use_height_for_output: boolean
      disable_tooltips: boolean
    }
  }

  export let formats:Array<string> = [];
  export let config:any;
  
  let avif = false;
  let webp = false;
  let jpg = false;
  
  function load_config() {
    avif = config.saved_state.output_avif;
    webp = config.saved_state.output_webp;
    jpg = config.saved_state.output_jpg;
  }

  function readConfig(){
    ReadConfig().then(res => {
      config = res;
      load_config();
    })    
  };

  function write_config(format, state){
    let send_format;
    switch(format){
      case "avif":
        send_format = "output_avif";
        break;
      case "webp":
        send_format = "output_webp";
        break;
      case "jpg":
        send_format = "output_jpg";
        break;
    };
    WriteOutputFormats(send_format, state)
  };

  readConfig()

  function check_current_formats(format:string){
    switch(format){
      case "avif":
        avif ? add_format_to_array(format) : delete_array_element(format);
        avif ? write_config(format, true) : write_config(format, false);
        break;
      case "jpg":
        jpg ? add_format_to_array(format) : delete_array_element(format);
        jpg ? write_config(format, true) : write_config(format, false);
        break;
      case "webp":
        webp ? add_format_to_array(format) : delete_array_element(format);
        webp ? write_config(format, true) : write_config(format, false);
        break;
    };
  };

  function delete_array_element(format:string){
    for (let i = 0; i <= formats.length - 1; i++) {
      if(format === formats[i]) {
        formats.splice(i, 1);
      }    
    };
  };
    
  function add_format_to_array(format:string){
    const findElement = formats.filter((el) => el.includes(format));
    if(findElement.toString() === ""){
      formats.push(format);
    };
    sort_array();
  }; 

  function sort_array(){
    switch(formats.length){
      case 1:        
        return;
      default:
        formats.sort();
        for (let i = 0; i < formats.length - 1; i++) {
          if(formats[i] === "jpg"){
            formats.splice(i, 1);
            formats.push("jpg");
          }
        }
        break;
      }
  };

  onMount(() => {
  })
</script>

<h2>Output formats:</h2>
<div class="container flex items-center gap-x-4">
  <label
    class="flex items-center gap-x-2 cursor-pointer">
    <input
      type="checkbox"        
      bind:checked={avif}
      on:change={() => check_current_formats("avif")}/>
      AVIF
  </label>
  <label
    class="flex items-center gap-x-2 cursor-pointer">
  <input
    type="checkbox"
    bind:checked={webp}
    on:change={() => check_current_formats("webp")}/>
    WEBP
  </label>
  <label
    class="flex items-center gap-x-2 cursor-pointer">
    <input
      type="checkbox"
      bind:checked={jpg}
      on:change={() => check_current_formats("jpg")}/>
      JPG
  </label>
</div>