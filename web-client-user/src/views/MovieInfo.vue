<!-- 电影信息页面 -->
<template>
  <div id="movieInfo">
    <div>
      <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
        <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
        <h1 slot="title" class="movie-info-header">
          {{ MovieInfo.MovieName }}
        </h1>
      </van-nav-bar>
    </div>
    <div class="movie-info-detail">
      <div class="movie-info-poster-bg">
        <img :src="'static/posters/' + MovieInfo.Poster" />
      </div>
      <div class="movie-info-box">
        <div class="movie-info-poster">
          <img :src="'static/posters/' + MovieInfo.Poster" />
        </div>
        <div class="movie-info-div">
          <div class="movie-info-name">{{ MovieInfo.MovieName }}</div>
          <div style="width: 60%;float: left">
            <div v-if="MovieInfo.Status == 3" class="movie-info-want-number">
              {{ computedWantNumber(MovieInfo.WantNumber) }}&nbsp;人想看
            </div>
            <div v-else class="movie-info-box-office">票房:{{ computedBoxOffice(MovieInfo.BoxOffice) }}</div>
            <div class="movie-info-type">{{ MovieInfo.MovieTypes }}</div>
            <div class="movie-info-time">{{ MovieInfo.Region }}/{{ MovieInfo.MovieLength }}分钟</div>
            <div class="movie-info-release">{{ MovieInfo.ReleaseTime }}大陆上映</div>
          </div>
          <div v-show="MovieInfo.Total >= 1000" class="movie-info-score">
            <p class="movie-info-score-p">
              {{ MovieInfo.Score.toFixed(1) }}
            </p>
            <van-rate v-model="stars" readonly allow-half gutter="0" />
            <p class="movie-info-total">{{ computedTotal(MovieInfo.Total) }}人评</p>
          </div>
          <div v-show="1000 > MovieInfo.Total" class="movie-info-score">
            <p class="movie-info-score-p2">
              暂无评分
            </p>
            <p class="movie-info-score-p3">评分人数不足</p>
          </div>
        </div>
      </div>
    </div>
    <van-divider content-position="left" :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }"
      >简介
    </van-divider>
    <div class="movie-info-introduction">
      {{ MovieInfo.Introduction }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'MovieInfo',
  computed: {
    MovieInfo() {
      return this.$store.state.movieInfo
    },
    stars() {
      return Number((this.MovieInfo.Score / 2).toFixed(1))
    },
  },
  //方法集合
  methods: {
    onClickLeft() {
      this.$router.replace(this.$store.state.backPath)
    },
    //计算票房
    computedBoxOffice(boxOffice) {
      if (boxOffice > 100000000) {
        return (boxOffice / 100000000).toFixed(2) + '亿'
      } else if (boxOffice < 100000000 && boxOffice > 10000) {
        return (boxOffice / 10000).toFixed(1) + '万'
      } else {
        return boxOffice
      }
    },
    //计算想看人数
    computedWantNumber(WantNumber) {
      if (WantNumber > 100000) {
        return (WantNumber / 10000).toFixed(1) + '万'
      } else {
        return WantNumber
      }
    },
    //计算评价人数
    computedTotal(total) {
      if (total > 100000) {
        return (total / 10000).toFixed(1) + '万'
      } else {
        return total
      }
    },
  },
}
</script>
<style>
@import url('../assets/css/movieInfo.css');
</style>
