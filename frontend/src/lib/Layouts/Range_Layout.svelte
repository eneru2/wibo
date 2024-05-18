<script lang="ts">
  import { onMount } from "svelte";    
  
  export let value:number;
  export let container_width:number;  
  export let max = 100;
  export let maxlength = 3;
  export let value_change;
  export let name:string;
  let input_element;

  let range:undefined | HTMLInputElement;
  let percentage;
  let slider_thumb;
  let is_pressed = false;  
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
    is_pressed = true
    on_click(event)

    document.addEventListener("mousemove", on_mouse_move)

    document.addEventListener("mouseup", () => { 
      is_pressed = false, 
      document.removeEventListener("mousemove",on_mouse_move)
    })
  }

  function on_mouse_move(event){
    if(is_pressed) {
      on_click(event)
    }
  }

  export function change_input(){
    setTimeout(() => {
      if (value > max) {
        value = max;
      };
      //value/max*100 - myvar + "%"
      // let myvar = thumb_width / container_width * 120
      // value es un valor porcentual y
      // x es una posicion relativa
      slider_thumb.style.left = Math.round(value*(100-(thumb_width))/max) + "%";
      percentage.style.width = Math.round(value/max*100) + "%";      
    }, 100);
  }

  // export function on_mount(avif, webp){
  //   if (name === "avif"){
  //     value = avif
  //     max = 63
  //     range.style.width = container_width + "px";
  //     slider_thumb.style.left = Math.round(value*(100-(thumb_width))/max) + "%";
  //     percentage.style.width = Math.round(value/max*100) + "%"; 
  //   } else {
  //     value = webp
  //     max = 100
  //     range.style.width = container_width + "px";
  //     slider_thumb.style.left = Math.round(value*(100-(thumb_width))/max) + "%";
  //     percentage.style.width = Math.round(value/max*100) + "%"; 
  //   }
  // }

  onMount(() => {
    // slider_thumb.style.left = value + "%"
    // alert(Math.round(value/max*100) + "%")
    // percentage.style.width = alert(Math.round(max) + "%");
      range.style.width = container_width + "px";
      slider_thumb.style.left = Math.round(value*(100-(thumb_width))/max) + "%";
      percentage.style.width = Math.round(value/max*100) + "%"; 
  })

</script>

<div class="flex items-center ml-auto gap-x-4">

  <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <range tabindex="0"
    bind:this={range}  
    on:click={on_click}
    on:mousedown={on_mouse_down}    
    on:mouseup={value_change(name, value)}
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
  <!-- bind:value={value}     -->
  <input
    class="w-12 pl-3.5"
    type="number"
    maxlength={maxlength}
    oninput="this.value=this.value.slice(0,this.maxLength);"
    pattern="[0-9]*"
    on:change={() => value_change(name, value)}
    bind:value={value}
    on:keydown={(evt) => {
      ["e", "E", "+", "-"].includes(evt.key) && evt.preventDefault();
      change_input()
      }}
    >
</div>