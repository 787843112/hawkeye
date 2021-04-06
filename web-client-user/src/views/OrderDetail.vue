<!-- 订单详情页面 -->
<template>
  <div id="orderDetail">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" class="order-detail-header">
        电&nbsp;影&nbsp;票&nbsp;详&nbsp;细
      </h1>
    </van-nav-bar>
    <div class="order-detail-body">
      <div class="order-movie-status">
        <b>{{ isPast(OrderInfo.Date + ' ' + OrderInfo.StartTime) ? '电影已放映' : '祝您观影愉快' }}</b>
      </div>
      <div class="order-detail-button">
        <van-button @click="toEvaluation" style="height: 35px;width: 150px;" round color="#1ccece"
          >查看影片评价</van-button
        >
      </div>
      <div class="order-detail-info">
        <div class="order-movie-info">
          <div class="order-movie-info-rigth">
            <p style="font-size: 25px;font-family: cursive">
              {{ MovieInfo.MovieName }}
            </p>
            <p style="margin-top: 10px;">
              {{ OrderInfo.Date }}&nbsp;{{ OrderInfo.StartTime }}&nbsp;~&nbsp;{{ OrderInfo.EndTime }}
            </p>
            <p>{{ OrderInfo.CinemaName }}&nbsp;{{ OrderInfo.HallName }}</p>
            <p style="font-size: 14px;color: grey;">{{ OrderInfo.Seat }}</p>
          </div>
          <div class="order-movie-info-img">
            <img :src="'static/posters/' + MovieInfo.Poster" />
          </div>
        </div>
        <div style="clear: both"></div>
        <van-divider dashed />
        <div class="cinema-info">
          <b>{{ OrderInfo.CinemaName }}</b>
          <p>{{ OrderInfo.Address }}</p>
        </div>
        <van-divider dashed />
        <div class="order-info">
          <b>实付金额:&nbsp;￥{{ OrderInfo.Price }}</b>
          <p style="margin-top: 10px;">订单号:&nbsp;{{ OrderInfo.OrderID }}</p>
          <p>
            购买时间:&nbsp;{{ OrderInfo.CreateTime.slice(0, 10) }}&nbsp;{{ &nbsp;OrderInfo.CreateTime.slice(11,19) }}
          </p>
          <p>手机号:&nbsp;{{ UserInfo.Phone }}</p>
          <van-divider dashed />
        </div>
        <div class="order-notice">
          <b>观影须知</b>
          <p>1.请提前到达影院现场,找到自助取票机,打印纸质电影票,完成取票。</p>
          <p>2.如现场自助取票机无法打印电影票,请联系影院工作人员处理。</p>
          <p>3.凭打印好的纸质电影票,检票入场观影。</p>
          <p>
            4.如有开具所购电影票发票的需求,请保留好电影票根,尽量在观影当天联系影城工作人员进行开具,如遇特殊情况请及时联系淘票票人工客服咨询。
          </p>
          <p>
            5.改签、退票服务由影城决定,特殊场次及使用兑换券的场次不支持改签、退票。
          </p>
        </div>
        <van-divider dashed />
        <div style="height: 65px;">
          <van-button
            :disabled="!isPast(OrderInfo.Date + ' ' + OrderInfo.StartTime)"
            round
            plain
            block
            style="width: 80%;margin-left: 10%;"
            @click="clickDelete"
            >删除订单</van-button
          >
        </div>
      </div>
      <div style="height: 1px;"></div>
    </div>
  </div>
</template>

<script>
import { reqDeleteOrder } from '../api'
export default {
  name: 'OrderDetail',
  //监听属性 类似于data概念
  computed: {
    UserInfo() {
      return this.$store.state.userInfo
    },
    MovieInfo() {
      return this.$store.state.movieInfo
    },
    OrderInfo() {
      return this.$store.state.orderInfo
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/myOrder')
    },
    //去评价页面
    toEvaluation() {
      this.$router.replace('/evaluation')
    },
    isPast(date) {
      return new Date().getTime() > new Date(date).getTime()
    },
    clickDelete() {
      this.$dialog
        .confirm({
          title: '删除后无法恢复,确认吗?',
        })
        .then(() => {
          this.delectOrder()
        })
        .catch(() => {})
    },
    //删除某张电影票(订单)
    async delectOrder() {
      let result
      const OrderID = this.OrderInfo.OrderID
      result = await reqDeleteOrder({ OrderID })
      if (result.code == 0) {
        this.$toast('删除成功!')
        this.$router.push({
          name: 'MyOrder',
          params: { orderID: OrderID },
        })
      } else {
        this.$toast(result.message)
      }
    },
  },
}
</script>
<style>
@import url('../assets/css/orderDetail.css');
</style>
