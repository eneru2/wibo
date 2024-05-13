<script lang="ts">
  import { onMount } from "svelte"
  import wailsapp__runtime from "@wails/runtime";
  import { Open_Output_Dir_Dialog, ConvertImage, CloseApp, MinimiseApp} from '../wailsjs/go/main/App';
  import Output_Code from './lib/Output_Code.svelte'
  import Input_File from './lib/Input_File.svelte'
  import Output_Formats from './lib/Output_Formats.svelte'
  import Output_Width from './lib/Output_Width.svelte'
  import Output_Schema from './lib/Output_Schema.svelte';    
  import Folder_Search_Icon from './lib/Icons/Folder_Search_Icon.svelte'  

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
  
  let formats:Array<string> = [];
  let code: (internal_output_schema?:string) => string;

  function open_output_dir_dialog(){
    Open_Output_Dir_Dialog("l").then((res) => {
      output_dir = res
    });
  }

  function check_for_res_argument(){
    const regex = /\[res]/g;
    const replacement = output_width.toString();
    internal_output_schema = internal_output_schema.replaceAll(regex, replacement);
  }

  function close_app(){
    CloseApp()
  }

  function minimise_app(){
    MinimiseApp()
  }
   
  function ffmpeg(){ 
    internal_output_schema = output_schema
    check_for_res_argument();
    ConvertImage(input_file, formats, output_width.toString(), output_dir, internal_output_schema);
    default_string = code(internal_output_schema);
  };

</script>
<nav
  class="flex justify-center items-center py-1 text-2xl relative"
  style="--wails-draggable: drag;">
  <h1 class="">
    Wibo
  </h1>
  <icons
    style="--wails-draggable: no-drag;"
    class="flex items-center absolute right-8 gap-x-2">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- on:click={() => runtime.Hide()} -->
    <div
      on:click={minimise_app}
      class="h-6 w-6 flex justify-center pt-0.5 text-sm
      border-black border-2 rounded-md items-center cursor-pointer
      dark:border-background
      hover:bg-black hover:text-background
      dark:hover:bg-background dark:hover:text-black">-</div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- on:click={() => runtime.Quit()} -->
    <div
      on:click={close_app}
      class="h-6 w-6 flex justify-center pt-0.5 text-sm
      dark:border-background
      border-black border-2 rounded-md items-center hover:bg-black hover:text-white cursor-pointer
      dark:hover:bg-background dark:hover:text-black">X</div>
  </icons>
</nav>
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
        <Output_Code
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
        bind:formats/>
      <Output_Width
        bind:output_width
        bind:defaultResolutions/>
      <div class="flex flex-col"> 
        <h2>Output directory:</h2>
        <div class="flex gap-x-2 items-center">
          <button
            on:click={open_output_dir_dialog}
            class="border-2 border-black w-full min-h-7 max-h-14 break-all px-2 text-left overflow-auto
            dark:border-background">
            {#if output_dir != undefined}
              {output_dir}
            {/if}
          </button>
          <button
            on:click={open_output_dir_dialog}
            class="border-2 border-black px-3 min-h-7
            dark:border-background"><Folder_Search_Icon/></button>
        </div>
      </div>
      <Output_Schema
        bind:output_schema/>
    </output>
  </div>
  <div class="flex h-full mt-4 gap-x-4 justify-between items-center">
    <div class="h-full border-dotted border-zinc-700 border-2 relative p-3 w-1/3 gap-y-2 flex flex-col
        dark:border-background">
<!-- Extra options: -->
      <h2
        class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
        dark:bg-black text-nowrap">Convert:</h2>
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
        </div> -->
        <button
          class="self-start justify-self-start
          text-black border-black dark:text-background dark:border-background
          border-2 rounded-lg font-bold w-fit px-4 py-1 mt-4"
          on:click={ffmpeg}>Convert</button>
    </div>
    <output
      class="h-full border-dotted border-zinc-700 border-2 relative p-3 w-2/3 gap-y-2 flex flex-col
      dark:border-background">
      <h2
        class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
        dark:bg-black">Code:</h2>
      <textarea
        class="resize-none h-full flex-grow mt-1"
        spellcheck="false">{default_string}</textarea>
    </output>
  </div>
</app>