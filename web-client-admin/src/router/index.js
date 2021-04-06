import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/movie',
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '../views/Login'),
    meta: {
      showMenu: false,
    },
  },
  {
    path: '/movie',
    name: 'Movie',
    component: () => import(/* webpackChunkName: "movie" */ '../views/Movie'),
    meta: {
      showMenu: true,
    },
  },
  {
    path: '/cinema',
    name: 'Cinema',
    component: () => import(/* webpackChunkName: "cinema" */ '../views/Cinema'),
    meta: {
      showMenu: true,
    },
  },
  {
    path: '/hall',
    name: 'Hall',
    component: () => import(/* webpackChunkName: "hall" */ '../views/Hall'),
    meta: {
      showMenu: true,
    },
  },
  {
    path: '/seat',
    name: 'Seat',
    component: () => import(/* webpackChunkName: "seat" */ '../views/Seat'),
    meta: {
      showMenu: true,
    },
  },
  {
    path: '/screening',
    name: 'Screening',
    component: () => import(/* webpackChunkName: "screening" */ '../views/Screening'),
    meta: {
      showMenu: true,
    },
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import(/* webpackChunkName: "admin" */ '../views/Admin'),
    meta: {
      showMenu: true,
    },
  },
]

const router = new VueRouter({
  mode: 'history',
  routes,
})

export default router
