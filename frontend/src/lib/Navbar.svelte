<script lang="ts">
  import { CloseApp, MinimiseApp } from "../../wailsjs/go/main/App";

  function close_app(){
    CloseApp()
  }

  function minimise_app(){
    MinimiseApp()
  }
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