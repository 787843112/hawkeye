<!-- 影院页面 -->
<template>
  <div id="cinema">
    <div class="cinema-div">
      <div class="cinema-header">
        <el-input
          clearable
          placeholder="按影院名称搜索"
          prefix-icon="el-icon-search"
          v-model="CinemaName"
          class="search-cinema-name"
        ></el-input>
        <el-cascader
          v-model="City"
          :props="{ emitPath: false, children: 'city', expandTrigger: 'hover' }"
          placeholder="试试搜索"
          :options="cityList"
          filterable
          style="margin-left: 4%;width: 20%;"
        ></el-cascader>
        <el-button
          v-if="this.$store.state.adminInfo.Permissions == 1"
          @click="openAddCinema"
          style="margin-left: 3%;"
          type="primary"
          icon="el-icon-plus"
          >添加影院</el-button
        >
        <el-button
          v-if="this.$store.state.adminInfo.Permissions == 1"
          @click="toDeleteSomeCinema"
          style="margin-left: 3%;"
          type="danger"
          :disabled="this.CinemaIDs.length == 0"
          icon="el-icon-delete"
          >批量删除</el-button
        >
      </div>
      <!-- 添加影院 -->
      <el-dialog title="添加影院" :visible.sync="showAddCinema" class="add-cinema-dialog">
        <el-input maxlength="32" show-word-limit v-model="addCinema.CinemaName" placeholder="影院名称" clearable>
          <template slot="prepend">影院名称</template></el-input
        >
        <el-cascader
          style="margin-top: 20px;width: 460px;"
          v-model="addCinema.City"
          :props="{ emitPath: false, children: 'city', expandTrigger: 'hover' }"
          placeholder="试试搜索"
          :options="cityList"
          filterable
        >
        </el-cascader>
        <el-input
          style="margin-top: 20px;"
          type="textarea"
          maxlength="50"
          show-word-limit
          v-model="addCinema.Address"
          placeholder="请输入详细地址"
        ></el-input>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showAddCinema = false">取 消</el-button>
          <el-button :disabled="addCinemaDisabled" type="primary" @click="toAddCinema" icon="el-icon-finished"
            >确认添加</el-button
          >
        </div>
      </el-dialog>
      <!-- 影院列表 -->
      <div class="cinema-body">
        <div class="cinema-list">
          <el-table
            max-height="670px"
            ref="multipleTable"
            stripe
            @selection-change="handleSelectionChange"
            :data="CinemaList"
            tooltip-effect="dark"
            class="cinematable"
          >
            <el-table-column align="center" type="selection" width="50"></el-table-column>
            <el-table-column align="center" prop="CinemaID" label="ID" width="80"></el-table-column>
            <el-table-column
              align="center"
              prop="CinemaName"
              label="名称"
              width="250"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              align="center"
              prop="City"
              label="所在城市"
              width="80"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              align="center"
              prop="Address"
              label="详细地址"
              width="450"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column label="操作" align="center" width="290">
              <template slot-scope="scope">
                <el-button type="primary" icon="el-icon-edit" size="mini" @click="openUpdateCinema(scope.row)"
                  >修改</el-button
                >
                <el-button type="success" icon="el-icon-data-analysis" size="mini" @click="toHallList(scope.row)"
                  >影厅列表</el-button
                >
                <el-button
                  v-if="$store.state.adminInfo.Permissions == 1"
                  size="mini"
                  icon="el-icon-delete-solid"
                  type="danger"
                  @click="toDeleteOneCinema(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
            <el-table-column></el-table-column>
          </el-table>
        </div>
        <div class="cinema-footer">
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
      <!-- 修改影院 -->
      <el-dialog title="修改影院" :visible.sync="showUpdateCinema" class="add-cinema-dialog">
        <el-input v-model="updateCinema.CinemaID" disabled>
          <template slot="prepend">影院ID</template>
        </el-input>
        <el-input
          style="margin-top: 20px;"
          maxlength="32"
          show-word-limit
          v-model="updateCinema.CinemaName"
          placeholder="影院名称"
          clearable
        >
          <template slot="prepend">影院名称</template></el-input
        >
        <el-cascader
          style="margin-top: 20px;width: 460px;"
          v-model="updateCinema.City"
          :props="{ emitPath: false, children: 'city', expandTrigger: 'hover' }"
          :options="cityList"
          filterable
        >
        </el-cascader>
        <el-input
          style="margin-top: 20px;"
          type="textarea"
          maxlength="50"
          show-word-limit
          v-model="updateCinema.Address"
          placeholder="请输入详细地址"
        ></el-input>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showUpdateCinema = false">取 消</el-button>
          <el-button :disabled="updateCinemaDisabled" type="primary" @click="toUpdateCinema" icon="el-icon-finished"
            >确认修改</el-button
          >
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { reqGetCinemas, reqDeleteCinemas, reqAddCinema, reqUpdateCinema, reqGetAdminManage } from '../api'
import arealist from '../assets/js/city'
export default {
  name: 'Cinema',
  data() {
    //这里存放数据
    return {
      cityList: arealist,
      CinemaName: '',
      City: '',
      Page: 1,
      Size: 10,
      count: 0,
      addCinema: {
        CinemaName: '',
        City: '',
        Address: '',
      },
      updateCinema: {
        CinemaName: '',
        City: '',
        Address: '',
      },
      CinemaIDs: [],
      showAddCinema: false,
      showUpdateCinema: false,
      CinemaList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    addCinemaDisabled() {
      return this.addCinema.CinemaName == '' || this.addCinema.City == '' || this.addCinema.Address == ''
    },
    updateCinemaDisabled() {
      return this.updateCinema.CinemaName == '' || this.updateCinema.City == '' || this.updateCinema.Address == ''
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/cinema') {
        this.toGetCinemaList()
      }
    },
    CinemaName() {
      this.toGetCinemaList()
    },
    City() {
      this.toGetCinemaList()
    },
    Page() {
      this.toGetCinemaList()
    },
    Size() {
      this.toGetCinemaList()
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
      this.CinemaIDs = val.map((item) => {
        return item.CinemaID
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
    openAddCinema() {
      this.addCinema.CinemaName = ''
      this.addCinema.Address = ''
      this.showAddCinema = true
    },
    openUpdateCinema(cinema) {
      this.updateCinema.CinemaID = cinema.CinemaID
      this.updateCinema.CinemaName = cinema.CinemaName
      this.updateCinema.City = cinema.City
      this.updateCinema.Address = cinema.Address
      this.showUpdateCinema = true
    },
    toGetCinemaList() {
      if (this.$store.state.adminInfo.Permissions == 1) {
        this.getCinemaList()
      } else {
        this.getManageCinema()
      }
    },
    //系统管理员获取影院列表
    async getCinemaList() {
      let result
      const { CinemaName, City, Page, Size } = this
      result = await reqGetCinemas({ CinemaName, City, Page, Size })
      if (result.code == 0) {
        this.count = result.count
        this.CinemaList = result.data
      } else {
        this.count = 0
        this.Page = 1
        this.CinemaList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    //影院管理员获取影院列表
    async getManageCinema() {
      let result
      const { City, CinemaName, Page, Size } = this
      result = await reqGetAdminManage({ City, CinemaName, Page, Size })
      if (result.code == 0) {
        this.count = result.count
        this.CinemaList = result.data
      } else {
        this.count = 0
        this.Page = 1
        this.CinemaList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    //添加影院
    async toAddCinema() {
      let result
      const CinemaName = this.addCinema.CinemaName
      const City = this.addCinema.City
      const Address = this.addCinema.Address
      result = await reqAddCinema({ CinemaName, City, Address })
      if (result.code == 0) {
        this.CinemaList.unshift(result.data)
        //this.getCinemaList()
        this.showAddCinema = false
        this.$message({
          showClose: true, //可关闭
          message: '添加成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.showAddCinema = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //修改影院
    async toUpdateCinema() {
      let result
      const CinemaID = this.updateCinema.CinemaID
      const CinemaName = this.updateCinema.CinemaName
      const City = this.updateCinema.City
      const Address = this.updateCinema.Address
      result = await reqUpdateCinema({ CinemaID, CinemaName, City, Address })
      if (result.code == 0) {
        this.getCinemaList()
        this.showUpdateCinema = false
        this.$message({
          showClose: true, //可关闭
          message: '修改成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.showUpdateCinema = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //去查看影院的影厅列表
    toHallList(cinema) {
      this.$store.dispatch('recordCinemaInfo', cinema)
      this.$router.replace('/hall')
    },
    //删除某个影院
    toDeleteOneCinema(cinema) {
      const { CinemaIDs } = this
      CinemaIDs[0] = cinema.CinemaID
      this.$confirm('此操作将永久删除该影院, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteCinemas({ CinemaIDs })
          if (result.code == 0) {
            this.getCinemaList()
            this.CinemaIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.CinemaIDs = []
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
          this.CinemaIDs = []
        })
    },
    //批量删除某些影院
    toDeleteSomeCinema() {
      const { CinemaIDs } = this
      this.$confirm('此操作将永久删除已选影院, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteCinemas({ CinemaIDs })
          if (result.code == 0) {
            this.CinemaIDs = []
            this.getCinemaList()
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.CinemaIDs = []
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
          this.CinemaIDs = []
        })
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.toGetCinemaList()
  },
}
</script>
<style>
@import url('../assets/css/cinema.css');
</style>
