<script lang="ts">
  export let formats:Array<string> = [];

  let avif = false;
  let webp = false;
  let jpg = false;

  function check_current_formats(format:string){
    switch(format){
      case "avif":
        avif ? add_format_to_array(format) : delete_array_element(format);
        break;
      case "jpg":
        jpg ? add_format_to_array(format) : delete_array_element(format);
        break;
      case "webp":
        webp ? add_format_to_array(format) : delete_array_element(format);
        break;
    };
  };

  function delete_array_element(format:string){
    for (let i = 0; i <= formats.length - 1; i++) {
      if(format === formats[i]) {
        formats.splice(i, 1);
      }    
    };
  };
    
  function add_format_to_array(format:string){
    const findElement = formats.filter((el) => el.includes(format));
    if(findElement.toString() === ""){
      formats.push(format);
    };
    sort_array();
  }; 

  function sort_array(){
    switch(formats.length){
      case 1:        
        return;
      default:
        formats.sort();
        for (let i = 0; i < formats.length - 1; i++) {
          if(formats[i] === "jpg"){
            formats.splice(i, 1);
            formats.push("jpg");
          }
        }
        break;
      }
  };
</script>

<h2>Output formats:</h2>
<div class="container flex items-center gap-x-4">
  <label
    class=" flex items-center gap-x-2">
    <input
      type="checkbox"        
      bind:checked={avif}
      on:change={() => check_current_formats("avif")}/>
      AVIF
  </label>
  <label
    class="flex items-center gap-x-2">
  <input
    type="checkbox"
    bind:checked={webp}
    on:change={() => check_current_formats("webp")}/>
    WEBP
  </label>
  <label
    class="flex items-center gap-x-2">
    <input
      type="checkbox"
      bind:checked={jpg}
      on:change={() => check_current_formats("jpg")}/>
      JPG
  </label>
</div>