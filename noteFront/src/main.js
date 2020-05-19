import Vue from 'vue'

import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App'
import './assets/global.css'
import axios from 'axios'

axios.defaults.withCredentials=true
Vue.prototype.$axios = axios
Vue.use(ElementUI)
Vue.config.productionTip = false


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
