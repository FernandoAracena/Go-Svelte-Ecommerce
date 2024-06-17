<script lang="ts">
  import { onMount } from "svelte";
  import ShoppingCart from "./components/ShoppingCart.svelte";
  import ProductCard from "./components/ProductCard.svelte";

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

  function checkout(cart: CartItem[], totalSum: number) {
    cart.totalprice = totalSum;
    checkoutClick = true;
    console.log("Cart:", cart);
  }
</script>

<main>
  <h1 class="text-2x1 font-bold mb-8 text-center">Coffee List</h1>
  <ShoppingCart
    {cart}
    {handleAddToCart}
    {handleRemoveFromCart}
    {checkout}
    {checkoutClick}
  />
  <div class="card gap-6 p-8 lg:w-1/2">
    {#each coffees as coffee}
      <ProductCard
        {coffee}
        {selectedCoffee}
        {editingCoffeeId}
        {handleSaveEditCoffee}
        {handleCancelEdit}
        {handleAddToCart}
      />
    {/each}
  </div>
</main>
