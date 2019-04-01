<template>
  <input v-model="bindValue" type="text" :placeholder="placeholder" />
</template>

<script>
export default {
  prop: {
    bindValue: String,
    placeholder: String
  }
}
</script>

<style>

</style>

