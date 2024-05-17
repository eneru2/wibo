<script lang="ts">
  import { onMount } from "svelte";    

  let range:undefined | HTMLInputElement;
  let percentage;
  let slider_thumb;
  export let value;
  let is_pressed = false;  
  export let container_width;  
  export let max = 100;
  const range_min = 0;
  const thumb_width = 16;
  

  function on_click(event){
    
    const range_container_size = range.getBoundingClientRect();  
    const x = event.clientX - range_container_size.left; 

    let track_width = Math.round(x*100/container_width);
    let thumb_pos: number;
    value = Math.round(x*max/container_width);

    if (x > container_width){
      thumb_pos = Math.round(container_width*(100-(thumb_width))/container_width);
    } else {
      thumb_pos = Math.round(x*(100-(thumb_width))/container_width);
    }

    if (track_width < 0){
      track_width = 0
    } if (track_width > 100) {
      track_width = 100
    }
    
    if (thumb_pos < 0){
      thumb_pos = 0
    } if (thumb_pos > 100) {
      thumb_pos = 100
    }

    if (value < 0){
      value = 0
    } if (value > max) {
      value = max
    }   

    percentage.style.width = track_width + "%";
    slider_thumb.style.left = thumb_pos + "%";
  }
  
  function on_mouse_down(event){    
    on_click(event)
    is_pressed = true

    document.addEventListener("mousemove", on_mouse_move)

    document.addEventListener("mouseup", () => { 
      is_pressed = false, 
      document.removeEventListener("mousemove",on_mouse_move)
    }, {once: true})
  }

  function on_mouse_move(event){
    if(is_pressed) {
      on_click(event)
    }
  }

  onMount(() => {
    range.style.width = container_width + "px"
    slider_thumb.style.left = value + "%"
  })

</script>

<!-- on:mousemove={on_mouse_move} -->
<div class="flex items-center ml-auto gap-x-4">

  <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
  <range tabindex="0"
    bind:this={range}  
    on:click={on_click}
    on:keydown={on_click}
    on:mousedown={on_mouse_down}
    class="block relative border-2 border-black h-4 p-0.5 cursor-pointer w-8
    dark:border-white">
  
    <percentage
      bind:this={percentage}
      class="h-full bg-black flex
      dark:bg-white">
    </percentage>
  
    <slider_thumb
      bind:this={slider_thumb}
      class="block absolute rounded-full bg-white border-black border-2 h-6 w-6 top-1/2 -translate-y-1/2
      dark:bg-black dark:border-white">
    </slider_thumb>
  </range>
  <input
    class="w-12 pl-3.5"
    type="number"
    pattern="[0-9]*"
    bind:value={value}
    on:keydown={(evt) => ["e", "E", "+", "-"].includes(evt.key) && evt.preventDefault()}>
</div>