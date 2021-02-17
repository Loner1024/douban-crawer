import SearchResult from './views/SearchResult.vue'
import Home from './views/Home.vue'



let routes = [
    {path:"/search",component:SearchResult},
    {path:"/",component:Home},
]

export {
    routes
}