<script>
  import "../app.css";
</script>

<svelte:head>
  <title>BigPay</title>
  <meta name="description" content="BigPay Assessment" />
</svelte:head>

<div
  class="min-h-screen bg-[radial-gradient(ellipse_at_right,_var(--tw-gradient-stops))] from-sky-400 to-indigo-900"
>
  <slot />
</div>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
