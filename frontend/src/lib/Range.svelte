<script lang="ts">
  let range:undefined | HTMLInputElement;
  let percentage;
  let slider_thumb;
  let value;  
  let is_pressed = false;  

  function on_click(event){
    const container_width = 128;  
    const range_max = 40;
    const range_min = 0;
    const thumb_width = 16;
    
    const range_container_size = range.getBoundingClientRect();  
    const x = event.clientX - range_container_size.left; 

    let track_width = Math.round(x*100/container_width);
    let thumb_pos: number;
    value = Math.round(x*range_max/container_width);

    if (x > container_width){
      thumb_pos = Math.round(container_width*(100-(thumb_width+4))/container_width);
    } else {
      thumb_pos = Math.round(x*(100-(thumb_width+4))/container_width);
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
    } if (value > range_max) {
      value = range_max
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

</script>

<!-- on:mousemove={on_mouse_move} -->
<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
<range tabindex="0"
  bind:this={range}  
  on:click={on_click}
  on:keydown={on_click}
  on:mousedown={on_mouse_down}
  class="block relative border-2 border-black  w-32 h-4 ml-16 p-0.5 cursor-pointer
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
<value>{value}</value>