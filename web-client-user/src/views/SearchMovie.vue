<!-- 搜索电影页面 -->
<template>
  <div id="searchMovie">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title">鹰 眼 电 影</h1>
    </van-nav-bar>
    <van-sticky :offset-top="46">
      <van-search v-model="MovieName" placeholder="搜电影" shape="round" background="#4fc08d"></van-search>
    </van-sticky>
    <div class="searchMovie-list">
      <van-pull-refresh v-model="searchMovieIsLoading" @refresh="searchMovieOnRefresh">
        <van-list
          :offset="10"
          :immediate-check="false"
          v-model="searchMovieLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :finished="searchMovieFinished"
          @load="searchMovieOnLoad"
          :error.sync="searchMovieError"
          error-text="请求失败，点击重新加载"
        >
          <van-cell v-for="item in MovieList" :key="item.MovieID">
            <div class="searchMovie-cell">
              <div class="searchMovie-img" @click="toMovieInfo(item)">
                <img :src="'static/posters/' + item.Poster" />
              </div>
              <div class="searchMovie-info">
                <div class="searchMovie-info_2" @click="toMovieInfo(item)">
                  <h2>{{ item.MovieName }}</h2>
                  <h4>类型:{{ item.MovieTypes }}</h4>
                  <h5 style="font-size: 13px;color: #999999">
                    {{ item.ReleaseTime != '' ? item.ReleaseTime : '即将' }}上映
                  </h5>
                  <span v-if="item.Status == 1" style="color: #cec6c3;size: 15px"
                    >累计票房:{{ computedBoxOffice(item.BoxOffice) }}</span
                  >
                  <span v-if="item.Status == 2" style="color: #fa4c1f;size: 15px"
                    >当前票房:{{ computedBoxOffice(item.BoxOffice) }}</span
                  >
                  <span v-if="item.Status == 3" style="color: #faaf00;size: 15px">{{ item.WantNumber }}</span>
                  <span v-if="item.Status == 3" style="font-size: 13px;color: #666">人想看</span>
                </div>
                <div v-if="item.Status == 2" class="ticket-button">
                  <div class="ticket-button_2">
                    <span class="ticket-span" @click="toMovieDetail(item)"> 购票</span>
                  </div>
                </div>
                <div v-if="item.Status == 1" class="shelves" @click="toMovieInfo(item)">
                  <div class="shelves_2">
                    <span class="shelves-span">已下架</span>
                  </div>
                </div>
                <div v-if="item.Status == 3" class="want-button">
                  <div v-if="isWant(item)" class="button-block-want_1">
                    <span class="fix" @click="deleteWantMovie(item)">已想看</span>
                  </div>
                  <div v-else class="button-block-want_2">
                    <span class="fix" @click="addWantMovie(item)">想看</span>
                  </div>
                </div>
              </div>
            </div>
          </van-cell>
        </van-list>
      </van-pull-refresh>
    </div>
  </div>
</template>

<script>
import { reqGetMovies, reqGetWantMovieID, reqAddWant, reqDeleteWant } from '../api'
export default {
  name: 'SearchMovie',
  data() {
    return {
      MovieName: '',
      Page: 1,
      Size: 10,
      searchMovieLoading: false,
      searchMovieFinished: false,
      searchMovieIsLoading: false,
      searchMovieError: false,
      wantMovieID: [],
      MovieList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {},
  //监控data中的数据变化
  watch: {
    MovieName() {
      this.Page = 1
      this.MovieList = []
      this.searchMovieFinished = false
      this.getSearchMovie()
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/movie')
    },
    //去查看电影详细信息
    toMovieInfo(movie) {
      this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieInfo')
    },
    //去购票
    toMovieDetail(movie) {
      this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieDetail')
    },
    //计算票房
    computedBoxOffice(boxOffice) {
      if (boxOffice > 100000000) {
        return (boxOffice / 100000000).toFixed(2) + '亿'
      } else if (boxOffice < 100000000 && boxOffice > 10000) {
        return (boxOffice / 10000).toFixed(0) + '万'
      } else {
        return boxOffice
      }
    },
    //加载更多电影
    searchMovieOnLoad() {
      this.Page++
      this.searchMovieError = false
      this.getSearchMovie()
    },
    //下拉刷新电影列表
    searchMovieOnRefresh() {
      this.Page = 1
      this.searchMovieIsLoading = true
      this.searchMovieFinished = false
      this.getSearchMovie()
    },
    //判断电影是否已想看
    isWant(movie) {
      return this.wantMovieID.some((item) => {
        return item == movie.MovieID
      })
    },
    //搜索电影
    async getSearchMovie() {
      let result
      const { MovieName, Page, Size } = this
      const MovieType = ''
      const Status = 0
      result = await reqGetMovies({ MovieName, MovieType, Status, Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.MovieList = result.data
        } else {
          this.MovieList = this.MovieList.concat(result.data)
        }
        this.searchMovieLoading = false
        this.searchMovieIsLoading = false
      } else if (result.code == -1) {
        this.searchMovieLoading = false
        this.searchMovieIsLoading = false
        this.searchMovieFinished = true
      } else {
        console.log(result.message)
        this.searchMovieError = true
        this.searchMovieLoading = false
        this.searchMovieIsLoading = false
      }
    },
    //获取用户所有想看的电影ID
    async getWantMovieID() {
      let result = await reqGetWantMovieID()
      if (result.code == 0) {
        this.wantMovieID = result.data
      } else if (result.code == -1) {
        this.wantMovieID = []
      } else {
        console.log(result.message)
      }
    },
    //添加电影到想看
    async addWantMovie(movie) {
      if (this.$store.state.userInfo.UserID == null) {
        this.$toast.fail('请您先登录!')
        return
      }
      let result
      const MovieID = movie.MovieID
      result = await reqAddWant({ MovieID })
      if (result.code == 0) {
        this.wantMovieID.unshift(MovieID)
        movie.WantNumber++
      } else {
        this.$toast(result.message)
      }
    },
    //取消此想看的电影
    async deleteWantMovie(movie) {
      let result
      const MovieID = movie.MovieID
      result = await reqDeleteWant({ MovieID })
      if (result.code == 0) {
        this.wantMovieID.some((item, i) => {
          if (MovieID == item) {
            this.wantMovieID.splice(i, 1)
          }
        })
        movie.WantNumber--
      } else {
        this.$toast(result.message)
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getWantMovieID()
  },
}
</script>
<style>
@import url('../assets/css/searchMovie.css');
</style>
