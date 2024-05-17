<script lang="ts">
  import { onMount } from "svelte";

  let wants_output_code = true;
  let tooltip_on_screen = false;
  export let output_schema = "";
  export let internal_output_schema:string;
  let alt = "";

  export let formats:Array<string> = [];

  let indenting_amount = 4;
  let mime_types = false;  
  let absolute_path = false;  
  let one_liners = false;
  
  let custom_path = true; 
  let custom_path_route = "";

  function indent(times = 1){
    let inner_spaces = "";
    for(let i = 0; i < indenting_amount*times; i++){
      inner_spaces += ` `
    }
    return inner_spaces;
  };

  function delete_one_liners(str, indent_amount = 2){
    const srcset_regex = /e s/g
    const srcset_replace = 'e\n' + indent(indent_amount) + 's'
    
    const mime_type_regex = /" t/g
    const mime_type_replace = '"\n' + indent(indent_amount) + 't'

    const src_regex = /g s/g
    const src_replace = 'g\n' + indent(indent_amount) + 's'

    const alt_regex = /" a/g
    const alt_replace = '"\n' + indent(indent_amount) + 'a'

    const result =
    str
      .replace(srcset_regex, srcset_replace)
      .replace(mime_type_regex, mime_type_replace)
      .replace(src_regex, src_replace)
      .replace(alt_regex, alt_replace)
    return result
  };

  function delete_mimes(str){
    const mime_regex = / type.*2?"/g
    const result = str.replace(mime_regex, "")
    return result;
  }

  function delete_absolute_paths(str){

    const srcset_regex = /srcset="\//g
    const srcset_replace = 'srcset="'

    const src_regex = /src="\//g
    const src_replace = 'src="'
    const result = 
    str
      .replace(src_regex, src_replace)
      .replace(srcset_regex, srcset_replace)
    return result;
  }

  function append_custom_path(str){
    if(absolute_path){
      const src_regex = /src="\//g
      const srcset_regex = /set="\//g
      const src_replace = `src="/${custom_path_route}/`
      const srcset_replace = `set="/${custom_path_route}/`

      const res = 
      str
        .replaceAll(srcset_regex, srcset_replace)
        .replaceAll(src_regex, src_replace)
      return res;
    } else {
      const src_regex = /src="/g
      const srcset_regex = /set="/g
      const src_replace = `src="${custom_path_route}/`
      const srcset_replace = `set="${custom_path_route}/`

      const res = 
      str
        .replaceAll(srcset_regex, srcset_replace)
        .replaceAll(src_regex, src_replace)
      return res;
    }
  }

  function initial_indenting(str){
    const string_start = /^.+</
    const end_of_tag_regex = />\n  </g
    const end_of_tag_replace = '>\n' + indent() + '<'

    const source_regex = /\n.*?<sou/g
    const source_replace = '\n' + indent() + `<sou`
    
    const img_regex = /\n.*?<img/g
    const img_replace = '\n' + indent() + `<img`
    const res =
    str
      .replace(end_of_tag_regex, end_of_tag_replace)
      .replace(string_start, "<")
      .replace(source_regex, source_replace)
      .replace(img_regex, img_replace)
    return res;
  }

  export const construct_output_string = (internal_output_schema?:string) => {
    switch(formats.length){
      case 0: {
        return default_string;
      };
      case 1: {
        let format = formats[0];
        let string = `<img src="/${internal_output_schema}.${format}" alt="${alt}">`;

        string = !one_liners ? string : delete_one_liners(string, 1);
        string = mime_types ? string : delete_mimes(string);
        string = absolute_path ? string : delete_absolute_paths(string);
        string = custom_path_route != "" ? append_custom_path(string) : string;
        return string;
      };
      case 2: {
        let first_format = formats[0];
        let second_format = formats[1];
        let string = `<picture>
          <source srcset="/${internal_output_schema}.${first_format}" type="image/${first_format}">
          <img src="/${internal_output_schema}.${second_format}" alt="${alt}">\n</picture>`;          
        string = initial_indenting(string);
        string = mime_types ? string : delete_mimes(string);
        string = !one_liners ? string : delete_one_liners(string);
        string = absolute_path ? string : delete_absolute_paths(string);
        string = custom_path_route != "" ? append_custom_path(string) : string;
        return string;
      };
      case 3: {
        let first_format = formats[0];
        let second_format = formats[1];
        let third_format = formats[2];
        let string = `<picture>
          <source srcset="/${internal_output_schema}.${first_format}" type="image/${first_format}">
          <source srcset="/${internal_output_schema}.${second_format}" type="image/${second_format}">
          <img src="/${internal_output_schema}.${third_format}" alt="${alt}">\n</picture>`;

        string = initial_indenting(string);
        string = mime_types ? string : delete_mimes(string);
        string = !one_liners ? string : delete_one_liners(string);
        string = absolute_path ? string : delete_absolute_paths(string);
        string = custom_path_route != "" ? append_custom_path(string) : string;
        return string;
      };
    };
  }

  export let default_string = `<picture>
  <source
    srcset="/my_super_image.avif"
    type="image/avif">
  <source
    srcset="/my_super_image.webp"
    type="image/webp">
  <img
    src="/my_super_image.jpg"
    alt="My super descriptive alt">\n</picture>`;

  onMount(() => {    
  })

  
// export let string = default_string
</script>
<output_code class="w-full relative flex flex-col gap-y-4">
  <!-- Do you want to output the corresponding code for the images? -->
  <div class="flex items-center">
    <h2 class="flex items-center mr-4">Output code for images
      <!-- svelte-ignore a11y-mouse-events-have-key-events -->
      <!-- svelte-ignore a11y-no-static-element-interactions -->
      <div
      class="w-5 h-5 ml-1"
      on:mouseover={() => (tooltip_on_screen = true)}
      on:mouseleave={() => (tooltip_on_screen = false)}>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
        {#if tooltip_on_screen}
          <div            
            class="absolute right-0 top-0 w-96 px-6 py-3 text-black border border-black bg-background z-10">
            <h3>DESCRIPTION</h3>
            <p class="mb-1">With this option you will get code that you can
              copy and paste in your editor. You can customize the output with the
              options below.
            </p>
            <h3>EXAMPLE</h3>
            <pre
              class="">{default_string}</pre>
          </div>
        {/if}
        </div>:
    </h2>
    <!-- <input
      bind:checked={wants_output_code}      
      type="checkbox"/> -->
  </div>

  <label class="flex gap-x-2 w-full text-nowrap">
    Image alt:
    <input
      spellcheck="false"
      bind:value={alt}
      type="text">
  </label>
  <label class="flex gap-x-2 w-full text-nowrap">
    Amount of spaces on indenting:
  <input
    type="number"
    min="1"
    max="10"
    pattern="[0-9]*"
    on:keydown={(evt) => ["e", "E", "+", "-"].includes(evt.key) && evt.preventDefault()}>    
  </label>
  <label class="flex gap-x-2 w-full text-nowrap">
    Include MIME types:
    <input
      bind:checked={mime_types}
      type="checkbox">
  </label>
  <label class="flex gap-x-2 w-full text-nowrap">
    Image src:
    <input
      spellcheck="false"
      bind:value={custom_path_route}
      type="text">
  </label>
  <label class="flex gap-x-2 w-full text-nowrap">
    Absolute paths:
    <input
      bind:checked={absolute_path}
      type="checkbox">
  </label>
  <label class="flex gap-x-2 w-full text-nowrap">
    New lines for properties:
    <input
      bind:checked={one_liners}
      type="checkbox">
  </label>
</output_code>

<style>
  h3 {
    text-decoration: underline;
    font-weight: bold;
  }
</style>