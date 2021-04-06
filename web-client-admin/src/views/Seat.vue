<!-- 座位信息页面 -->
<template>
  <div id="seat">
    <div class="seat-div">
      <div class="seat-header">
        <el-input
          v-model="this.$store.state.cinemaInfo.CinemaID"
          disabled
          style="margin-top: 20px;width: 150px;margin-left: 10px;"
        >
          <template slot="prepend">影院ID</template>
        </el-input>
        <el-input v-model="HallInfo.HallID" disabled style="margin-top: 20px;width: 150px;margin-left: 10px;">
          <template slot="prepend">影厅ID</template>
        </el-input>
        <el-input
          v-model="HallInfo.HallName"
          maxlength="20"
          show-word-limit
          clearable
          style="width: 300px;margin-left: 10px;"
        >
          <template slot="prepend">影厅名称</template>
        </el-input>
        <span style="font-size: 20px;margin-left: 10px">行数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 150px"
          :min="1"
          :max="127"
          placeholder="座位行数"
          @change="changeRow"
          v-model="HallInfo.HallRow"
        ></el-input-number>
        <!-- 列数 -->
        <span style="font-size: 20px;margin-left: 10px;">列数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 150px"
          :min="1"
          :max="127"
          placeholder="座位列数"
          @change="changeColumn"
          v-model="HallInfo.HallColumn"
        ></el-input-number>
        <el-button
          @click="saveChanges"
          style="margin-left: 15px;"
          type="primary"
          icon="el-icon-finished"
          :disabled="saveDisabled"
          >保存修改</el-button
        >
      </div>
      <div class="screen-center" :style="{ minWidth: HallInfo.HallColumn * 58 + 'px' }">
        <div class="mid-line"></div>
      </div>
      <!-- 座位布局 -->
      <div class="seat-body" :style="{ minWidth: HallInfo.HallColumn * 58 + 'px' }">
        <div class="illustration">
          <div class="illustration-img-wrapper white-seat"></div>
          <span class="illustration-text">座位</span>
          <div class="illustration-img-wrapper red-seat" style="margin-left: 20%;"></div>
          <span class="illustration-text">不设座位</span>
        </div>
        <div class="seat-scrn">
          <img src="../assets/images/scrn.jpg" />
        </div>
        <div class="seat-grid">
          <div class="seat-rowNumber">
            <tr v-for="index in HallInfo.HallRow" :key="index">
              {{
                index
              }}
            </tr>
          </div>

          <div class="inner-seat-wrapper">
            <div v-for="row in HallInfo.HallRow" :key="row">
              <div v-for="col in HallInfo.HallColumn" :key="col" class="seat">
                <div
                  @click="handleChooseSeat(row, col)"
                  :class="changeTheClass(row, col) ? 'white-seat' : 'red-seat'"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reqGetHall, reqUpdateHall, reqGetSeats, reqAddSeats, reqDeleteSeats } from '../api'
export default {
  name: 'Seat',
  data() {
    //这里存放数据
    return {
      Seats: [],
      addSeats: [],
      deleteSeats: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    HallInfo() {
      return this.$store.state.hallInfo
    },
    saveDisabled() {
      return this.HallInfo.HallName == ''
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/seat') {
        this.getHallInfo()
        this.getSeatInfo()
      }
    },
  },
  //方法集合
  methods: {
    //保存修改
    saveChanges() {
      this.toUpdateHall()
      setTimeout(() => {
        if (this.addSeats.length !== 0) {
          this.toAddSeats()
        }
        if (this.deleteSeats.length !== 0) {
          this.toDeleteSeats()
        }
      }, 500)
      setTimeout(() => {
        this.getSeatInfo()
        this.addSeats = []
        this.deleteSeats = []
      }, 1000)
    },
    //改变行数时
    changeRow(newRow, oldRow) {
      if (newRow > oldRow) {
        for (let i = 1; i <= this.HallInfo.HallColumn; i++) {
          this.Seats.unshift({
            HallID: this.HallInfo.HallID,
            SeatRow: newRow,
            SeatColumn: i,
          })
        }
      } else {
        let Seats = this.Seats.filter((item) => {
          return item.SeatRow < oldRow
        })
        this.Seats = Seats
      }
    },
    //列数改变时
    changeColumn(newCol, oldCol) {
      if (newCol > oldCol) {
        for (let i = 1; i <= this.HallInfo.HallRow; i++) {
          this.Seats.unshift({
            HallID: this.HallInfo.HallID,
            SeatRow: i,
            SeatColumn: newCol,
          })
        }
      } else {
        let Seats = this.Seats.filter((item) => {
          return item.SeatColumn < oldCol
        })
        this.Seats = Seats
      }
    },
    //点击座位时的处理
    handleChooseSeat(row, col) {
      if (this.changeTheClass(row, col)) {
        this.addSeats.map((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.addSeats.splice(i, 1)
          }
        })
        this.deleteSeats.map((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.deleteSeats.splice(i, 1)
          }
        })
        this.deleteSeats.unshift({
          HallID: this.HallInfo.HallID,
          SeatRow: row,
          SeatColumn: col,
        })
        this.Seats.some((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.Seats.splice(i, 1)
          }
        })
      } else {
        this.addSeats.map((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.addSeats.splice(i, 1)
          }
        })
        this.deleteSeats.map((item, i) => {
          if (item.SeatRow == row && item.SeatColumn == col) {
            this.deleteSeats.splice(i, 1)
          }
        })
        this.addSeats.unshift({
          HallID: this.HallInfo.HallID,
          SeatRow: row,
          SeatColumn: col,
        })
        this.Seats.unshift({
          HallID: this.HallInfo.HallID,
          SeatRow: row,
          SeatColumn: col,
        })
      }
    },
    //改变样式
    changeTheClass(row, col) {
      return this.Seats.some((item) => {
        return item.SeatRow == row && item.SeatColumn == col
      })
    },
    //获取当前影厅的信息
    async getHallInfo() {
      let result
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const HallID = this.HallInfo.HallID
      result = await reqGetHall({ CinemaID, HallID })
      if (result.code == 0) {
        // this.HallInfo = result.data
        this.$store.dispatch('recordHallInfo', result.data)
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //获取座位信息
    async getSeatInfo() {
      let result
      const HallID = this.HallInfo.HallID
      const CinemaID = this.HallInfo.CinemaID
      result = await reqGetSeats({ HallID, CinemaID })
      if (result.code == 0) {
        this.Seats = result.data
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //修改影厅
    async toUpdateHall() {
      let result
      const HallID = this.HallInfo.HallID
      const CinemaID = this.HallInfo.CinemaID
      const HallName = this.HallInfo.HallName
      const HallRow = this.HallInfo.HallRow
      const HallColumn = this.HallInfo.HallColumn
      result = await reqUpdateHall({
        HallID,
        CinemaID,
        HallName,
        HallRow,
        HallColumn,
      })
      if (result.code == 0) {
        this.$message({
          showClose: true, //可关闭
          message: '修改影厅成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //添加某些座位
    async toAddSeats() {
      let result
      const CinemaID = this.HallInfo.CinemaID
      const Seats = this.addSeats
      result = await reqAddSeats({ CinemaID, Seats })
      if (result.code == 0) {
        this.$message({
          showClose: true, //可关闭
          message: '添加座位成功!',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //删除某些座位
    async toDeleteSeats() {
      let result
      const CinemaID = this.HallInfo.CinemaID
      const Seats = this.deleteSeats
      result = await reqDeleteSeats({ CinemaID, Seats })
      if (result.code == 0) {
        this.$message({
          showClose: true, //可关闭
          message: '删除座位成功!',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getSeatInfo()
  },
}
</script>
<style>
@import url('../assets/css/seat.css');
</style>
