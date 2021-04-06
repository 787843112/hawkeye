<!-- 影厅页面 -->
<template>
  <div id="screening">
    <div class="screening-div">
      <!-- 顶部筛选 -->
      <div class="screening-header">
        <el-input
          v-model="this.$store.state.cinemaInfo.CinemaID"
          disabled
          style="width: 150px;margin-left: 5px;margin-top: 20px;"
        >
          <template slot="prepend">影院ID</template>
        </el-input>
        <el-input v-model="HallInfo.HallID" disabled style="width: 150px;margin-left: 5px;margin-top: 20px;">
          <template slot="prepend">影厅ID</template>
        </el-input>
        <el-input v-model="HallInfo.HallName" disabled style="width: 300px;margin-left: 5px;margin-top: 20px;">
          <template slot="prepend">影厅名称</template>
        </el-input>
        <!-- 日期选择 -->
        <el-date-picker
          :editable="false"
          style="margin-left: 5px;margin-top: 20px"
          v-model="dateRange"
          type="daterange"
          range-separator="至"
          format="yyyy 年 MM 月 dd 日"
          value-format="yyyy-MM-dd"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        >
        </el-date-picker>
        <el-button
          @click="openAddScreening"
          style="margin-left: 5px;margin-top: 20px"
          type="primary"
          icon="el-icon-plus"
          >添加场次</el-button
        >
        <el-button
          @click="toDeleteSomeScreening"
          style="margin-top: 20px;margin-left: 5px;margin-top: 20px"
          type="danger"
          :disabled="this.ScreeningIDs.length == 0"
          icon="el-icon-delete"
          >批量删除</el-button
        >
      </div>
      <!-- 添加场次 -->
      <el-dialog title="添加场次" :visible.sync="showAddScreening" class="add-screening-dialog">
        <el-input v-model="HallInfo.HallID" disabled style="width: 150px;">
          <template slot="prepend">影厅ID</template>
        </el-input>
        <!-- 选择电影 -->
        <span style="font-size: 20px;margin-left: 20px;">电影:&nbsp;</span>
        <el-select
          v-model="addScreening.MovieID"
          filterable
          @change="changeMovie"
          placeholder="请选择放映的电影"
          style="width: 339px;"
        >
          <el-option
            v-for="item in ReleasedMovieList"
            :key="item.MovieID"
            :label="item.MovieName"
            :value="item.MovieID"
          >
          </el-option>
        </el-select>
        <!-- 日期 -->
        <span style="font-size: 20px;">日期:&nbsp;</span>
        <el-date-picker
          :editable="false"
          :default-value="new Date()"
          style="margin-top: 20px;width: 190px;"
          v-model="addScreening.Date"
          type="date"
          placeholder="选择日期"
          format="yyyy 年 MM 月 dd 日"
          value-format="yyyy-MM-dd"
        >
        </el-date-picker>
        <!-- 开始时间 -->
        <span style="font-size: 20px;margin-left: 15px;">时间:&nbsp;</span>
        <el-time-picker
          :editable="false"
          :default-value="new Date()"
          style="width: 120px;"
          v-model="addScreening.StartTime"
          placeholder="开始时间"
          format="HH:mm"
          value-format="HH:mm:00"
        >
        </el-time-picker>
        <!-- 结束时间 -->
        <span style="font-size: 20px">-</span>
        <el-time-picker
          :editable="false"
          disabled
          readonly
          style="width: 120px;"
          v-model="this.EndTime"
          placeholder="结束时间"
          format="HH:mm"
          value-format="HH:mm:00"
        >
        </el-time-picker>
        <span style="font-size: 20px;margin-left: 130px">价格:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 200px"
          :min="1"
          :max="127"
          placeholder="请输入价格"
          v-model="addScreening.Price"
        ></el-input-number>
        <span style="font-size: 20px;margin-left: 10px">元</span>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showAddScreening = false">取 消</el-button>
          <el-button :disabled="addScreeningDisabled" type="primary" @click="toAddScreening" icon="el-icon-finished"
            >确认添加</el-button
          >
        </div>
      </el-dialog>
      <!-- 场次列表 -->
      <div class="screening-body">
        <div class="screening-list">
          <el-table
            max-height="670px"
            ref="multipleTable"
            stripe
            @selection-change="handleSelectionChange"
            :data="ScreeningList"
            tooltip-effect="dark"
            class="screeningtable"
          >
            <el-table-column align="center" type="selection" width="50"></el-table-column>
            <el-table-column align="center" prop="ScreeningID" label="ID" width="80"></el-table-column>
            <el-table-column
              align="center"
              prop="MovieName"
              label="放映电影"
              :formatter="formatterMovieName"
              width="300"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              align="center"
              prop="StartTime"
              label="日期"
              :formatter="formatterDate"
              width="100"
            ></el-table-column>
            <el-table-column
              align="center"
              prop="StartTime"
              label="开始时间"
              :formatter="formatterStartTime"
              width="100"
            ></el-table-column>
            <el-table-column
              align="center"
              prop="EndTime"
              label="结束时间"
              :formatter="formatterEndTime"
              width="100"
            ></el-table-column>
            <el-table-column align="center" prop="Price" label="价格" width="80"></el-table-column>
            <el-table-column label="操作" align="center" width="100">
              <template slot-scope="scope">
                <el-button
                  size="mini"
                  icon="el-icon-delete-solid"
                  type="danger"
                  @click="toDeleteOneScreening(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
            <el-table-column></el-table-column>
          </el-table>
        </div>
        <div class="screening-footer">
          <el-pagination
            background
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :page-sizes="[5, 10, 20, 30, 50, 100]"
            :page-size="this.Size"
            :current-page="this.Page"
            layout="total,sizes,prev,pager,next,jumper"
            :total="count"
          >
          </el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reqReleasedMovies, reqAddScreening, reqGetScreenings, reqDeleteScreenings } from '../api'
export default {
  name: 'Screening',
  data() {
    //这里存放数据
    return {
      StartDate: '',
      EndDate: '',
      Page: 1,
      Size: 10,
      count: 0,
      dateRange: [],
      MovieLength: '',
      addScreening: {
        MovieID: '',
        Date: '',
        StartTime: '',
        Price: '',
      },
      ScreeningIDs: [],
      showAddScreening: false,
      ScreeningList: [],
      ReleasedMovieList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    HallInfo() {
      return this.$store.state.hallInfo
    },
    EndTime() {
      var date = new Date(
        new Date(this.addScreening.Date + ' ' + this.addScreening.StartTime).getTime() + 1000 * 60 * this.MovieLength
      )
      var YY = date.getFullYear() + '-'
      var MM = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
      var DD = date.getDate() < 10 ? '0' + date.getDate() : date.getDate()
      var hh = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':'
      var mm = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':'
      var ss = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()
      return YY + MM + DD + ' ' + hh + mm + ss
    },
    addScreeningDisabled() {
      return (
        this.addScreening.MovieID == '' ||
        this.addScreening.Date == '' ||
        this.addScreening.StartTime == '' ||
        this.EndTime == '' ||
        this.addScreening.Price == null ||
        this.addScreening.Price == ''
      )
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/screening') {
        this.getReleasedMovieList()
        this.getScreeningList()
      }
    },
    dateRange() {
      if (this.dateRange !== null) {
        this.StartDate = this.dateRange[0]
        this.EndDate = this.dateRange[1]
      } else {
        this.StartDate = ''
        this.EndDate = ''
      }
      this.getScreeningList()
    },
    Page() {
      this.getScreeningList()
    },
    Size() {
      this.getScreeningList()
    },
  },
  //方法集合
  methods: {
    //格式化表格
    formatterMovieName(row) {
      // setTimeout(() => {
      return this.ReleasedMovieList.find((item) => {
        return row.MovieID == item.MovieID
      }).MovieName
      // }, 500)
    },
    formatterDate(row) {
      return row.StartTime.slice(0, 10)
    },
    formatterStartTime(row) {
      return row.StartTime.slice(11, 16)
    },
    formatterEndTime(row) {
      return row.EndTime.slice(11, 16)
    },
    //改变选择的电影时
    changeMovie(movieID) {
      this.ReleasedMovieList.some((item) => {
        if (item.MovieID == movieID) {
          this.MovieLength = item.MovieLength
        }
      })
    },
    //改变每页显示条数
    handleSizeChange(val) {
      this.Size = val
    },
    //改变当前页
    handleCurrentChange(val) {
      this.Page = val
    },
    //多选发生改变时
    handleSelectionChange(val) {
      this.ScreeningIDs = val.map((item) => {
        return item.ScreeningID
      })
    },
    //清空选择
    toggleSelection(rows) {
      if (rows) {
        rows.forEach((row) => {
          this.$refs.multipleTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleTable.clearSelection()
      }
    },
    openAddScreening() {
      this.showAddScreening = true
    },
    //获取正在上映的电影列表
    async getReleasedMovieList() {
      let result = await reqReleasedMovies()
      if (result.code == 0) {
        this.ReleasedMovieList = result.data
      } else {
        this.ReleasedMovieList = []
      }
    },
    //获取场次列表
    async getScreeningList() {
      let result
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const HallID = this.HallInfo.HallID
      const { StartDate, EndDate, Page, Size } = this
      result = await reqGetScreenings({
        CinemaID,
        HallID,
        StartDate,
        EndDate,
        Page,
        Size,
      })
      if (result.code == 0) {
        this.count = result.count
        this.ScreeningList = result.data
      } else {
        this.count = 0
        this.Page = 1
        this.ScreeningList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //添加场次
    async toAddScreening() {
      let result
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const HallID = this.HallInfo.HallID
      const MovieID = this.addScreening.MovieID
      const StartTime = this.addScreening.Date + ' ' + this.addScreening.StartTime
      const EndTime = this.EndTime
      const Price = this.addScreening.Price
      result = await reqAddScreening({
        CinemaID,
        HallID,
        MovieID,
        StartTime,
        EndTime,
        Price,
      })
      if (result.code == 0) {
        this.showAddScreening = false
        this.ScreeningList.unshift(result.data)
        this.count++
        this.$message({
          showClose: true, //可关闭
          message: '添加成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        //this.showAddScreening = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //删除某个场次
    toDeleteOneScreening(screening) {
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const { ScreeningIDs } = this
      ScreeningIDs[0] = screening.ScreeningID
      this.$confirm('此操作将永久删除该场次, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteScreenings({ CinemaID, ScreeningIDs })
          if (result.code == 0) {
            this.getScreeningList()
            this.ScreeningIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.ScreeningIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'error',
              message: result.message,
              center: true, //文字居中
            })
          }
        })
        .catch(() => {
          this.toggleSelection()
          this.ScreeningIDs = []
        })
    },
    //批量删除某些影厅
    toDeleteSomeScreening() {
      const CinemaID = this.$store.state.cinemaInfo.CinemaID
      const { ScreeningIDs } = this
      this.$confirm('此操作将永久删除已选场次, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteScreenings({ CinemaID, ScreeningIDs })
          if (result.code == 0) {
            this.ScreeningIDs = []
            this.getScreeningList()
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.ScreeningIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'error',
              message: result.message,
              center: true, //文字居中
            })
          }
        })
        .catch(() => {
          this.toggleSelection()
          this.ScreeningIDs = []
        })
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getReleasedMovieList()
    this.getScreeningList()
  },
}
</script>
<style>
@import url('../assets/css/screening.css');
</style>
