import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';
import 'view-design/dist/styles/iview.css';
import VueRouter from "vue-router";
import router from './router'
import axios from 'axios'


Vue.use(ElementUI);
Vue.use(VueRouter);
//配置请求的基本路径
// axios.defaults.baseURL = 'http://10.128.164.14:8410'
axios.defaults.timeout = 3000
axios.defaults.baseURL = 'http://120.79.214.70:8410'
Vue.prototype.$http = axios


//关闭生产提示
Vue.config.productionTip = false



new Vue({
  render: h => h(App),
  router:router
}).$mount('#app')
