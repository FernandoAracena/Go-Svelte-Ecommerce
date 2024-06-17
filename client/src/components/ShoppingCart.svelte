<script lang="ts">
  export let cart = [];
  export let handleAddToCart;
  export let handleRemoveFromCart;
  export let checkout;
  export let checkoutClick;

  $: totalSum = getTotalSum(cart);

  function getTotalSum(item: CartItem[]) {
    return cart.reduce((total, item) => total + item.price * item.quantity, 0);
  }
</script>

{#if cart.length > 0}
  <div class="p-2 rounded top-0 right-0 fixed z-10 bg-indigo-50">
    {#if checkoutClick == false}
      {#each cart as item}
        <div class="flex items-center justify-between w-64">
          <div>
            <p class="text-blue-500">{item.quantity}</p>
          </div>
          <div>
            <p class="p-2">{item.name}</p>
          </div>
          <div>
            <p class="p-2">{item.price} nok</p>
          </div>
          <div class="text-3xl">
            <button
              on:click={() => handleAddToCart(item)}
              class="hover:text-4xl text-green-500">+</button
            >
            <button
              on:click={() => handleRemoveFromCart(item)}
              class="hover:text-4xl text-red-500">-</button
            >
          </div>
        </div>
      {/each}
      <div class="flex items-center justify-between w-64 my-4">
        <h3 class="text-lg font-bold">Total: {totalSum} nok</h3>
        <button
          on:click={() => checkout(cart, totalSum)}
          class="bg-yellow-500 text-white p-2 rounded">Checkout</button
        >
      </div>
    {:else}
      <div class="row items-center justify-between w-64 my-4">
        <h2>Thanks for your purchase!</h2>
        <p>Your Coffee/s will be ready on 15 minutes</p>
        {#each cart as item}
          <div class="flex items-center justify-between w-full">
            <p>{item.quantity} - {item.name}</p>
          </div>
        {/each}
        <p class="font-bold">Total Price: {totalSum} nok</p>
      </div>
    {/if}
  </div>
{/if}
