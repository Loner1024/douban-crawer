import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'

let app = createApp(App)

let router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

app.use(router)
app.mount('#app')
