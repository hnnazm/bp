<script>
  import { env } from "$env/dynamic/public";

  const API_ENDPOINT = env.PUBLIC_API_ENDPOINT || "http://localhost:1323";

  /** @type {import('./$types').PageData} */
  export let data;

  async function incrementCounter() {
    await fetch(`${API_ENDPOINT}/counter`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        value: 1,
      }),
    });
  }
</script>

<header class="mb-24 p-8 flex flex-row justify-between items-center">
  <h1 class="text-5xl text-slate-50 font-bold">
    {env.PUBLIC_SITE_NAME || "My Site"}
  </h1>
</header>

<main
  class="md:max-w-screen-lg md:mx-auto px-4 md:px-0 flex flex-col justify-center items-center space-y-4"
>
  <div
    class="w-full py-3 flex flex-row justify-between items-center border border-slate-300 rounded-lg"
  >
    <p id="counter-value" class="text-3xl text-slate-50 text-center grow">
      {data.counter}
    </p>

    <button on:click={incrementCounter} class="p-2 group">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 -960 960 960"
        class="w-8 h-8 fill-slate-50 group-disabled:opacity-50"
      >
        <path
          d="M480-529 324-373q-11 11-28 11t-28-11q-11-11-11-28t11-28l184-184q6-6 13-8.5t15-2.5q8 0 15 2.5t13 8.5l184 184q11 11 11 28t-11 28q-11 11-28 11t-28-11L480-529Z"
        />
      </svg>
    </button>
  </div>
</main>
