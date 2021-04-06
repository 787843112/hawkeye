import Vue from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import Vant from 'vant'
import 'vant/lib/index.css'
import 'vant/lib/icon/local.css'
Vue.config.productionTip = false
Vue.use(Vant)

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app')
