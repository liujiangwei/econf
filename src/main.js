import Vue from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue'


import "ant-design-vue/dist/antd.css";
import VueRouter from 'vue-router';

import HelloWorld from "./components/HelloWorld"
import Etcd from "./components/Etcd"

Vue.use(Antd)
Vue.use(VueRouter)

Vue.config.productionTip = false

const routes = [
  { path: '/', component: Etcd },
  { path: '/config-center/', component: HelloWorld },
  { path: '/config-project', component: HelloWorld }
]
const router = new VueRouter({
  routes // (缩写) 相当于 routes: routes
})


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
