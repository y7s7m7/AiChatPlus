import {createApp} from 'vue'
import {createRouter, createWebHashHistory} from "vue-router";
import App from './App.vue'
import './style.css';
import IndexView from "./Views/IndexView.vue";
import {createPinia} from "pinia";
import ChatPage from "./Pages/ChatPage.vue";
import ChatViewPage from "./Views/ChatViewPage.vue";
import HelloChat from "./Pages/HelloChat.vue";
const routes = [
    {
        path: '/',
        component: IndexView,
        children: [
            {
                path:'/chat',
                name:'chat',
                component:ChatPage,
            }
        ]
    }//每一个对象就是一个路由
]
const router = createRouter({
    history: createWebHashHistory(),
    routes,
})
const pinia = createPinia()
const app=createApp(App)
app.use(router)
app.use(pinia)
app.mount('#app')
