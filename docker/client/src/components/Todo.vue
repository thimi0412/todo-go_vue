<template>
  <div class='hello'>
    <h1>TODO</h1>
    <article class="todo-list">
      <div v-for="todo in todos" :key="todo.id">
        <ul>{{ todo.context }}, {{ todo.limit_date }}, {{ todo.updated_at }}</ul>
      </div>
    </article>

    <input v-model="form.context" type="text" placeholder="context" />
    <input @input="updateValue" type="datetime-local" placeholder="date-time" />
    <button @click="postTodo()" style="background-color: gray;">post</button>

    <p>{{ form }}</p>
  </div>
</template>

<script>
import axios from 'axios'

const tmp = document.cookie.split('; ')
let data = tmp[0].split('=')[1]
axios.defaults.headers.common['Authorization'] = data

export default {
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      form: {
        context: '',
        dateTime: ''
      },
      todos: [],
      date: new Date().toLocaleString().replace(/\//g, '-')
    }
  },
  created () {
    this.get()
  },
  methods: {
    get () {
      axios.get('http://localhost:8090/todo')
        .then(res => {
          console.log(res)
          let result = []
          for(let r of res.data.result){
            console.log(r)
            result.push(
              {
                context: r.context,
                limit_date: formatDate(r.limit_date),
                updated_at: formatDate(r.updated_at),
              }
            )
          }
          this.todos = result
        })
        .catch(err => { console.log(err.response) })
    },
    updateValue (event) {
      this.form.dateTime = event.target.value.replace(/T/, ' ').replace(/\//, '-') + ':00'
    },
    postTodo () {
      let params = new URLSearchParams();
      params.append('context', this.form.context);
      params.append('limit_date', this.form.dateTime);
      axios.post('http://localhost:8090/todo', params)
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
function formatDate(date) {
  let result = date.replace(/T/, ' ')
  result = result.split('+')[0]
  return result
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