import SearchResult from './views/SearchResult.vue'
import Home from './views/Home.vue'



let routes = [
    {path:"/search",component:SearchResult,props: (route) => ({ query: route.query.q })},
    {path:"/",component:Home},
]

export {
    routes
}