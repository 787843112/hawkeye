<!-- 选择座位页面 -->
<template>
  <div id="selectSeat" v-if="Screening.ScreeningID != null">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title" style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
        {{ CinemaInfo.CinemaName }}
      </h1>
    </van-nav-bar>
    <div class="seat-movie-info">
      <div class="movie-name">{{ MovieInfo.MovieName }}</div>
      <div class="start-time">
        {{ SelectDate }}&nbsp;{{ Screening.StartTime.slice(11, 16) }}&nbsp;单价: ¥&nbsp;{{ Screening.Price }}
      </div>
    </div>
    <div class="seat-body">
      <div class="seat-scrn" :style="{ width: HallInfo.HallColumn * 35 + 35 + 'px' }">
        <div class="moviehall-name">
          {{ HallInfo.HallName }}
        </div>
      </div>
      <div class="screen-center" :style="{ width: HallInfo.HallColumn * 35 + 35 + 'px' }">
        <div class="mid-line" :style="{ height: HallInfo.HallRow * 35 + 25 + 'px' }"></div>
      </div>
      <div class="seat-grid">
        <div class="seat-rowNumber">
          <tr v-for="index in HallInfo.HallRow" :key="index">
            {{
              index
            }}
          </tr>
        </div>

        <div class="inner-seat-wrapper" :style="{ width: HallInfo.HallColumn * 35 + 35 + 'px' }">
          <div v-for="row in HallInfo.HallRow" :key="row">
            <div v-for="col in HallInfo.HallColumn" :key="col" class="seat">
              <div
                class="inner-seat"
                v-if="isSeat(row, col)"
                @click="handleClickSeat(row, col)"
                :class="
                  changeTheClass(row, col).Status == 1 ? 'red-seat' : isSelected(row, col) ? 'green-seat' : 'white-seat'
                "
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="illustration" :style="{ maxWidth: HallInfo.HallColumn * 35 + 35 + 'px' }">
      <div class="illustration-img-wrapper white-seat"></div>
      <span class="illustration-text">可选</span>
      <div class="illustration-img-wrapper green-seat"></div>
      <span class="illustration-text">已选</span>
      <div class="illustration-img-wrapper red-seat"></div>
      <span class="illustration-text">已售</span>
    </div>
    <van-divider />
    <div v-if="selectSeat.length != 0" class="select-seat-div">
      <p style="font-size: 20px">已选座位</p>
      <van-tag
        class="seat-info-tab"
        v-for="(item, index) in selectSeat"
        color="#f2526a"
        text-color="#000"
        closeable
        :key="index"
        @close="closeSeatTag(index)"
        >{{ item.SeatRow + '排' + item.SeatColumn + '座' }}
      </van-tag>
    </div>
    <div class="seat-button">
      <van-button class="selectSeat-button-1" disabled round v-if="selectSeat.length == 0" block type="warning">
        请先选择座位
      </van-button>
      <van-button class="selectSeat-button-2" v-else round block icon="balance-list" type="info" @click="toCreateOrder">
        <i>¥{{ selectSeat.length * Screening.Price }}</i>
      </van-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SelectSeat',
  data() {
    //这里存放数据
    return {
      selectSeat: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
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
    SeatStatus() {
      return this.$store.state.seatStatus
    },
    SelectDate() {
      return this.$store.state.selectDate
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to, from) {
      if (from.path == '/shows') {
        this.selectSeat = []
      }
    },
  },
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.selectSeat = []
      this.$router.replace('/shows')
    },
    //判断是否有位置
    isSeat(row, col) {
      return this.SeatStatus.some((item) => {
        return item.SeatRow == row && item.SeatColumn == col
      })
    },
    //判断是否已出售
    changeTheClass(row, col) {
      return this.SeatStatus.find((item) => {
        return item.SeatRow == row && item.SeatColumn == col
      })
    },
    //判断是否已选
    isSelected(row, col) {
      return this.selectSeat.some((item) => {
        return item.SeatRow == row && item.SeatColumn == col
      })
    },
    //处理点击某个座位
    handleClickSeat(row, col) {
      let seat = this.SeatStatus.find((item) => {
        return item.SeatRow == row && item.SeatColumn == col
      })
      if (seat.Status == 1) return
      if (this.isSelected(row, col)) {
        this.selectSeat.map((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.selectSeat.splice(i, 1)
          }
        })
      } else if (this.selectSeat.length < 5) {
        this.selectSeat.unshift(seat)
      } else {
        this.$toast('最多选择5个座位!')
      }
    },
    //关闭某个座位Tag
    closeSeatTag(i) {
      this.selectSeat.splice(i, 1)
    },
    //确认座位去支付
    toCreateOrder() {
      this.$store.dispatch('recordSelectSeat', this.selectSeat)
      this.$router.replace('/createOrder')
    },
  },
}
</script>
<style>
@import url('../assets/css/selectSeat.css');
</style>
