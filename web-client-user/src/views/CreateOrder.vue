<!-- 确认创建订单页面 -->
<template>
  <div id="createOrder">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
        支 付 订 单
      </h1>
    </van-nav-bar>
    <div style="margin-top: 46px">
      <van-notice-bar left-icon="volume-o">
        <span>请在</span>
        <span v-show="minutes != 0" style="color: red">{{ minutes }}</span>
        <span v-show="minutes != 0">分</span>
        <span style="color: red">{{ seconds }}</span>
        <span>秒内完成支付!</span>
      </van-notice-bar>
    </div>
    <div>
      <van-progress
        :color="payTime > 400 ? '#00bb00' : 200 > payTime ? '#ff0000' : '#f2826a'"
        :show-pivot="false"
        :percentage="payTime / 6"
        stroke-width="8"
      />
    </div>
    <van-divider></van-divider>
    <p class="movie-info-p">
      {{ MovieInfo.MovieName }}
    </p>
    <van-divider></van-divider>
    <p class="select-time-p">{{ SelectDate }}&nbsp;{{ Screening.StartTime.slice(11, 16) }}</p>
    <p class="cinema-info-p">{{ CinemaInfo.CinemaName }}</p>
    <p class="cinema-info-p">{{ HallInfo.HallName }}</p>
    <van-divider></van-divider>
    <div style="text-align: center">
      <van-tag class="seat-info-tab" type="warning" text-color="#000" v-for="(item, index) in SelectSeat" :key="index"
        >&nbsp;{{ item.SeatRow + '排' + item.SeatColumn + '座' }}&nbsp;</van-tag
      >
    </div>
    <van-divider></van-divider>
    <van-cell-group>
      <van-cell icon="user-circle-o" title="用户">{{ UserInfo.UserName }} </van-cell>
      <van-cell icon="phone-o" title="手机号码">{{ UserInfo.Phone }} </van-cell>
      <van-cell icon="gold-coin-o" title="小计">¥{{ SelectSeat.length * Screening.Price }} </van-cell>
    </van-cell-group>
    <van-divider></van-divider>
    <van-button
      loading-text="支付中..."
      class="create-order-button"
      loading-type="spinner"
      type="primary"
      block
      round
      icon="alipay"
      :loading="payLoad"
      @click="toPay"
    >
      <span> ¥{{ SelectSeat.length * Screening.Price }}&nbsp;确认支付</span>
    </van-button>
  </div>
</template>

<script>
import { reqUpdateStatus, reqAddOrder } from '../api'
export default {
  name: 'CreateOrder',
  data() {
    return {
      intervalId: null,
      payLoad: false,
      payTime: 0,
      minutes: 0,
      seconds: 0,
      flag: 1,
    }
  },
  computed: {
    UserInfo() {
      return this.$store.state.userInfo
    },
    CinemaInfo() {
      return this.$store.state.cinemaInfo
    },
    MovieInfo() {
      return this.$store.state.movieInfo
    },
    Screening() {
      return this.$store.state.screening
    },
    HallInfo() {
      return this.$store.state.hallInfo
    },
    SelectDate() {
      return this.$store.state.selectDate
    },
    SelectSeat() {
      return this.$store.state.selectSeat
    },
  },
  //监控data中的数据变化
  watch: {
    $route() {
      this.countDown()
      clearInterval(this.intervalId)
      this.intervalId = null
    },
    payTime() {
      if (this.payTime == 0) {
        this.$toast('已超时')
        setTimeout(() => {
          this.intervalId = null
          this.$router.replace('/movie')
        }, 1000)
      }
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.intervalId = null
      this.$router.replace('/selectSeat')
    },
    //倒计时
    countDown() {
      this.payTime = 600
      if (this.intervalId != null) return
      this.intervalId = setInterval(() => {
        this.payTime--
        this.minutes = Math.floor(this.payTime / 60)
        this.seconds = this.payTime % 60
        if (this.payTime <= 0) {
          clearInterval(this.intervalId)
          this.intervalId = null
        }
      }, 1000)
    },
    async toPay() {
      this.payLoad = true
      await this.createOrder()
      setTimeout(() => {
        // 加载状态结束
        if (this.flag == 0) {
          this.$store.dispatch('recordSelectSeat', [])
          this.$store.dispatch('recordSelectScreening', {})
          this.$store.dispatch('recordHallInfo', {})
          this.$store.dispatch('recordSeatStatus', {})
          this.$toast('支付成功!')
          this.payLoad = false
          this.$router.replace('/movie')
        } else {
          this.payLoad = false
          this.$toast('支付失败!')
        }
      }, 700)
    },
    //创建订单
    async createOrder() {
      let result
      const Status = this.SelectSeat
      result = await reqUpdateStatus({ Status })
      if (result.code == 0) {
        let result1
        const MovieID = this.MovieInfo.MovieID
        const MovieName = this.MovieInfo.MovieName
        const Date = this.SelectDate
        const StartTime = this.Screening.StartTime.slice(11, 16)
        const EndTime = this.Screening.EndTime.slice(11, 16)
        const CinemaName = this.CinemaInfo.CinemaName
        const Address = this.CinemaInfo.Address
        const HallName = this.HallInfo.HallName
        const Number = this.SelectSeat.length
        const Price = this.SelectSeat.length * this.Screening.Price
        result1 = await reqAddOrder({
          MovieID,
          MovieName,
          Date,
          StartTime,
          EndTime,
          CinemaName,
          Address,
          HallName,
          Number,
          Price,
        })
        if (result1.code == 0) {
          this.flag = 0
        } else {
          this.flag = 1
          console.log(result1.message)
        }
      } else {
        this.flag = 1
        console.log(result.message)
      }
    },
  },
  created() {
    this.countDown()
  },
}
</script>
<style>
@import url('../assets/css/createOrder.css');
</style>
