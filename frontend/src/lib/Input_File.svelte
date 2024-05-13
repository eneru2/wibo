<script lang="ts">
  import { Open_Input_File_Dialog } from '../../wailsjs/go/main/App';
  import File_Search_Icon from './Icons/File_Search_Icon.svelte'
  
  export let output_schema:string;
  export let input_file:string;
  let input_file_name;

  function open_file_dialog(){
    Open_Input_File_Dialog("hello").then(res => {
      input_file = res
      const changeEscapesOnWindows = /\\/g
      res = res.replaceAll(changeEscapesOnWindows, "/")

      const regex = /(?<=\/)(?!.*\/).*/
      const delete_after_dot = /.(?<=\.)(?!.*\.).*/
      res = res.match(regex).toString()
      
      input_file_name = res
      res = res.replace(delete_after_dot, "")
      output_schema = res
    })
  };
</script>
<div
  class="flex gap-x-2 border-dotted border-2 border-zinc-700 relative p-3 w-full h-fit
  dark:border-background">
  <h2 class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
  dark:bg-black">Input:</h2>
  <h3 class="w-fit text-nowrap">Input file:</h3>  
  <button  
    on:click={open_file_dialog}
    class="w-full min-h-7 max-h-7 px-3 border-2 border-black text-nowrap overflow-auto text-left
    dark:border-background">
      {#if (input_file_name != undefined) }
        {input_file_name}
      {/if}
  </button>
  <button
    on:click={open_file_dialog}
    class="px-2 py-1 font-semibold w-fit
    border-2 border-black dark:border-background"><i><File_Search_Icon/></i></button>
</div>