<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<tooltip
  class="absolute right-0 top-0 max-w-96 px-6 py-3 text-black bg-background border border-black z-10">
  <h3>DESCRIPTION</h3>
  <p>Allows to create custom file names for the output files.</p>
  <h3>DEFAULT</h3>
  <p>If this field is empty the output will default to the input's file name.</p>
  <h3>OPTIONS</h3>
  <div>
    [res]<br />
      If you want to use the target res in the name of your output, you can do
      it with this argument, as an example,<br /><b>flying_bird-[res]x</b> would
      output:<br/>
      <p class="pl-2">
        - flying_bird-512x.jpg<br />
        - flying_bird-512x.webp<br />
        - flying_bird-512x.avif
      </p>
  </div>
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
