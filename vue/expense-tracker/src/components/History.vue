<script setup>
import Transaction from "./Transaction.vue";
import { ref } from "vue";

const items = [
  { item: "hello", price: "world" },
  { item: "goodbye", price: "world" },
];
</script>

<template>
  <div>
    <div>History</div>
    <li v-for="item in items">
      <Transaction :item="item.item" :price="item.price" />
    </li>
  </div>
</template>
