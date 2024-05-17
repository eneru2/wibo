<script lang="ts">
  import { onMount } from "svelte"
  import { ConvertImage, ReadConfig, InitialConfigRead } from '../wailsjs/go/main/App';
  import Output_Code_Options from './lib/Output_Code_Options.svelte'
  import Input_File from './lib/Input_File.svelte'
  import Output_Formats from './lib/Output/Output_Formats.svelte'
  import Output_Width from './lib/Output/Output_Width.svelte'
  import Output_Schema from './lib/Output/Output_Schema.svelte';    
  import Output_Quality from './lib/Output/Output_Quality.svelte'
  import OutputDirectory from "./lib/Output/Output_Directory.svelte";
  import Navbar from "./lib/Navbar.svelte";
  import Code from "./lib/Code.svelte";
  import ExtraOptions from "./lib/Extra_Options.svelte";
  
  let avif_crf;
  let webp_quality;
  let output_height:boolean;
  let default_string:string;
  let input_file:string;
  let output_width = 512;
  let output_dir:string;  
  let output_schema = ""; 
  let internal_output_schema:string;
  let defaultResolutions:Array<number> = [
    1920,
    1280,
    1024,
    768,
    512,
    384,
  ];

  let load_formats;
  let load_qualities;
  let load_code_options;
  InitialConfigRead()
  
  let formats:Array<string> = [];
  let code: (internal_output_schema?:string) => string;
  
  function readConfig() {
    return ReadConfig()
    .then(res => {
      load_formats(res);
      load_qualities(res);
      load_code_options(res);
      return res;
    });
  };
  
  let config = readConfig()

  function check_for_res_argument(){
    const regex = /\[res]/g;
    const replacement = output_width.toString();
    internal_output_schema = internal_output_schema.replaceAll(regex, replacement);
  };
   
  function ffmpeg(){ 
    internal_output_schema = output_schema
    check_for_res_argument();
    ConvertImage(input_file, formats, output_width.toString(),
    output_dir, internal_output_schema,
    avif_crf.toString(), webp_quality.toString(), output_height);
    default_string = code(internal_output_schema);
  };


  onMount(() => {
    // TODO: Create custom contextmenu
    document.addEventListener("contextmenu", (event) => {
      event.preventDefault()
    })
  })

</script>
<Navbar/>
<app
  class="flex flex-col flex-grow p-8 overflow-auto
  -outline-offset-[10px] outline outline-3 outline-black border-2 border-black
  dark:outline-background dark:border-background">
  <div class="flex gap-4">
    <div class="flex flex-col gap-y-4 w-1/2 h-full">
      <Input_File bind:input_file bind:output_schema/>
      <div class="flex h-full flex-col gap-y-4 border-dotted border-2 border-zinc-700 relative p-3 w-full dark:border-background">
        <h2
          class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
          dark:bg-black">Output code:</h2>
        <Output_Code_Options
          bind:load_config={load_code_options}
          bind:default_string
          bind:construct_output_string={code}
          bind:internal_output_schema
          bind:output_schema
          bind:formats/>
      </div>
    </div>
    <output
      class="border-dotted border-zinc-700 border-2 relative p-3 w-1/2 gap-y-2 flex flex-col
      dark:border-background">
      <h2
        class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
        dark:bg-black">Output:</h2>
      <Output_Formats
        bind:load_config={load_formats}
        bind:formats/>
      <Output_Width
        bind:output_height
        bind:output_width
        bind:defaultResolutions/>
      <OutputDirectory
        bind:output_dir/>
      <Output_Schema
        bind:output_schema/>
      <Output_Quality
        bind:load_config={load_qualities}
        bind:avif_crf
        bind:webp_quality/>
    </output>
  </div>
  <div class="flex h-full mt-4 gap-4 justify-between">
    <div class="w-1/3 flex flex-col gap-4 h-full">
      <ExtraOptions
        bind:output_height/>
      <div class="h-fit w-full border-dotted border-zinc-700 border-2 relative p-3  gap-y-2 flex flex-col
    dark:border-background">        
        <h2
          class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
          dark:bg-black text-nowrap">Convert:</h2>
        <button
          class="self-start justify-self-start
          text-black border-black dark:text-background dark:border-background
          border-2 rounded-lg font-bold w-fit px-4 py-1 mt-4"
          on:click={ffmpeg}>Convert</button>
      </div>
    </div>
    <Code      
      bind:default_string></Code>
  </div>
</app>