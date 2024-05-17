<script>
  import OutputCode from "./Output_Code_Options.svelte";
  export let output_name = "my_image"
  export let alt = "My super alt that's very descriptive"

  export let string = `<picture>
  <source
    srcset="${output_name}.avif"
    type="image/avif">
  <source
    srcset="${output_name}.avif"
    type="image/webp">
  <img 
    src="${output_name}.jpg"
    alt="${alt}">
</picture>`
</script>
{string}