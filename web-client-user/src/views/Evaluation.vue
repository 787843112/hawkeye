<!-- 评价页面 -->
<template>
  <div id="evaluation">
    <van-nav-bar
      fixed
      left-arrow
      style="background-color: rgb(84, 157, 240)"
      @click-left="onClickLeft"
      @click-right="updateEvaluation"
    >
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" class="evaluation-header">
        {{ SeenInfo.MovieInfo.MovieName }}
      </h1>
      <template #right>
        <span class="update-evaluation-button">发布</span>
      </template>
    </van-nav-bar>
    <div style="margin-top: 46px;">
      <div class="evaluation-score">
        <div style="text-align: center;height: 80px;">
          <div style="padding-top: 30px;">
            <span>我的评分</span>&nbsp;<span style="color: #ffd21e;">{{ SeenInfo.Score }}分</span>
          </div>
        </div>
        <div class="evaluation-score-stars">
          <van-rate v-model="stars" allow-half size="40" gutter="25" void-icon="star" void-color="#eee" />
          <div style="width: 300px;margin: 0 auto;">
            <span
              :style="{
                color: SeenInfo.Score == 1 || SeenInfo.Score == 2 ? '#ffd21e' : '#808080',
              }"
              >极差</span
            >
            <span
              class="show-text-span"
              :style="{
                color: SeenInfo.Score == 3 || SeenInfo.Score == 4 ? '#ffd21e' : '#808080',
              }"
              >较差</span
            >
            <span
              :style="{
                color: SeenInfo.Score == 5 || SeenInfo.Score == 6 ? '#ffd21e' : '#808080',
              }"
              class="show-text-span"
              >一般</span
            >
            <span
              :style="{
                color: SeenInfo.Score == 7 || SeenInfo.Score == 8 ? '#ffd21e' : '#808080',
              }"
              class="show-text-span"
              >不错</span
            >
            <span
              :style="{
                color: SeenInfo.Score == 9 || SeenInfo.Score == 10 ? '#ffd21e' : '#808080',
              }"
              class="show-text-span"
              >很棒</span
            >
          </div>
        </div>
      </div>
      <van-divider />
      <div class="evaluation-review">
        <van-field
          v-model="SeenInfo.Review"
          rows="5"
          autosize
          type="textarea"
          placeholder="写几句评价"
          maxlength="1000"
          show-word-limit
        />
      </div>
    </div>
    <van-icon @click="clickDelete" name="delete" color="#1989fa" size="35" style="float: right;margin-right: 20px;" />
  </div>
</template>

<script>
import { reqUpdateEvaluation, reqDeleteEvaluation } from '../api'
export default {
  name: 'Evaluation',
  data() {
    //这里存放数据
    return {
      oldScore: '',
      oldReview: '',
    }
  },
  computed: {
    stars: {
      get() {
        return this.SeenInfo.Score / 2
      },
      set(val) {
        this.SeenInfo.Score = val * 2
      },
    },
    SeenInfo() {
      return this.$store.state.seenInfo
    },
  },
  watch: {
    $route() {
      this.oldReview = this.SeenInfo.Review
      this.oldScore = this.SeenInfo.Score
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.SeenInfo.oldReview = this.oldReview
      this.SeenInfo.Score = this.oldScore
      this.$router.replace(this.$store.state.backPath)
    },
    clickDelete() {
      this.$dialog
        .confirm({
          title: '确认删除评分和评论吗?',
        })
        .then(() => {
          this.deleteEvaluation()
        })
        .catch(() => {})
    },
    //发布评价
    async updateEvaluation() {
      let result
      const MovieID = this.SeenInfo.MovieInfo.MovieID
      const Score = this.SeenInfo.Score
      if (Score == 0) {
        this.$toast('请选择一个分值')
        return
      }
      const Review = this.SeenInfo.Review
      result = await reqUpdateEvaluation({ MovieID, Score, Review })
      if (result.code == 0) {
        this.$router.replace(this.$store.state.backPath)
        this.$toast('发布成功')
      } else {
        this.$toast(result.message)
      }
    },
    //删除评价
    async deleteEvaluation() {
      let result
      const MovieID = this.SeenInfo.MovieInfo.MovieID
      result = await reqDeleteEvaluation({ MovieID })
      if (result.code == 0) {
        this.SeenInfo.Score = 0
        this.SeenInfo.Review = ''
        this.$router.replace(this.$store.state.backPath)
        this.$toast('删除成功')
      } else {
        this.$toast(result.message)
      }
    },
  },
  created() {
    this.oldReview = this.SeenInfo.Review
    this.oldScore = this.SeenInfo.Score
  },
}
</script>
<style>
@import url('../assets/css/evaluation.css');
</style>
