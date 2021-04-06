import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/movie',
  },
  {
    path: '/movie',
    name: 'Movie',
    component: () => import(/* webpackChunkName: "movie" */ '../views/Movie'),
    meta: {
      showFooter: true,
    },
  },
  {
    path: '/cinema',
    name: 'Cinema',
    component: () => import(/* webpackChunkName: "cinema" */ '../views/Cinema'),
    meta: {
      showFooter: true,
    },
  },
  {
    path: '/myCenter',
    name: 'MyCenter',
    component: () =>
      import(/* webpackChunkName: "myCenter" */ '../views/MyCenter'),
    meta: {
      showFooter: true,
    },
  },
  {
    path: '/searchMovie',
    name: 'SearchMovie',
    component: () =>
      import(/* webpackChunkName: "searchMovie" */ '../views/SearchMovie'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/searchCinema',
    name: 'SearchCinema',
    component: () =>
      import(/* webpackChunkName: "searchCinema" */ '../views/SearchCinema'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '../views/Login'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/movieDetail',
    name: 'MovieDetail',
    component: () =>
      import(/* webpackChunkName: "movieDetail" */ '../views/MovieDetail'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/movieInfo',
    name: 'MovieInfo',
    component: () =>
      import(/* webpackChunkName: "movieInfo" */ '../views/MovieInfo'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/shows',
    name: 'Shows',
    component: () => import(/* webpackChunkName: "shows" */ '../views/Shows'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/selectSeat',
    name: 'SelectSeat',
    component: () =>
      import(/* webpackChunkName: "selectSeat" */ '../views/SelectSeat'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/createOrder',
    name: 'CreateOrder',
    component: () =>
      import(/* webpackChunkName: "createOrder" */ '../views/CreateOrder'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/myOrder',
    name: 'MyOrder',
    component: () =>
      import(/* webpackChunkName: "myOrder" */ '../views/MyOrder'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/orderDetail',
    name: 'OrderDetail',
    component: () =>
      import(/* webpackChunkName: "orderDetail" */ '../views/OrderDetail'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/mySeen',
    name: 'MySeen',
    component: () => import(/* webpackChunkName: "mySeen" */ '../views/MySeen'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/myWant',
    name: 'MyWant',
    component: () => import(/* webpackChunkName: "myWant" */ '../views/MyWant'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/myInfo',
    name: 'MyInfo',
    component: () => import(/* webpackChunkName: "myInfo" */ '../views/MyInfo'),
    meta: {
      showFooter: false,
    },
  },
  {
    path: '/evaluation',
    name: 'Evaluation',
    component: () =>
      import(/* webpackChunkName: "evaluation" */ '../views/Evaluation'),
    meta: {
      showFooter: false,
    },
  },
]

const router = new VueRouter({
  mode: 'history',
  routes,
})

export default router
