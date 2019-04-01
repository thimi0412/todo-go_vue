<template>
  <div>
    <InputText></InputText>
    <InputDateTime></InputDateTime>
    <ButtonCommon></ButtonCommon>
  </div>
</template>

<script>
import ButtonCommon from '../Atom/ButtonCommon'
import InputDateTime from '../Atom/InputDateTime'
import InputText from '../Atom/InputText'

export default {
  components: [
    ButtonCommon,
    InputDateTime,
    InputText
  ]
}
</script>

<style>

</style>
