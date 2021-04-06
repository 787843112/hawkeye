<!-- 选择场次页面 -->
<template>
  <div id="shows">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
        {{ CinemaInfo.CinemaName }}
      </h1>
    </van-nav-bar>
    <div class="cinema-info">
      <div class="cinema-name">{{ CinemaInfo.CinemaName }}</div>
      <div class="cinema-location">
        <van-icon name="location"></van-icon>
        {{ CinemaInfo.Address }}
      </div>
    </div>
    <div v-if="MovieList.length != 0" class="show-movie">
      <div class="movie-bg">
        <img :src="'static/posters/' + MovieList[movieActive].Poster" />
      </div>
      <div class="movie-swipe">
        <van-tabs v-model="movieActive" :border="false" :line-width="80">
          <van-tab v-for="item in MovieList" :key="item.MovieID">
            <div slot="title" style="text-align: center;">
              <img :src="'static/posters/' + item.Poster" />
            </div>
          </van-tab>
        </van-tabs>
      </div>
    </div>
    <van-sticky :offset-top="45">
      <div v-if="MovieList.length != 0" class="movieInfo-more">
        <div class="showMovie-info">
          <h3>{{ MovieList[movieActive].MovieName }}</h3>
          <div class="more-movieInfo">
            <span>{{ MovieList[movieActive].MovieLength }}分钟&nbsp;|</span>
            <span>{{ MovieList[movieActive].MovieTypes }}&nbsp;|</span>
            <span>{{ MovieList[movieActive].Starring }}</span>
          </div>
        </div>
      </div>
      <SelectDate v-if="MovieList.length != 0" />
    </van-sticky>
    <van-pull-refresh v-model="screeningIsLoading" @refresh="screeningOnRefresh">
      <div v-if="MovieList.length != 0" class="time-select">
        <van-list
          :offset="5"
          v-model="screeningLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :immediate-check="false"
          :finished="screeningFinished"
          @load="screeningOnLoad"
          :error.sync="screeningError"
          error-text="请求失败，点击重新加载"
        >
          <van-cell v-for="item in ScreeningList" :key="item.ScreeningID" @click="toSelectSeat(item)">
            <div class="time-cell">
              <div class="time-block">
                <div class="start-time">
                  {{ item.StartTime.slice(11, 16) }}
                </div>
                <div class="end-time">{{ item.EndTime.slice(11, 16) }}散场</div>
              </div>
              <div class="movie-hall">
                {{ item.HallName }}
              </div>
              <div class="movie-fare">¥{{ item.Price }}</div>
              <div class="button-ticket">
                <div class="button-ticket_2">
                  <span class="fix-1">选座位</span>
                </div>
              </div>
            </div>
          </van-cell>
        </van-list>
      </div>
    </van-pull-refresh>
  </div>
</template>

<script>
import SelectDate from '../components/SelectDate'
import { reqGetScreenings, reqGetStatus } from '../api'
export default {
  name: 'Shows',
  components: {
    SelectDate,
  },
  data() {
    //这里存放数据
    return {
      movieActive: 0,
      Page: 1,
      Size: 10,
      ScreeningList: [],
      screeningLoading: false,
      screeningFinished: false,
      screeningIsLoading: false,
      screeningError: false,
    }
  },
  //监听属性 类似于data概念
  computed: {
    MovieList() {
      return this.$store.state.movieList
    },
    MovieInfo() {
      return this.$store.state.movieInfo
    },
    CinemaInfo() {
      return this.$store.state.cinemaInfo
    },
    SelectDate() {
      return this.$store.state.selectDate
    },
  },
  //监控data中的数据变化
  watch: {
    $route() {
      this.movieIndex()
    },
    movieActive() {
      this.Page = 1
      this.ScreeningList = []
      this.screeningFinished = false
      this.getScreeningList()
    },
    SelectDate() {
      this.Page = 1
      this.ScreeningList = []
      this.screeningFinished = false
      this.getScreeningList()
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.ScreeningList = []
      this.$router.replace(this.$store.state.backPath == '/selectSeat' ? '/movie' : this.$store.state.backPath)
    },
    //去选择座位
    async toSelectSeat(screening) {
      await this.getSeatStatus(screening)
      this.$store.dispatch('recordSelectScreening', screening)
      this.$router.replace('/selectSeat')
    },
    movieIndex() {
      this.MovieList.some((item, i) => {
        if (item.MovieID == this.MovieInfo.MovieID) {
          this.movieActive = i
        }
      })
    },
    //加载更多场次
    screeningOnLoad() {
      this.Page++
      this.screeningError = false
      this.getScreeningList()
    },
    //下拉刷新场次列表
    screeningOnRefresh() {
      this.Page = 1
      this.ScreeningList = []
      this.screeningIsLoading = true
      this.screeningFinished = false
      this.getScreeningList()
    },
    //获取场次
    async getScreeningList() {
      let result
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const MovieID = this.MovieList[this.movieActive].MovieID
      const { SelectDate, Page, Size } = this
      result = await reqGetScreenings({
        CinemaID,
        MovieID,
        SelectDate,
        Page,
        Size,
      })
      if (result.code == 0) {
        if (Page == 1) {
          this.ScreeningList = result.data
        } else {
          this.ScreeningList = this.ScreeningList.concat(result.data)
        }
        this.screeningLoading = false
        this.screeningIsLoading = false
      } else if (result.code == -1) {
        this.screeningLoading = false
        this.screeningIsLoading = false
        this.screeningFinished = true
      } else {
        console.log(result.message)
        this.screeningError = true
        this.screeningIsLoading = false
        this.screeningLoading = false
      }
    },
    //获取座位信息
    async getSeatStatus(screening) {
      let result
      const ScreeningID = screening.ScreeningID
      result = await reqGetStatus({ ScreeningID })
      if (result.code == 0) {
        this.$store.dispatch('recordHallInfo', result.hall)
        this.$store.dispatch('recordSeatStatus', result.data)
      } else {
        console.log(result.message)
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.movieIndex()
  },
}
</script>
<style>
@import url('../assets/css/shows.css');
</style>
