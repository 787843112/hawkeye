<!-- 影院页面 -->
<template>
  <div id="cinema">
    <van-nav-bar fixed style="background-color: rgb(84, 157, 240)"><h1 slot="title">影 院</h1></van-nav-bar>
    <van-sticky :offset-top="46">
      <div class="cinema-tab">
        <van-row class="cinema-tab-row" style="height: 44px;">
          <van-col span="5">
            <CitySelect />
          </van-col>
          <van-col span="19">
            <div class="search" @click="toSearchCinema">
              <van-icon name="search"></van-icon>
              <span>搜影院</span>
            </div>
          </van-col>
        </van-row>
      </div>
    </van-sticky>
    <div class="cinema-list">
      <van-pull-refresh v-model="cinemaIsLoading" @refresh="cinemaOnRefresh">
        <van-list
          :offset="50"
          v-model="cinemaLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :immediate-check="false"
          :finished="cinemaFinished"
          @load="cinemaOnLoad"
          :error.sync="cinemaError"
          error-text="请求失败，点击重新加载"
        >
          <van-cell v-for="item in CinemaList" :key="item.CinemaID" @click="toShows(item)">
            <div class="cinema-cell">
              <h3>{{ item.CinemaName }}</h3>
              <van-icon name="location-o"></van-icon>
              <span>{{ item.Address }}</span>
            </div>
          </van-cell>
        </van-list>
      </van-pull-refresh>
      <div style="height: 50px;"></div>
    </div>
  </div>
</template>

<script>
import CitySelect from '../components/CitySelect'
import { reqGetCinemas } from '../api'
export default {
  name: 'Cinema',
  components: { CitySelect },
  data() {
    return {
      Page: 1,
      Size: 15,
      CinemaList: [],
      cinemaLoading: false,
      cinemaFinished: false,
      cinemaIsLoading: false,
      cinemaError: false,
    }
  },
  //计算属性
  computed: {
    City() {
      return this.$store.state.city
    },
  },
  //监控data中的数据变化
  watch: {
    City() {
      this.Page = 1
      this.CinemaList = []
      this.cinemaFinished = false
      this.getCinemaList()
    },
  },
  //方法集合
  methods: {
    toSearchCinema() {
      this.$router.replace('/searchCinema')
    },
    async toShows(cinema) {
      if (this.$store.state.movieList.length == 0) {
        await this.$store.dispatch('toGetMovieList')
      }
      await this.$store.dispatch('recordMovieInfo', {})
      await this.$store.dispatch('recordCinemaInfo', cinema)
      this.$router.replace('/shows')
    },
    //加载更多影院
    cinemaOnLoad() {
      this.Page++
      this.cinemaError = false
      this.getCinemaList()
    },
    //下拉刷新影院列表
    cinemaOnRefresh() {
      this.Page = 1
      this.cinemaIsLoading = true
      this.cinemaFinished = false
      this.getCinemaList()
    },
    //获取影院列表
    async getCinemaList() {
      let result
      const CinemaName = ''
      const { City, Page, Size } = this
      result = await reqGetCinemas({ CinemaName, City, Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.CinemaList = result.data
        } else {
          this.CinemaList = this.CinemaList.concat(result.data)
        }
        this.cinemaLoading = false
        this.cinemaIsLoading = false
      } else if (result.code == -1) {
        this.cinemaLoading = false
        this.cinemaIsLoading = false
        this.cinemaFinished = true
      } else {
        console.log(result.message)
        this.cinemaError = true
        this.cinemaIsLoading = false
        this.cinemaLoading = false
      }
    },
  },
  created() {
    this.getCinemaList()
  },
}
</script>
<style>
@import url('../assets/css/cinema.css');
</style>
