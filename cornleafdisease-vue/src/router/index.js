import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import Login from "@/views/Login.vue";
import SignUp from '@/views/SignUp.vue'
import Outcome from "@/views/Outcome.vue";
// const originalPush = VueRouter.prototype.push;
// VueRouter.prototype.push = function push(location) {
//     return originalPush.call(this, location).catch(err => err);
// }
// Vue.use(VueRouter)
const routes = [
    {
        path: '/',
        name:"Login",
        component: Login
    },
    {
        path: "/home",
        name: "Home",
        component: Home,
    },
    {
        path: '/signup',
        name:"SignUp",
        component: SignUp
    },
    {
        path: '/outcomt',
        name:"Outcome",
        component: Outcome
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;