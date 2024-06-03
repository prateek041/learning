<script setup lang="ts">
// here we will learn about importing and exporting props into component

const props = defineProps({
  ['item', 'price']
});

console.log(prope)
</script>

<template>
  <div>
    <p>{{ item }}</p>
    <p>{{ price }}</p>
  </div>
</template>
