<script lang="ts">
  import { Open_Output_Dir_Dialog } from '../../wailsjs/go/main/App';
  import Folder_Search_Icon from './Icons/Folder_Search_Icon.svelte'
  
  export let output_dir:string;  

  function open_output_dir_dialog(){
    Open_Output_Dir_Dialog().then((res) => {
      output_dir = res
    });
  }
</script>
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
      dark:border-background">
        <Folder_Search_Icon />
    </button>
  </div>
</div>
