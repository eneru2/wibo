<script lang="ts">
    import { onMount } from "svelte";
    import Mind from "./Mind.svelte"

  let wants_output_code;
  let tooltip_on_screen = false;
  let output_name = "my_super_image";
  let alt = "Super alt";
  


  let formats = {
    webp: true,
  }
  let indenting_amount = 4;
  let mime_types = false;
  let relative_path = false;
  let absolute_path = true;
  let relative_path_string;
  let one_liners = false;
  
  let string = ``

  function indent(times = 1){
    for(let i = 0; i < indenting_amount*times; i++){
      string += ` `
    }
  }

  function constructString() {
    if(Object.keys(formats).length === 1){
      string += `<img`
      let format = Object.keys(formats)[0]
      if(one_liners){
        string += ` src="`
        if(relative_path) {
          if(absolute_path){
            string += `/${relative_path_string}/${output_name}.${format}" alt="${alt}">`
            if(mime_types)
            alert(string)
          } else {
            string += `${relative_path_string}/${output_name}.${format}" alt="${alt}">`
            alert(string)
          }
        } else {
          string += `${output_name}.${format}" alt="${alt}">`
            alert(string)
        }
      } else {
        string += `\n`
        indent() 
        if(relative_path) {
          if(absolute_path){
            string += `src="/${relative_path_string}/${output_name}.${format}"\n`;
            indent();
            string += `alt="${alt}">`;
            alert(string)
          } else {
            string += `src="${relative_path_string}/${output_name}.${format}"\n`;
            indent();
            string += `alt="${alt}">`;        
            alert(string)
          }
        } else {
          if(absolute_path){
            string += `src="/${output_name}.${format}"\n`;
            indent();
            string += `alt="${alt}">`
            alert(string)
          } else {
            string += `src="${output_name}.${format}"\n`;
            indent();
            string += `alt="${alt}">`
            alert(string)
          }
        }
      }
    } else if(Object.keys(formats).length === 2){
      string += `<picture>\n`;
     indent() 
      string += `<source`
      let first_format = Object.keys(formats)[0];
      let second_format = Object.keys(formats)[1];
      if (one_liners) {
        string += ` srcset="`
        if(relative_path){
          if(absolute_path){
            string += `/${relative_path_string}/${output_name}.${first_format}"`
            if(mime_types){
              string += ` type="image/${first_format}">\n`
             indent() 
              string +=`<img src="/${relative_path_string}/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img src="/${relative_path_string}/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            }
          } else {
            string += `${relative_path_string}/${output_name}.${first_format}"`
            if(mime_types){
              string += ` type="image/${first_format}">\n`
             indent() 
              string +=`<img src="${relative_path_string}/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img src="${relative_path_string}/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            }
          }
        } else {
          if(absolute_path){
            string += `/${output_name}.${first_format}"`
            if(mime_types){
              string += ` type="image/${first_format}">\n`
             indent() 
              string +=`<img src="/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img src="/${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            }
          } else {
            string += `${output_name}.${first_format}"`
            if(mime_types){
              string += ` type="image/${first_format}">\n`
             indent() 
              string +=`<img src="${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img src="${output_name}.${second_format}" alt="${alt}">\n</picture>`
              alert(string)
            }
          }
        }
      } else {
        string += `\n`;
        indent(2)
        string += `srcset="`
        if(relative_path){
          if(absolute_path){
            string += `/${relative_path_string}/${output_name}.${first_format}"`
            if(mime_types){
              string += `\ntype="image/${first_format}">\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="/${relative_path_string}/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="/${relative_path_string}/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            }
          } else {
            string += `${relative_path_string}/${output_name}.${first_format}"`
            if(mime_types){
              string += `\ntype="image/${first_format}">\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="${relative_path_string}/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="${relative_path_string}/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            }
        }
        } else {
          if(absolute_path){
            string += `/${output_name}.${first_format}"`
            if(mime_types){
              string += `\n`
              indent(2)
              string += `type="image/${first_format}">\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="/${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            }
          } else {
            string += `${output_name}.${first_format}"`
            if(mime_types){
              string += `\n`
              indent(2)
              string += `type="image/${first_format}">\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            } else {
              string += `>\n`
             indent() 
              string +=`<img\n`;
              indent(2)
              string += `src="${output_name}.${second_format}"\n`; 
              indent(2)
              string +=`alt="${alt}">\n</picture>`
              alert(string)
            }
          }
        }
      }
    }
  }

  $: default_string = `<picture>
  <source
    srcset="${output_name}.avif"
    type="image/avif">
  <source
    srcset="${output_name}.avif"
    type="image/webp">
  <img 
    src="${output_name}.jpg"
    alt="${alt}">
</picture>`;

  $: no_mime_string = `<picture>
  <source srcset="${output_name}.avif">
  <source srcset="${output_name}.avif">
  <img 
    src="${output_name}.jpg"
    alt="${alt}">
</picture>`;

//   $: relative_path_string = `<picture>
//   <source
//     srcset="${relative_path}/${output_name}.avif"
//     type="image/avif">
//   <source
//     srcset="${relative_path}/${output_name}.avif"
//     type="image/webp">
//   <img 
//     src="${relative_path}/${output_name}.jpg"
//     alt="${alt}">
// </picture>`;



  onMount(() => {
    wants_output_code = true;
    output_name = "my_super_image"
  })

  
// export let string = default_string
</script>
<output_code class="w-fit relative">
  <Mind></Mind>
  <!-- Do you want to output the corresponding code for the images? -->
  <h2 on:click={constructString} class="text-3xl">Optional options</h2>
  <div class="flex justify-between">
    <h2>Output code for images</h2>
    <!-- svelte-ignore a11y-mouse-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <p
      class="w-6 h-6 text-blue-imperial"
      on:mouseover={() => (tooltip_on_screen = true)}
      on:mouseout={() => (tooltip_on_screen = false)}>
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
    </p>
    <input
      bind:this={wants_output_code}
      on:change={() => wants_output_code = !wants_output_code}
      type="checkbox" />
  </div>
  
  {#if wants_output_code}
    <div class="pl-8">
      <div class="flex justify-between">
        <p class="text-blue-imperial font-bold">
          Image Alt
        </p> 
      </div>
      <div>
        <input type="text">
        <h2>Code format options:</h2>
        <ul>
          <li>indenting_amount</li>
          <li>Include MIME types</li>
          <li>Relative path</li>
          <li>One liners</li>
        </ul>
      </div>
    </div>
    {:else}
    "fuyo"
    {/if}

    {#if tooltip_on_screen}
    <div
      class="absolute right-[-24rem] top-0 w-96 px-6 py-3 bg-blue-imperial text-background">
      <h3>EXAMPLE</h3>
      <pre 
        class="">{default_string}</pre>
    </div>
    {/if}
  </output_code>

<style>
  h3 {
    text-decoration: underline;
    font-weight: bold;
  }
</style>