<!-- 搜索影院页面 -->
<template>
  <div id="searchCinema">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title">影 院</h1>
    </van-nav-bar>
    <van-sticky :offset-top="46">
      <van-search v-model="CinemaName" placeholder="搜影院" shape="round" background="#4fc08d"></van-search>
    </van-sticky>
    <div class="searchCinema-list">
      <van-pull-refresh v-model="searchCinemaIsLoading" @refresh="searchCinemaOnRefresh">
        <van-list
          :offset="10"
          :immediate-check="false"
          v-model="searchCinemaLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :finished="searchCinemaFinished"
          @load="searchCinemaOnLoad"
          :error.sync="searchCinemaError"
          error-text="请求失败，点击重新加载"
        >
          <van-cell v-for="item in CinemaList" :key="item.CinemaID" @click="toShows(item)">
            <div class="searchCinema-cell">
              <h3>{{ item.CinemaName }}</h3>
              <van-icon name="location-o"></van-icon>
              <span>{{ item.Address }}</span>
            </div>
          </van-cell>
        </van-list>
      </van-pull-refresh>
    </div>
  </div>
</template>

<script>
import { reqGetCinemas } from '../api'
export default {
  name: 'SearchCinema',
  data() {
    //这里存放数据
    return {
      CinemaName: '',
      Page: 1,
      Size: 15,
      searchCinemaLoading: false,
      searchCinemaFinished: false,
      searchCinemaIsLoading: false,
      searchCinemaError: false,
      CinemaList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    City() {
      return this.$store.state.city
    },
  },
  //监控data中的数据变化
  watch: {
    CinemaName() {
      this.Page = 1
      this.CinemaList = []
      this.searchCinemaFinished = false
      this.getSearchCinema()
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/cinema')
    },
    //去购票
    async toShows(cinema) {
      if (this.$store.state.movieList.length == 0) {
        await this.$store.dispatch('toGetMovieList')
      }
      await this.$store.dispatch('recordMovieInfo', {})
      this.$store.dispatch('recordCinemaInfo', cinema)
      this.$router.replace('/shows')
    },
    //加载更多影院
    searchCinemaOnLoad() {
      this.Page++
      this.searchMovieError = false
      this.getSearchCinema()
    },
    //下拉刷新影院列表
    searchCinemaOnRefresh() {
      this.Page = 1
      this.searchCinemaIsLoading = true
      this.searchCinemaFinished = false
      this.getSearchCinema()
    },
    //获取搜索到的影院
    async getSearchCinema() {
      let result
      const { CinemaName, City, Page, Size } = this
      result = await reqGetCinemas({ CinemaName, City, Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.CinemaList = result.data
        } else {
          this.CinemaList = this.CinemaList.concat(result.data)
        }
        this.searchCinemaLoading = false
        this.searchCinemaIsLoading = false
      } else if (result.code == -1) {
        this.searchCinemaLoading = false
        this.searchCinemaIsLoading = false
        this.searchCinemaFinished = true
      } else {
        console.log(result.message)
        this.searchCinemaError = true
        this.searchCinemaLoading = false
        this.searchCinemaIsLoading = false
      }
    },
  },
}
</script>
<style>
@import url('../assets/css/searchCinema.css');
</style>
