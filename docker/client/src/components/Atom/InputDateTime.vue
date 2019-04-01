<template>
  <input @input="emitValue" type="datetime-local" :placeholder="placeholder" />
</template>

<script>
export default {
  props: {
    placeholder: String
  },
  methods: {
    emitValue (e) {
      e.target.value.replace(/T/, ' ').replace(/\//, '-') + ':00'
    }
  }
}
</script>

<style>

</style>
