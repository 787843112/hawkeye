<template>
  <div id="app">
    <keep-alive>
      <router-view />
    </keep-alive>
    <Menu v-if="$route.meta.showMenu"></Menu>
  </div>
</template>

<script>
import Menu from './components/Menu'
export default {
  name: 'App',
  components: {
    Menu,
  },
  watch: {
    $route(to, from) {
      if (this.$store.state.adminInfo.AdminID == null) {
        this.$store.dispatch('toGetAdminInfo')
      }
      if (this.$store.state.adminInfo.AdminID != null) {
        if (from.path === '/' && to.path !== '/movie' && to.path !== '/cinema') {
          this.$router.replace('/movie')
        }
      } else {
        this.$router.replace('/login')
      }
    },
  },
  created() {
    this.$store.dispatch('toGetAdminInfo')
  },
}
</script>
