<script lang="ts">
  import { Copy_To_Clipboard } from "../../wailsjs/go/main/App";  

  export let default_string: string;
  let success_message: HTMLElement;
  let error_message: HTMLElement;

  function copy_to_clipboard () {    
    Copy_To_Clipboard(default_string).then(() => {      
      success_message.style.opacity = "100%";

      setTimeout(() => {
        success_message.style.opacity = "0%";
      }, 4000);

    }).catch(() => {
      error_message.style.opacity = "100%";

      setTimeout(() => {
        error_message.style.opacity = "0%";
      }, 4000);
    });
  };
</script>

<output
  class="h-full border-dotted border-zinc-700 border-2 relative p-3 pb-1.5 w-2/3 flex flex-col
  dark:border-background">
  <h2
    class="absolute -top-3 left-1/2 -translate-x-1/2 bg-background pl-1.5
    dark:bg-black">
    Code:
  </h2> 
  <div class="flex-grow h-full mt-1 relative">
    <textarea 
      class="resize-none h-full" spellcheck="false" bind:value={default_string}></textarea>
    <button
      class="absolute right-4 top-4 w-fit h-fit px-2 py-1 border-2 border-white text-white
      dark:text-black dark:border-black"
      on:click={copy_to_clipboard}>Copy</button>
    </div>
      <success_message
        bind:this={success_message}
        class="text-green-500 h-fit ml-1 absolute bottom-2 right-5 transition-all"
        style="opacity: 0;">Text copied to clipboard successfuly!</success_message>
      <success_message
      bind:this={error_message}
      class="text-red-500 h-fit ml-1 absolute bottom-2 right-5 transition-all"
      style="opacity: 0;">There was an error copying to the clipboard</success_message>
</output>

<style lang="postcss">
  textarea {
    scrollbar-width: none;
    -ms-overflow-style: none;
  }  
</style>
