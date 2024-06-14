<script lang="ts">
  import { authToken } from "../stores/auth";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let error = "";

  async function login() {
    const res = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    if (res.ok) {
      const data = await res.json();
      authToken.set(data.token);
      error = "";
    } else {
      const data = await res.json();
      error = data.error;
    }
  }
</script>

<main>
  <h2>Login</h2>
  <form on:submit|preventDefault={login}>
    <label>
      Email:
      <input type="email" bind:value={email} />
    </label>
    <label>
      Password:
      <input type="password" bind:value={password} />
    </label>
    <button type="submit">Login</button>
  </form>
  {#if error}
    <p>{error}</p>
  {/if}
</main>
