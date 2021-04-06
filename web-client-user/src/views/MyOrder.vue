<!-- 电影票页面 -->
<template>
  <div id="myOrder">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
        电&nbsp;影&nbsp;票
      </h1>
    </van-nav-bar>
    <van-sticky :offset-top="46">
      <van-notice-bar color="#1989fa" background="#ecf9ff" left-icon="coupon-o" :scrollable="false"
        >总共&nbsp;<span style="color: crimson;">{{ count }}</span
        >&nbsp;张电影票</van-notice-bar
      >
    </van-sticky>
    <!-- 电影票列表 -->
    <div class="order-list-div">
      <van-pull-refresh v-model="orderIsLoading" @refresh="orderOnRefresh">
        <van-list
          :offset="5"
          v-model="orderLoading"
          finished-text="没有更多了o(ﾟДﾟ)っ!"
          :immediate-check="false"
          :finished="orderFinished"
          @load="orderOnLoad"
          :error.sync="orderError"
          error-text="请求失败，点击重新加载"
        >
          <van-swipe-cell
            :disabled="!isPast(item.Date + ' ' + item.StartTime)"
            right-width="50"
            style="height: 150px;margin-top: 10px;"
            v-for="(item, index) in OrderList"
            :key="item.OrderID"
          >
            <div class="order-cell-left">
              <div>
                <span style="color: #fff;font-size: 20px;">座</span>
              </div>
            </div>
            <div class="order-cell-mid"></div>
            <!-- 电影票信息div -->
            <div class="order-info-div">
              <div @click="toOrderDetail(item)">
                <!-- 电影名称和张数div -->
                <div class="order-info-movie-name">
                  <span
                    :style="{
                      fontSize: 20 + 'px',
                      color: isPast(item.Date + ' ' + item.StartTime) ? '#c4c4c4' : '#000',
                    }"
                    >{{ item.MovieName }}&nbsp;{{ item.Number }}张</span
                  >
                </div>
                <!-- 观影日期和时间div -->
                <div class="order-info-date">
                  <span
                    :style="{
                      color: isPast(item.Date + ' ' + item.StartTime) ? '#c4c4c4' : '#000',
                    }"
                    >{{ item.Date }}&nbsp;{{ item.StartTime }}</span
                  >
                </div>
                <!-- 电影院名称div -->
                <div class="order-info-cinema-name">
                  <span
                    :style="{
                      color: isPast(item.Date + ' ' + item.StartTime) ? '#c4c4c4' : '#000',
                    }"
                    >{{ item.CinemaName }}</span
                  >
                </div>
              </div>
              <!-- 查看评价 -->
              <van-button class="order-button" plain round @click="toEvaluation(item)">查看影片评价</van-button>
            </div>
            <template #right>
              <van-button @click="clickDelete(item, index)" square text="删除" type="danger" style="height: 100%" />
            </template>
          </van-swipe-cell>
        </van-list>
      </van-pull-refresh>
    </div>
  </div>
</template>

<script>
import { reqGetOrders, reqDeleteOrder, reqGetMovieInfo, reqGetSeenByMovieID } from '../api'
export default {
  name: 'MyOrder',
  data() {
    //这里存放数据
    return {
      Page: 1,
      Size: 5,
      OrderList: [],
      count: 0,
      orderLoading: false,
      orderFinished: false,
      orderIsLoading: false,
      orderError: false,
    }
  },
  //监听属性 类似于data概念
  computed: {},
  //监控data中的数据变化
  watch: {
    $route(to, from) {
      if (from.path == '/orderDetail') {
        if (this.$route.params.orderID != null) {
          for (let i = 0; i < this.OrderList.length; i++) {
            if (this.OrderList[i].OrderID == this.$route.params.orderID) {
              this.OrderList.splice(i, 1)
              if (this.OrderList.length < this.Size) {
                this.Page = 1
                this.Size = 5
                this.getOrderList()
              }
            }
          }
        }
      }
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/myCenter')
    },
    async toOrderDetail(order) {
      await this.getMovieInfo(order.MovieID)
      await this.getSeenInfo(order.MovieID)
      await this.$store.dispatch('recordOrderInfo', order)
      this.$router.replace('/orderDetail')
    },
    async toEvaluation(order) {
      await this.getSeenInfo(order.MovieID)
      this.$router.replace('/evaluation')
    },
    isPast(date) {
      return new Date().getTime() > new Date(date).getTime()
    },
    clickDelete(order, index) {
      this.$dialog
        .confirm({
          title: '删除后无法恢复,确认吗?',
        })
        .then(() => {
          this.delectOrder(order.OrderID, index)
        })
        .catch(() => {})
    },
    //加载更多电影票
    orderOnLoad() {
      this.Page++
      this.orderError = false
      this.getOrderList()
    },
    //下拉刷新电影票列表
    orderOnRefresh() {
      this.Page = 1
      this.orderIsLoading = true
      this.orderFinished = false
      this.getOrderList()
    },
    //获取电影票
    async getOrderList() {
      let result
      const { Page, Size } = this
      result = await reqGetOrders({ Page, Size })
      if (result.code == 0) {
        if (Page == 1) {
          this.OrderList = result.data
          this.count = result.count
        } else {
          this.OrderList = this.OrderList.concat(result.data)
        }
        this.orderLoading = false
        this.orderIsLoading = false
      } else if (result.code == -1) {
        this.orderLoading = false
        this.orderIsLoading = false
        this.orderFinished = true
      } else {
        // this.$toast(result.message)
        console.log(result.message)
        this.orderError = true
        this.orderIsLoading = false
        this.orderLoading = false
      }
    },
    //根据电影ID获取看过的信息
    async getSeenInfo(MovieID) {
      let result = await reqGetSeenByMovieID({ MovieID })
      if (result.code == 0) {
        this.$store.dispatch('recordSeenInfo', result.data)
      } else {
        console.log(result.message)
      }
    },
    //根据电影ID获取电影信息
    async getMovieInfo(MovieID) {
      let result = await reqGetMovieInfo({ MovieID })
      if (result.code == 0) {
        this.$store.dispatch('recordMovieInfo', result.data)
      } else {
        console.log(result.message)
      }
    },
    //删除某张电影票(订单)
    async delectOrder(OrderID, index) {
      let result = await reqDeleteOrder({ OrderID })
      if (result.code == 0) {
        this.$toast('删除成功!')
        this.OrderList.splice(index, 1)
        if (this.OrderList.length < this.Size) {
          this.Page = 1
          this.Size = 5
          this.getOrderList()
        }
      } else {
        this.$toast(result.message)
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getOrderList()
  },
}
</script>
<style>
@import url('../assets/css/myOrder.css');
</style>
