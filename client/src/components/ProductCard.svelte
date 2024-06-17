<script lang="ts">
  export let coffee;
  export let selectedCoffee;
  export let editingCoffeeId;
  export let handleSaveEditCoffee;
  export let handleCancelEdit;
  export let handleAddToCart;
</script>

<div class="border rounded-lg p-4 shadow-lg bg-white">
  {#if editingCoffeeId === coffee.id}
    <diV>
      <label class="block mb-2">
        Name:
        <input
          type="text"
          bind:value={selectedCoffee.name}
          class="block w-full p-2 border rounded-mt-1"
        />
      </label>
      <label class="block mb-2">
        Picture URL:
        <input
          type="text"
          bind:value={selectedCoffee.picture_url}
          class="block w-full p-2 border rounded-mt-1"
        />
      </label>
      <label class="block mb-2">
        Description:
        <input
          type="text"
          bind:value={selectedCoffee.description}
          class="block w-full p-2 border rounded-mt-1"
        />
      </label>
      <label class="block mb-2">
        Price:
        <input
          type="number"
          bind:value={selectedCoffee.price}
          class="block w-full p-2 border rounded-mt-1"
        />
      </label>
      <button
        on:click={handleSaveEditCoffee}
        class="bg-green-500 text-white p-1 rounded mt-4">Save</button
      >
      <button
        on:click={handleCancelEdit}
        class="bg-red-500 text-white p-1 rounded mt-4 ml-2">Cancel</button
      >
    </diV>
  {:else}
    <div>
      <h2 class="text-xl font-semibold">{coffee.name}</h2>
      {#if coffee.picture_url}
        <img
          src={coffee.picture_url}
          alt={coffee.name}
          class="w-full h-48 object-cover mt-4"
        />
      {/if}
      <p class="text-gray-700">{coffee.description}</p>
      <p class="text-gray-900">{coffee.price} nok</p>
      <div class="mt-4">
        <!-- <button
          on:click={() => handleEditClick(coffee)}
          class="bg-blue-500 text-white p-1 rounded mr-2">Edit</button
        >
        <button
          on:click={() => handleDeleteCoffee(coffee.id)}
          class="bg-red-500 text-white p-1 rounded mr-2">Delete</button
        > -->
        <button
          on:click={() => handleAddToCart(coffee)}
          class="bg-yellow-500 text-white p-1 rounded">Add to Cart</button
        >
      </div>
    </div>
  {/if}
</div>
