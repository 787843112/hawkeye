<!-- 看过页面 -->
<template>
  <div id="mySeen">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" class="mySeen-header">
        看&nbsp;过&nbsp;的&nbsp;电&nbsp;影
      </h1>
    </van-nav-bar>
    <van-sticky :offset-top="46">
      <van-notice-bar color="#1989fa" background="#ecf9ff" left-icon="coupon-o" :scrollable="false"
        >共&nbsp;<span style="color: crimson;">{{ count }}</span
        >&nbsp;部影片<span style="color: grey;font-size: 12px;">(购过票标记为看过)</span></van-notice-bar
      >
    </van-sticky>
    <!-- 电影列表 -->
    <div class="seen-list-div">
      <van-pull-refresh v-model="seenIsLoading" @refresh="seenOnRefresh">
        <van-list
          :offset="5"
          v-model="seenLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :immediate-check="false"
          :finished="seenFinished"
          @load="seenOnLoad"
          :error.sync="seenError"
          error-text="请求失败，点击重新加载"
        >
          <div v-for="(item, index) in SeenList" :key="item.MovieInfo.MovieID">
            <van-swipe-cell right-width="50" class="seen-cell">
              <!-- 电影信息div -->
              <div @click="toMovieInfo(item.MovieInfo)" class="seen-cell-movie-info">
                <div class="movie-info-img">
                  <img :src="'static/posters/' + item.MovieInfo.Poster" />
                </div>
                <div class="movie-info-div">
                  <b>{{ item.MovieInfo.MovieName }}</b>
                  <p>{{ item.MovieInfo.ReleaseTime }}&nbsp;上映</p>
                  <p style="margin-top: 15px">导演:&nbsp;{{ item.MovieInfo.Director }}</p>
                  <p>主演:&nbsp;{{ item.MovieInfo.Starring }}</p>
                </div>
              </div>
              <template #right>
                <van-button @click="clickDelete(item, index)" square text="删除" type="danger" style="height: 100%" />
              </template>
            </van-swipe-cell>
            <!-- 我的评分div -->
            <div class="score-info-div" @click="toEvaluation(item)">
              <span style="color: grey;font-size: 18px;margin-left: 10px;">我评</span>
              <!-- count="10" allow-half -->
              <van-rate
                style="margin-top: 12.5px;margin-left: 20px;"
                v-model="item.Score"
                readonly
                count="10"
                void-icon="star"
                void-color="#eee"
              />
            </div>
            <van-divider />
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
  </div>
</template>

<script>
import { reqGetSeen, reqDeleteSeen } from '../api'
export default {
  name: 'MySeen',
  data() {
    //这里存放数据
    return {
      Page: 1,
      Size: 5,
      SeenList: [],
      count: 0,
      seenLoading: false,
      seenFinished: false,
      seenIsLoading: false,
      seenError: false,
    }
  },
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/myCenter')
    },
    clickDelete(seen, index) {
      this.$dialog
        .confirm({
          title: '删除后无法恢复,确认吗?',
        })
        .then(() => {
          this.delectOrder(seen.MovieID, index)
        })
        .catch(() => {})
    },
    //去查看电影详细信息
    async toMovieInfo(movie) {
      await this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieInfo')
    },
    async toEvaluation(seen) {
      await this.$store.dispatch('recordSeenInfo', seen)
      this.$router.replace('/evaluation')
    },
    //加载更多电影票
    seenOnLoad() {
      this.Page++
      this.seenError = false
      this.getSeenList()
    },
    //下拉刷新电影票列表
    seenOnRefresh() {
      this.Page = 1
      this.seenIsLoading = true
      this.seenFinished = false
      this.getSeenList()
    },
    //获取电影票
    async getSeenList() {
      let result
      const { Page, Size } = this
      result = await reqGetSeen({ Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.SeenList = result.data
          this.count = result.count
        } else {
          this.SeenList = this.SeenList.concat(result.data)
        }
        this.seenLoading = false
        this.seenIsLoading = false
      } else if (result.code == -1) {
        this.seenLoading = false
        this.seenIsLoading = false
        this.seenFinished = true
      } else {
        console.log(result.message)
        this.seenError = true
        this.seenIsLoading = false
        this.seenLoading = false
      }
    },
    //删除某部看过的电影
    async delectSeen(MovieID, index) {
      let result = await reqDeleteSeen({ MovieID })
      if (result.code == 0) {
        this.$toast('删除成功!')
        this.SeenList.splice(index, 1)
        if (this.SeenList.length < this.Size) {
          this.Page = 1
          this.Size = 5
          this.getSeenList()
        }
      } else {
        this.$toast(result.message)
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getSeenList()
  },
}
</script>
<style>
@import url('../assets/css/mySeen.css');
</style>
