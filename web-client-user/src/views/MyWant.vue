<!-- 想看页面 -->
<template>
  <div id="myWant">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" class="myWant-header">
        想&nbsp;看&nbsp;的&nbsp;电&nbsp;影
      </h1>
    </van-nav-bar>
    <van-sticky :offset-top="46">
      <van-notice-bar color="#1989fa" background="#ecf9ff" left-icon="coupon-o" :scrollable="false"
        >共&nbsp;<span style="color: crimson;">{{ count }}</span
        >&nbsp;部影片</van-notice-bar
      >
    </van-sticky>
    <!-- 电影列表 -->
    <div class="want-list-div">
      <van-pull-refresh v-model="wantIsLoading" @refresh="wantOnRefresh">
        <van-list
          :offset="5"
          v-model="wantLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :immediate-check="false"
          :finished="wantFinished"
          @load="wantOnLoad"
          :error.sync="wantError"
          error-text="请求失败，点击重新加载"
        >
          <div v-for="(item, index) in WantList" :key="item.MovieID">
            <van-swipe-cell right-width="50" class="want-cell">
              <!-- 电影信息div -->
              <div @click="toMovieInfo(item)" class="want-cell-movie-info">
                <div class="movie-info-img">
                  <img :src="'static/posters/' + item.Poster" />
                </div>
                <div id="movie-info-div">
                  <b>{{ item.MovieName }}</b>
                  <p>
                    <span style="color: #d38208;height: 21px;">
                      {{ computedWantNumber(item.WantNumber) }}
                    </span>
                    &nbsp;人想看
                  </p>
                  <p>{{ item.Director }}</p>
                  <p>{{ item.ReleaseTime }}&nbsp;在中国大陆上映</p>
                </div>
              </div>
              <template #right>
                <van-button
                  @click="delectWant(item.MovieID, index)"
                  square
                  text="删除"
                  type="danger"
                  style="height: 100%"
                />
              </template>
            </van-swipe-cell>
            <van-divider />
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
  </div>
</template>

<script>
import { reqGetWant, reqDeleteWant } from '../api'
export default {
  name: 'MyWant',
  data() {
    return {
      Page: 1,
      Size: 6,
      count: 0,
      WantList: [],
      wantLoading: false,
      wantFinished: false,
      wantIsLoading: false,
      wantError: false,
    }
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/myCenter')
    },
    //计算想看人数
    computedWantNumber(wantNumber) {
      if (wantNumber > 100000) {
        return (wantNumber / 10000).toFixed(1) + '万'
      } else {
        return wantNumber
      }
    },
    //去查看电影详细信息
    async toMovieInfo(movie) {
      await this.$store.dispatch('recordMovieInfo', movie)
      this.$router.replace('/movieInfo')
    },
    //加载更多电影票
    wantOnLoad() {
      this.Page++
      this.wantError = false
      this.getWantList()
    },
    //下拉刷新电影票列表
    wantOnRefresh() {
      this.Page = 1
      this.wantIsLoading = true
      this.wantFinished = false
      this.getWantList()
    },
    //获取电影票
    async getWantList() {
      let result
      const { Page, Size } = this
      result = await reqGetWant({ Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.WantList = result.data
          this.count = result.count
        } else {
          this.WantList = this.WantList.concat(result.data)
        }
        this.wantLoading = false
        this.wantIsLoading = false
      } else if (result.code == -1) {
        this.wantLoading = false
        this.wantIsLoading = false
        this.wantFinished = true
      } else {
        console.log(result.message)
        this.wantError = true
        this.wantIsLoading = false
        this.wantLoading = false
      }
    },
    //删除某部看过的电影
    async delectWant(MovieID, index) {
      let result = await reqDeleteWant({ MovieID })
      if (result.code == 0) {
        this.$toast('删除成功')
        this.WantList.splice(index, 1)
        this.count--
        if (this.WantList.length < this.Size) {
          this.Page = 1
          this.Size = 6
          this.getWantList()
        }
      } else {
        this.$toast(result.message)
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getWantList()
  },
}
</script>
<style>
@import url('../assets/css/myWant.css');
</style>
