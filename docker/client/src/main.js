import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VModal from 'vue-js-modal'

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  render: h => h(App)
});
Vue.use(VModal)