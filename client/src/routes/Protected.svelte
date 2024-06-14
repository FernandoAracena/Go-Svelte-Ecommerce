<script lang="ts">
  import { authToken } from "../stores/auth";
  import { onMount } from "svelte";
  import { goto } from "@sapper/app";

  $: if (!authToken) {
    goto("/login");
  }

  function logout() {
    authToken.set("");
  }
</script>

<main>
  <h1>Protected Routes</h1>
  <button on:click={logout}>Logout</button>
</main>
