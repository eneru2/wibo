<script lang="ts">
  import { Open_Output_Dir_Dialog, ConvertAnyFormatToPng} from '../wailsjs/go/main/App';
  import Output_Code from './lib/Output_Code_Options.svelte'
  import Input_File from './lib/Input_File.svelte'
  import Output_Formats from './lib/Output_Formats.svelte'
  import Output_Width from './lib/Output_Width.svelte'
  import Output_Schema from './lib/Output_Schema.svelte';
  import Flowers from './lib/Flowers.svelte';
    import OutputSchema from './lib/Output_Schema.svelte';

  let name;
  let greeting = "";

  let input_file;
  let output_width = 512;
  let output_dir;
  let output_schema = ""; 

  let defaultResolutions = [
    1920,
    1536,
    1280,
    1024,
    512,
    256,
    128
  ]

  export let formats = [    
  ]

  function open_output_dir_dialog(){
    Open_Output_Dir_Dialog("l").then(res => output_dir = res)
  }


  function greet() {
    Greet(name).then((result) => {
      greeting = result;
    });
  }

  function check_for_res_argument(){
    const regex = /\[res]/g;
    const replacement = output_width.toString();

    output_schema = output_schema.replaceAll(regex, replacement)    
  }

  function ffmpeg(){
    check_for_res_argument();
    ConvertAnyFormatToPng(input_file, formats, output_width.toString(), output_dir, output_schema);
    // ConvertAnyFormatToPng("G:\\Andrés\\Download\\ffmpeg-7.0-essentials_build\\bin\\hola.avif")
    //   .then(res => alert(res))
  }

</script>
<!-- <h1>Image Optimizer for the Web</h1> -->
<!-- <h1 class="text-5xl text-blue-imperial font-bold mb-4">Wabi</h1> -->

<app class="flex flex-col gr h-screen px-8 py-12 border-[12px] border-red-brown overflow-auto">
  <!-- <Flowers 
    x = "2"
    y = "5"
    rotation = "90"/> -->
  <!-- <img class="absolute w-screen text-red-brown fill-red-brown bg-red-brown" src="/src/assets/flowers-thin-ornamental-frame-svgrepo-com.svg" alt=""> -->
  <h1 class="text-5xl text-blue-imperial font-bold mb-4 pb-3 border-b-2 border-b-blue-imperial">Wibo</h1>
  <button on:click={() => alert(input_file)}>caca</button>
  <Input_File bind:input_file/>
  <Output_Formats bind:formats/>
  <Output_Width bind:output_width bind:defaultResolutions/>
  <label class="flex flex-col"> 
    <h2>Output directory</h2>
    <div>
      {#if output_dir != undefined}        
        {output_dir}
      {/if}
    </div>
    <button
      on:click={open_output_dir_dialog}>click</button>
    <!-- <input placeholder={output_dir} on:click={open_output_dir_dialog} type="text"> -->
  </label>
  <Output_Schema bind:output_schema/>
  <button
    class="bg-blue-imperial text-background font-bold w-fit px-4 py-1 mt-4"
    on:click={ffmpeg}>Convert</button>
  <details class="overflow-auto">
    <summary>Options</summary>    
    <Output_Code/>
  </details>
</app>

<style lang="postcss">

  .gr {    
    background: url("/src/assets/img-noise-361x370.png");
  }

</style>