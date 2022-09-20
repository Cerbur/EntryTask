import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from "@/components/Login";
import Home from "@/components/Home";

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/home',
        name: 'home',
        component: Home
    },
    {
        path: '/',
        name: 'defalut',
        component: Login
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
