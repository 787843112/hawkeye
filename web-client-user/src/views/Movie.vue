<!-- 电影首页 -->
<template>
  <div id="movie">
    <van-nav-bar fixed style="background-color: rgb(84, 157, 240)"><h1 slot="title">鹰 眼 电 影</h1></van-nav-bar>
    <van-sticky :offset-top="46">
      <div class="movie-tab">
        <van-row>
          <van-col span="5">
            <CitySelect />
          </van-col>
          <van-col span="14">
            <van-tabs v-model="tabShow" @click="onClickTab">
              <van-tab title="正在热映"></van-tab>
              <van-tab title="即将上映"></van-tab>
            </van-tabs>
          </van-col>
          <van-col span="5" class="movie-tab-search">
            <p @click="toSearchMovie">
              <van-icon name="search"></van-icon>
            </p>
          </van-col>
        </van-row>
      </div>
    </van-sticky>
    <!--正在热映列表-->
    <div v-if="tabShow == 0" class="movie-list">
      <van-pull-refresh v-model="releaseIsLoading" @refresh="releaseOnRefresh">
        <van-list
          :offset="10"
          :immediate-check="false"
          v-model="releaseLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :finished="releaseFinished"
          @load="releaseOnLoad"
          :error.sync="releaseError"
          error-text="请求失败，点击重新加载"
        >
          <van-cell v-for="item in releaseMovieList" :key="item.MovieID" @click="toMovieDetail(item)">
            <div class="movie-cell">
              <div class="movie-img">
                <img :src="'static/posters/' + item.Poster" />
              </div>
              <div class="movie-info">
                <div class="movie-info_2">
                  <h2>{{ item.MovieName }}</h2>
                  <h4>类型: {{ item.MovieTypes }}</h4>
                  <h4>主演: {{ item.Starring }}</h4>
                  <h4>简介:{{ item.Introduction }}</h4>
                </div>
                <div class="button-block">
                  <div class="button-block_2">
                    <span class="fix">购票</span>
                  </div>
                </div>
              </div>
            </div>
          </van-cell>
        </van-list>
      </van-pull-refresh>
      <div style="height: 50px;"></div>
    </div>
    <!--即将上映列表-->
    <div v-if="tabShow == 1" class="movie-list">
      <van-pull-refresh v-model="unReleaseIsLoading" @refresh="unReleaseOnRefresh">
        <!--近期最受期待列表-->
        <div style="height: 200px;width: 100%">
          <p
            style="margin-top: 5px;margin-left: 20px;color: #ff0073;text-overflow: ellipsis;white-space: nowrap;overflow: hidden"
          >
            近期最受期待
          </p>
          <div style="overflow: scroll;white-space: nowrap;margin-top: 10px;margin-left: 10px;margin-right: 10px">
            <div
              v-for="item in expectMovieList"
              style="overflow: hidden;margin-right: 10px;display: inline-block;"
              :key="item.MovieID"
              @click="toMovieInfo(item)"
            >
              <div style="position: relative;">
                <img style="height: 115px;width: 85px" :src="'static/posters/' + item.Poster" />
              </div>
              <h5 style="width: 85px;text-overflow: ellipsis;white-space: nowrap;overflow: hidden;">
                {{ item.MovieName }}
              </h5>
              <p style="font-size: 12px;color: #999">{{ computedWantNumber(item.WantNumber) }}人想看</p>
            </div>
          </div>
        </div>
        <!--近期即将上映列表-->
        <div>
          <p
            style="margin-top: 5px;margin-left: 20px;color: #ff0073;text-overflow: ellipsis;white-space: nowrap;overflow: hidden"
          >
            近期即将上映
          </p>
          <van-list
            :offset="10"
            v-model="unReleaseLoading"
            finished-text="没有更多了o(ﾟДﾟ)っ!"
            :finished="unReleaseFinished"
            @load="unReleaseOnLoad"
            :error.sync="unReleaseError"
            error-text="请求失败，点击重新加载"
          >
            <van-cell v-for="item in unReleaseMovieList" :key="item.MovieID">
              <div class="movie-cell">
                <div class="movie-img" @click="toMovieInfo(item)">
                  <img :src="'static/posters/' + item.Poster" />
                </div>
                <div class="movie-info" @click="toMovieInfo(item)">
                  <div class="movie-info_2">
                    <h2>{{ item.MovieName }}</h2>
                    <span style="color: #faaf00;size: 15px">{{ computedWantNumber(item.WantNumber) }}</span>
                    <span style="font-size: 13px;color: #666">人想看</span>
                    <div
                      style="font-size: 13px;color: #666;width: 70%;text-overflow: ellipsis;white-space: nowrap;overflow: hidden"
                    >
                      主演: {{ item.Starring }}
                    </div>
                    <div style="font-size: 13px;color: #666">
                      {{ item.ReleaseTime != '' ? item.ReleaseTime : '即将' }}上映
                    </div>
                  </div>
                </div>
                <div class="button-block-want">
                  <div v-if="isWant(item)" class="button-block-want_1" @click="deleteWantMovie(item)">
                    <span class="fix">已想看</span>
                  </div>
                  <div v-else class="button-block-want_2" @click="addWantMovie(item)">
                    <span class="fix">想看</span>
                  </div>
                </div>
              </div>
            </van-cell>
          </van-list>
        </div>
      </van-pull-refresh>
      <div style="height: 50px;"></div>
    </div>
  </div>
</template>

<script>
import CitySelect from '../components/CitySelect'
import { reqGetMovies, reqExpectMovies, reqGetWantMovieID, reqAddWant, reqDeleteWant } from '../api'
export default {
  name: 'Movie',
  components: { CitySelect },
  data() {
    return {
      releasePage: 1,
      unReleasePage: 1,
      Size: 7,
      tabShow: 0,
      wantMovieID: [],
      releaseMovieList: [], //正在上映的电影列表
      unReleaseMovieList: [], //近期即将上映的电影列表
      expectMovieList: [], //近期最受期待的电影列表
      releaseIsLoading: false,
      releaseLoading: false,
      releaseFinished: false,
      releaseError: false,
      unReleaseIsLoading: false,
      unReleaseLoading: false,
      unReleaseFinished: false,
      unReleaseError: false,
    }
  },
  //方法集合
  methods: {
    //去搜索电影页面
    toSearchMovie() {
      this.$router.replace('/searchMovie')
    },
    //去详细页面
    async toMovieDetail(movie) {
      await this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieDetail')
    },
    //去查看电影信息页面
    async toMovieInfo(movie) {
      await this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieInfo')
    },
    //计算想看人数
    computedWantNumber(WantNumber) {
      if (WantNumber > 100000) {
        return (WantNumber / 10000).toFixed(1) + '万'
      } else {
        return WantNumber
      }
    },
    onClickTab(name) {
      if (name == 1 && this.expectMovieList.length == 0) {
        this.getExpectMovieList()
      }
      if (name == 1 && this.unReleaseMovieList.length == 0) {
        this.getUnReleaseMovies()
      }
      this.tabShow = name
    },
    //加载更多正在上映的电影
    releaseOnLoad() {
      this.releasePage++
      this.releaseError = false
      this.getReleaseMovies()
    },
    //下拉刷新正在上映的电影
    releaseOnRefresh() {
      this.releasePage = 1
      this.releaseIsLoading = true
      this.releaseFinished = false
      this.getReleaseMovies()
    },
    //加载更多即将上映的电影
    unReleaseOnLoad() {
      this.unReleasePage++
      this.unReleaseError = false
      this.getUnReleaseMovies()
    },
    //下拉刷新最受期待的10部电影和即将上映的电影
    unReleaseOnRefresh() {
      this.unReleasePage = 1
      this.unReleaseIsLoading = true
      this.unReleaseFinished = false
      this.getExpectMovieList()
      this.getWantMovieID()
      this.getUnReleaseMovies()
    },
    //判断电影是否已想看
    isWant(movie) {
      return this.wantMovieID.some((item) => {
        return item == movie.MovieID
      })
    },
    //获取正在上映的电影
    async getReleaseMovies() {
      let result
      const MovieName = ''
      const MovieType = ''
      const Status = 2
      const Page = this.releasePage
      const Size = this.Size
      result = await reqGetMovies({ MovieName, MovieType, Status, Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.releaseMovieList = result.data
        } else {
          this.releaseMovieList = this.releaseMovieList.concat(result.data)
        }
        this.releaseLoading = false
        this.releaseIsLoading = false
      } else if (result.code == -1) {
        this.releaseIsLoading = false
        this.releaseLoading = false
        this.releaseFinished = true
      } else {
        console.log(result.message)
        this.releaseError = true
        this.releaseIsLoading = false
        this.releaseLoading = false
      }
    },
    //获取最受期待的10部电影
    async getExpectMovieList() {
      let result
      result = await reqExpectMovies()
      if (result.code == 0) {
        this.expectMovieList = result.data
      } else {
        this.expectMovieList = []
        console.log('获取近期最受期待失败')
      }
    },
    //获取即将上映的电影
    async getUnReleaseMovies() {
      let result
      const MovieName = ''
      const MovieType = ''
      const Status = 3
      const Page = this.unReleasePage
      const Size = this.Size
      result = await reqGetMovies({ MovieName, MovieType, Status, Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.unReleaseMovieList = result.data
        } else {
          this.unReleaseMovieList = this.unReleaseMovieList.concat(result.data)
        }
        this.unReleaseIsLoading = false
        this.unReleaseLoading = false
      } else if (result.code == -1) {
        this.unReleaseIsLoading = false
        this.unReleaseFinished = true
        this.unReleaseLoading = false
      } else {
        console.log(result.message)
        this.unReleaseError = true
        this.unReleaseIsLoading = false
        this.unReleaseLoading = false
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
        this.$toast('请您先登录!')
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
    //根据想看人数排序
    // contrast(obj1, obj2) {
    //   return obj2.WantNumber - obj1.WantNumber
    // },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getReleaseMovies()
    if (this.$store.state.userInfo.UserID != null) {
      this.getWantMovieID()
    }
  },
}
</script>
<style>
@import url('../assets/css/movie.css');
</style>
