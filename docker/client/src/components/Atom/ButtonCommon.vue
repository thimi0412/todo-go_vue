<template>
  <button type="button" @click="$emit('emitAction')">{{ buttonName }}</button>
</template>

<script>
export default {
  props: {
    buttonName: String
  },
  methods: {
    emitAction () {
      this.$emit('emitAction')
    }
  }
}
</script>

<style>

</style>
