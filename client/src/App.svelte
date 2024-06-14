<script lang="ts">
  import { onMount } from "svelte";

  interface Coffee {
    id: number;
    name: string;
    description: string;
    price: number;
    picture_url: string;
  }

  interface CartItem extends Coffee {
    quantity: number;
    totalprice: number;
  }

  let coffees: Coffee[] = [];
  let cart: CartItem[] = [];
  let checkoutClick: boolean = false;
  let editingCoffeeId: number | null = null;
  let selectedCoffee: Coffee = {
    id: 0,
    name: "",
    description: "",
    price: 0,
    picture_url: "",
  };

  onMount(async () => {
    await fetchCoffees();
  });
  async function fetchCoffees() {
    const res = await fetch("/api/coffees");
    const data: Coffee[] = await res.json();
    coffees = data;
  }

  async function handleDeleteCoffee(id: number) {
    const res = await fetch(`/api/coffees/${id}`, {
      method: "DELETE",
    });
    if (res.ok) {
      coffees = coffees.filter((coffee) => coffee.id !== id);
    } else {
      console.error("Failed to Delete Coffee");
    }
  }

  function handleEditClick(coffee: Coffee) {
    editingCoffeeId = coffee.id;
    selectedCoffee = { ...coffee };
  }

  async function handleSaveEditCoffee() {
    const res = await fetch(`/api/coffees/${selectedCoffee.id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(selectedCoffee),
    });
    if (res.ok) {
      await fetchCoffees();
      editingCoffeeId = null;
    } else {
      console.error("Failed to Update Coffee");
    }
  }

  function handleCancelEdit() {
    editingCoffeeId = null;
    selectedCoffee = {
      id: 0,
      name: "",
      description: "",
      price: 0,
      picture_url: "",
    };
  }

  function handleAddToCart(coffee: Coffee) {
    const index = cart.findIndex((item) => item.id === coffee.id);
    if (index !== -1) {
      cart[index].quantity += 1;
    } else {
      cart = [...cart, { ...coffee, quantity: 1 }];
    }
  }

  function handleRemoveFromCart(coffee: Coffee) {
    const index = cart.findIndex((item) => item.id === coffee.id);
    if (index !== -1) {
      cart[index].quantity -= 1;
      if (cart[index].quantity < 1) {
        cart = [...cart.slice(0, index), ...cart.slice(index + 1)];
      }
    }
  }

  $: totalSum = getTotalSum(cart);

  function getTotalSum(item: CartItem[]) {
    return cart.reduce((total, item) => total + item.price * item.quantity, 0);
  }

  function checkout(cart: CartItem[], totalSum: number) {
    cart.totalprice = totalSum;
    checkoutClick = true;
    console.log("Cart:", cart);
  }
</script>

<main>
  <h1 class="text-2x1 font-bold mb-8 text-center">Coffee List</h1>
  {#if cart.length > 0}
    <div class="p-2 rounded top-0 right-0 fixed bg-indigo-50">
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
  <div class="grid grid-cols-1 sm:grid-cols-1 lg:grid-cols-2 gap-6 p-8">
    {#each coffees as coffee}
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
    {/each}
  </div>
</main>
