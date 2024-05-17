<script lang="ts">
  import { onMount } from "svelte"
  import { ConvertImage, ReadConfig, InitialConfigRead } from '../wailsjs/go/main/App';
  import Output_Code_Options from './lib/Output_Code_Options.svelte'
  import Input_File from './lib/Input_File.svelte'
  import Output_Formats from './lib/Output_Formats.svelte'
  import Output_Width from './lib/Output_Width.svelte'
  import Output_Schema from './lib/Output_Schema.svelte';    
  import Output_Quality from './lib/Output_Quality.svelte'
  import OutputDirectory from "./lib/Output_Directory.svelte";
  import Navbar from "./lib/Navbar.svelte";
    import Code from "./lib/Code.svelte";

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

  InitialConfigRead("/index.html")
  
  let formats:Array<string> = [];
  let code: (internal_output_schema?:string) => string;
  let config
  
  function readConfig() {
    let result = ReadConfig()
      .then(res => config = res)
    //alert(JSON.stringify(result))
  };

  readConfig()

  function check_for_res_argument(){
    const regex = /\[res]/g;
    const replacement = output_width.toString();
    internal_output_schema = internal_output_schema.replaceAll(regex, replacement);
  };
   
  function ffmpeg(){ 
    internal_output_schema = output_schema
    check_for_res_argument();
    ConvertImage(input_file, formats, output_width.toString(), output_dir, internal_output_schema);
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
    <div class="flex flex-col gap-y-4 w-1/2">
      <Input_File bind:input_file bind:output_schema/>
      <div class="flex flex-col gap-y-4 border-dotted border-2 border-zinc-700 relative p-3 w-full h-fit dark:border-background">
        <h2
          class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
          dark:bg-black">Output code:</h2>
        <Output_Code_Options
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
        bind:config
        bind:formats/>
      <Output_Width
        bind:output_width
        bind:defaultResolutions/>
      <OutputDirectory
        bind:output_dir/>
      <Output_Schema
        bind:output_schema/>
      <Output_Quality
        />
    </output>
  </div>
  <div class="flex h-full mt-4 gap-x-4 justify-between items-center">
    <div class="h-full border-dotted border-zinc-700 border-2 relative p-3 w-1/3 gap-y-2 flex flex-col
        dark:border-background">
<!-- Extra options: -->
      <h2>Extra options:</h2>
      <h2
        class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
        dark:bg-black text-nowrap">Convert:</h2>
        <div class="flex items-center gap-x-4">
          <h3>Disable tooltips:</h3>
          <input type="checkbox" name="" id="">
        </div>
         <!-- <div>
          Default output formats:
          <label class="flex items-center gap-x-4">
            AVIF
            <input type="checkbox" name="" id="">
          </label>
          <label class="flex items-center gap-x-2">
            WEBP
            <input type="checkbox" name="" id="">
          </label>
          <label class="flex items-center gap-x-2">
            JPG
            <input type="checkbox" name="" id="">
          </label>
        </div>--><!-- 
        <div>        
          Default output widths:
          <label class="flex items-center gap-x-4">
            AVIF
            <input type="checkbox" name="" id="">
          </label>
          <label class="flex items-center gap-x-2">
            WEBP
            <input type="checkbox" name="" id="">
          </label>
          <label class="flex items-center gap-x-2">
            JPG
            <input type="checkbox" name="" id="">
          </label>
        </div>
      default output formats
      AVIF crf
      webp quality
      use output height instead
      custom width/heights
      -->    
        <button
          class="self-start justify-self-start
          text-black border-black dark:text-background dark:border-background
          border-2 rounded-lg font-bold w-fit px-4 py-1 mt-4"
          on:click={ffmpeg}>Convert</button>
    </div>
    <Code      
      bind:default_string></Code>
  </div>
</app>