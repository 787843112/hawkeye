<template>
  <div id="app">
    <keep-alive>
      <router-view />
    </keep-alive>
    <Footer v-show="$route.meta.showFooter" />
  </div>
</template>

<script>
import Footer from './components/FooterGuide'
import { mapActions } from 'vuex'

export default {
  name: 'App',
  components: {
    Footer,
  },
  watch: {
    $route(to, from) {
      this.$store.dispatch('recordBackPath', from.path)
      if (from.path === '/' && to.path !== '/movie') {
        this.$router.replace('/movie')
      }
    },
  },
  async mounted() {
    this.toGetUserInfo()
  },
  methods: {
    ...mapActions(['toGetUserInfo']),
  },
}
</script>
