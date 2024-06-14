<script lang="ts">
  import { authToken } from "../stores/auth";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let error = "";

  async function register() {
    const res = await fetch("/api/register", {
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
  <h2>Register</h2>
  <form on:submit|preventDefault={register}>
    <label>
      Email:
      <input type="email" bind:value={email} />
    </label>
    <label>
      Password:
      <input type="password" bind:value={password} />
    </label>
    <button type="submit">Register</button>
  </form>
  {#if error}
    <p>{error}</p>
  {/if}
</main>
