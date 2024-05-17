<script>
  export let description = ""
  export let default_behaviour = ""
  export let options = ""
</script>

<tooltip class="absolute right-0 top-0 max-w-96 w-96 px-6 py-3 text-black bg-background border border-black z-10">
  <h3>DESCRIPTION</h3>
  <p>{description}</p>
  {#if default_behaviour !== ""}    
    <h3>DEFAULT</h3>
    <p>{default_behaviour}</p>
  {/if}
  {#if options !== ""}
    <h3>OPTIONS</h3>
    <p>{options}</p>
  {/if}
</tooltip>

<style lang="postcss">
  h3 {
    text-decoration: underline;
    font-weight: bold;
  }
  p {
    @apply mb-1;
  }
</style>
