<!-- 影厅页面 -->
<template>
  <div id="hall">
    <div class="hall-div">
      <div class="hall-header">
        <el-input v-model="CinemaInfo.CinemaID" disabled style="width: 150px;margin-left: 10px;">
          <template slot="prepend">影院ID</template>
        </el-input>
        <el-input v-model="CinemaInfo.CinemaName" disabled style="width: 300px;margin-left: 20px;">
          <template slot="prepend">影院名称</template>
        </el-input>
        <el-input
          clearable
          placeholder="按影影厅名称搜索"
          prefix-icon="el-icon-search"
          v-model="HallName"
          class="search-hall-name"
        ></el-input>
        <el-button @click="openAddHall" style="margin-left: 3%;" type="primary" icon="el-icon-plus">添加影厅</el-button>
        <el-button
          @click="toDeleteSomeHall"
          style="margin-left: 3%;"
          type="danger"
          :disabled="this.HallIDs.length == 0"
          icon="el-icon-delete"
          >批量删除</el-button
        >
      </div>
      <!-- 添加影厅 -->
      <el-dialog title="添加影厅" :visible.sync="showAddHall" class="add-hall-dialog">
        <el-input v-model="CinemaInfo.CinemaID" disabled>
          <template slot="prepend">影院ID</template>
        </el-input>
        <el-input
          style="margin-top: 20px;"
          maxlength="20"
          show-word-limit
          v-model="addHall.HallName"
          placeholder="请输入影厅名称"
          clearable
        >
          <template slot="prepend">影厅名称</template></el-input
        >
        <!-- 行数 -->
        <span style="font-size: 20px">行数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 170px"
          :min="1"
          :max="127"
          placeholder="座位行数"
          v-model="addHall.HallRow"
        ></el-input-number>
        <!-- 列数 -->
        <span style="font-size: 20px;margin-left: 15px;">列数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 170px"
          :min="1"
          :max="127"
          placeholder="座位列数"
          v-model="addHall.HallColumn"
        ></el-input-number>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showAddHall = false">取 消</el-button>
          <el-button :disabled="addHallDisabled" type="primary" @click="toAddHall" icon="el-icon-finished"
            >确认添加</el-button
          >
        </div>
      </el-dialog>
      <!-- 影厅列表 -->
      <div class="hall-body">
        <div class="hall-list">
          <el-table
            max-height="670px"
            ref="multipleTable"
            stripe
            @selection-change="handleSelectionChange"
            :data="HallList"
            tooltip-effect="dark"
            class="halltable"
          >
            <el-table-column align="center" type="selection" width="50"></el-table-column>
            <el-table-column align="center" prop="HallID" label="ID" width="80"></el-table-column>
            <el-table-column
              align="center"
              prop="HallName"
              label="名称"
              width="250"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              align="center"
              prop="HallRow"
              label="座位行数"
              width="100"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              align="center"
              prop="HallColumn"
              label="座位列数"
              width="100"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column label="操作" align="center" width="400">
              <template slot-scope="scope">
                <el-button type="primary" icon="el-icon-edit" size="mini" @click="openUpdateHall(scope.row)"
                  >修改</el-button
                >
                <el-button type="success" icon="el-icon-s-grid" size="mini" @click="toSeatInfo(scope.row)"
                  >座位信息</el-button
                >
                <el-button type="warning" icon="el-icon-time" size="mini" @click="toScreeningInfo(scope.row)"
                  >场次信息</el-button
                >
                <el-button size="mini" icon="el-icon-delete-solid" type="danger" @click="toDeleteOneHall(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
            <el-table-column></el-table-column>
          </el-table>
        </div>
        <div class="hall-footer">
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
      <!-- 修改影厅 -->
      <el-dialog title="修改影厅" :visible.sync="showUpdateHall" class="add-hall-dialog">
        <el-input v-model="updateHall.HallID" disabled>
          <template slot="prepend">影厅ID</template>
        </el-input>
        <el-input
          style="margin-top: 20px;"
          maxlength="20"
          show-word-limit
          v-model="updateHall.HallName"
          placeholder="请输入影厅名称"
          clearable
        >
          <template slot="prepend">影厅名称</template></el-input
        >
        <!-- 行数 -->
        <span style="font-size: 20px">行数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 170px"
          :min="1"
          :max="127"
          placeholder="座位行数"
          v-model="updateHall.HallRow"
        ></el-input-number>
        <!-- 列数 -->
        <span style="font-size: 20px;margin-left: 15px;">列数:&nbsp;</span>
        <el-input-number
          style="margin-top: 20px;width: 170px"
          :min="1"
          :max="127"
          placeholder="座位列数"
          v-model="updateHall.HallColumn"
        ></el-input-number>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showUpdateHall = false">取 消</el-button>
          <el-button icon="el-icon-finished" :disabled="updateHallDisabled" type="primary" @click="toUpdateHall"
            >确认修改</el-button
          >
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { reqGetHalls, reqDeleteHalls, reqAddHall, reqUpdateHall } from '../api'
export default {
  name: 'Hall',
  data() {
    //这里存放数据
    return {
      HallName: '',
      Page: 1,
      Size: 10,
      count: 0,
      addHall: {
        CinemaID: '',
        HallName: '',
        HallRow: '',
        HallColumn: '',
      },
      updateHall: {
        HallID: '',
        HallName: '',
        HallRow: '',
        HallColumn: '',
      },
      HallIDs: [],
      showAddHall: false,
      showUpdateHall: false,
      HallList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    CinemaInfo() {
      return this.$store.state.cinemaInfo
    },
    addHallDisabled() {
      return this.addHall.HallName == ''
    },
    updateHallDisabled() {
      return this.updateHall.HallName == ''
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/hall') {
        this.getHallList()
      }
    },
    HallName() {
      this.getHallList()
    },
    Page() {
      this.getHallList()
    },
    Size() {
      this.getHallList()
    },
    CinemaInfo() {
      this.getHallList()
    },
  },
  //方法集合
  methods: {
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
      this.HallIDs = val.map((item) => {
        return item.HallID
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
    openAddHall() {
      this.addHall.HallName = ''
      this.addHall.HallRow = 1
      this.addHall.HallColumn = 1
      this.showAddHall = true
    },
    openUpdateHall(hall) {
      this.updateHall.HallID = hall.HallID
      this.updateHall.HallName = hall.HallName
      this.updateHall.HallRow = hall.HallRow
      this.updateHall.HallColumn = hall.HallColumn
      this.showUpdateHall = true
    },
    //去查看座位信息
    toSeatInfo(hall) {
      this.$store.dispatch('recordHallInfo', hall)
      this.$router.replace('/seat')
    },
    //去查看场次信息
    async toScreeningInfo(hall) {
      await this.$store.dispatch('recordHallInfo', hall)
      this.$router.replace('/screening')
    },
    //获取影厅列表
    async getHallList() {
      let result
      const CinemaID = this.CinemaInfo.CinemaID
      const { HallName, Page, Size } = this
      result = await reqGetHalls({ HallName, CinemaID, Page, Size })
      if (result.code == 0) {
        this.count = result.count
        this.HallList = result.data
      } else {
        this.count = 0
        this.Page = 1
        this.HallList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //添加影厅
    async toAddHall() {
      let result
      const HallName = this.addHall.HallName
      const CinemaID = this.CinemaInfo.CinemaID
      const HallRow = this.addHall.HallRow
      const HallColumn = this.addHall.HallColumn
      result = await reqAddHall({ HallName, CinemaID, HallRow, HallColumn })
      if (result.code == 0) {
        this.showAddHall = false
        this.HallList.unshift(result.data)
        this.count++
        this.$message({
          showClose: true, //可关闭
          message: '添加成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.showAddHall = false
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
      const HallID = this.updateHall.HallID
      const CinemaID = this.CinemaInfo.CinemaID
      const HallName = this.updateHall.HallName
      const HallRow = this.updateHall.HallRow
      const HallColumn = this.updateHall.HallColumn
      result = await reqUpdateHall({
        HallID,
        CinemaID,
        HallName,
        HallRow,
        HallColumn,
      })
      if (result.code == 0) {
        this.getHallList()
        this.showUpdateHall = false
        this.$message({
          showClose: true, //可关闭
          message: '修改成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.showUpdateHall = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //删除某个影厅
    toDeleteOneHall(hall) {
      const CinemaID = this.CinemaInfo.CinemaID
      const { HallIDs } = this
      HallIDs[0] = hall.HallID
      this.$confirm('此操作将永久删除该影厅, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteHalls({ CinemaID, HallIDs })
          if (result.code == 0) {
            this.getHallList()
            this.HallIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.HallIDs = []
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
          this.HallIDs = []
        })
    },
    //批量删除某些影厅
    toDeleteSomeHall() {
      const CinemaID = this.CinemaInfo.CinemaID
      const { HallIDs } = this
      this.$confirm('此操作将永久删除已选影厅, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteHalls({ CinemaID, HallIDs })
          if (result.code == 0) {
            this.HallIDs = []
            this.getHallList()
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.HallIDs = []
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
          this.HallIDs = []
        })
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getHallList()
  },
}
</script>
<style>
@import url('../assets/css/hall.css');
</style>
