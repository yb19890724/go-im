import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
        path: '/login',
        name: 'Login',
        components: require("../components/auth/login")
    }
  ],
  mode:"history"
})
