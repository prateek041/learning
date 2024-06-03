<script setup>
import Balance from "./components/Balance.vue";
import History from "./components/History.vue";
</script>

<template>
  <div class="flex justify-center h-screen w-1/2 mx-auto">
    <div class="flex flex-col items-start bg-red-200 mt-8 w-1/2 h-full">
      <Balance />
      <History />
    </div>
  </div>
</template>
