<template>
  <div class='hello'>
    <h1>TODO</h1>
    <article class="todo-list">
      <table align="center">
        <tr>
          <th>TODO</th>
          <th>LIMIT</th>
          <th>UPDATED_AT</th>
          <th>STUTAS</th>
        </tr>
        <tr v-for="todo in todos" :key="todo.id">
          <td>{{ todo.context }}</td>
          <td>{{ todo.limit_date }}</td>
          <td>{{ todo.updated_at }}</td>
          <td><button>delete</button></td>
        </tr>
      </table>
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
          this.get()
        })
        .catch(err => {
          console.log(err)
        })
    },
    // 作成途中
    deleteTodo () {
      axios.delete('http://localhost:8090/todo')
        .then(res => {
          console.log(res)
          this.get()
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
/**
 * @param {String} date datetime形式の文字列
 * @returns {String} yyyy/mm/dd hh:mm形式
 */
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
table{
  border-collapse: collapse;
  border-spacing: 0;
  width: 100%;
}

table tr{
  border-bottom: solid 1px #eee;
  cursor: pointer;
}

table tr:hover{
  background-color: #d4f0fd;
}

table th,table td{
  text-align: center;
  width: 25%;
  padding: 15px 0;
}
</style>