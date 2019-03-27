<template>
  <div class='hello'>
    <h1>TODO</h1>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  created () {
    this.getSessions()
  },
  methods: {
    getSessions () {
      axios.get('http://localhost:8090/todo')
        .then(res => { console.log(res) })
        .catch(err => { console.log(err.response) })
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>