<template>
    <div class="signin">
        <h2>Sign in</h2>
        <input type="email" placeholder="Email" v-model="email">
        <input type="password" placeholder="Password" v-model="password">
        <button @click="signin">Signin</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Signin',
  data: function () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signin: function(){
      let params = new URLSearchParams();
      params.append('email', this.email);
      params.append('password', this.password);
      axios.post('http://localhost:8090/signin', params)
      .then(response => {
        console.log(response)
        document.cookie = 'jwt=' + response.data.token
        this.$router.push({ path: '/todo' })
      })
      .catch(response => (console.log(response)))
    }
  }

}
</script>

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
.signin {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>