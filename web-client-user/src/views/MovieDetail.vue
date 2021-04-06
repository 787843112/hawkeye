<!-- 细节页面 -->
<template>
  <div id="movieDetail">
    <div>
      <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
        <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
        <h1 slot="title" style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap">
          {{ MovieInfo.MovieName }}
        </h1>
      </van-nav-bar>
    </div>
    <div class="movie-detail" @click="toMovieInfo">
      <div class="poster-bg">
        <img :src="'static/posters/' + MovieInfo.Poster" />
      </div>
      <div class="detail-box">
        <div class="poster">
          <img :src="'static/posters/' + MovieInfo.Poster" />
        </div>
        <div class="content">
          <div class="content-title">{{ MovieInfo.MovieName }}</div>
          <div class="content-type">{{ MovieInfo.MovieTypes }}</div>
          <div class="content-time">{{ MovieInfo.Region }}/{{ MovieInfo.MovieLength }}分钟</div>
          <div class="content-release">{{ MovieInfo.ReleaseTime }}大陆上映</div>
        </div>
        <van-icon name="arrow" size="40" style="margin-top: 55px;overflow: hidden"></van-icon>
      </div>
    </div>
    <!-- 日期选择 -->
    <van-sticky :offset-top="45">
      <SelectDate />
    </van-sticky>
    <!-- 城市的影院列表 -->
    <van-pull-refresh v-model="cinemaIsLoading" @refresh="cinemaOnRefresh">
      <div class="select-cinema">
        <van-list
          :offset="50"
          :immediate-check="false"
          v-model="cinemaLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
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
      </div>
    </van-pull-refresh>
  </div>
</template>

<script>
import SelectDate from '../components/SelectDate'
import { reqGetCinemas } from '../api'
export default {
  name: 'MovieDetail',
  components: {
    SelectDate,
  },
  data() {
    return {
      Page: 1,
      Size: 10,
      CinemaList: [],
      cinemaLoading: false,
      cinemaFinished: false,
      cinemaIsLoading: false,
      cinemaError: false,
    }
  },
  computed: {
    City() {
      return this.$store.state.city
    },
    MovieInfo() {
      return this.$store.state.movieInfo
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
    //返回
    onClickLeft() {
      if (this.$store.state.backPath == '/movieInfo' || this.$store.state.backPath == '/shows') {
        this.$router.replace('/movie')
      } else {
        this.$router.replace(this.$store.state.backPath)
      }
    },
    //去查看电影信息页面
    toMovieInfo() {
      this.$router.replace('/movieInfo')
    },
    //去shows页面
    async toShows(cinema) {
      if (this.$store.state.movieList.length == 0) {
        await this.$store.dispatch('toGetMovieList')
      }
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
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getCinemaList()
  },
}
</script>
<style>
@import url('../assets/css/movieDetail.css');
</style>
