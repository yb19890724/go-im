import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            redirect: '/login'
        },
        {
            path: '/login',
            name: 'Login',
            components: require("../components/auth/login")
        },
        {
            path:'/home',
            name:'home',
            beforeEnter:requireAuth,
            components: require("../components/home")
        }
    ],
    mode:"history",

})

function requireAuth (to, from, next) {
    return next();
   /* if (localStorage.getItem("currentUserToken")) {
        return next();
    }else{
        return next('/')
    }*/
}

