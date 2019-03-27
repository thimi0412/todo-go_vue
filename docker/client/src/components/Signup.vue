<template>
  <div class="signup">
    <h2>Sign up</h2>
    <input type="email" placeholder="Emali" v-model="email" style="border: 1px solid gray">
    <input type="password" placeholder="Password" v-model="password" style="border: 1px solid gray">
    <button @click="signUp">Register</button>
    <p>Do you have an account?
      <router-link to="/signin">sign in now!!</router-link>
    </p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Signup',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signUp: function(){
      let params = new URLSearchParams();
      params.append('email', this.email);
      params.append('password', this.password);
      axios.post('http://localhost:8090/signup', params)
      .then(response => (console.log(response)))
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
.signup {
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