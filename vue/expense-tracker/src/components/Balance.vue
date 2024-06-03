<script setup>
const Balance = 1000;
const number = 100;
</script>

<template>
  <div class="w-full">
    <h1 class="text-2xl mb-5">Expense Tracker</h1>
    <div>
      <h3>Your Balance</h3>
      <h1 class="text-3xl">${{ Balance }}</h1>
    </div>
    <div class="flex divide-x gap-2 w-full h-full">
      <div
        class="w-1/2 h-24 bg-sky-50 flex flex-col justify-center items-center"
      >
        <h2>INCOME</h2>
        <h1>${{ number }}</h1>
      </div>
      <div
        class="w-1/2 h-24 bg-sky-50 flex flex-col justify-center items-center"
      >
        <h2>EXPENSE</h2>
        <h1>${{ number }}</h1>
      </div>
    </div>
  </div>
</template>
