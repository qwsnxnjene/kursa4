import Vue from 'vue'
import App from './App.vue'
import router from './router' // Импортируем роутер
import axios from 'axios'

Vue.config.productionTip = false
axios.defaults.baseURL = 'http://localhost:8081'
Vue.prototype.$http = axios

new Vue({
  router, // Добавляем роутер в экземпляр Vue
  render: h => h(App),
}).$mount('#app')