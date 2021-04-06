<!-- 影院管理员页面 -->
<template>
  <div id="admin">
    <div class="admin-div">
      <el-tabs type="border-card" v-model="editableTabsValue">
        <!-- 管理员列表 -->
        <el-tab-pane>
          <span slot="label"><i class="el-icon-s-custom"></i> 管理员列表</span>
          <div class="admin-body">
            <div class="admin-list">
              <el-table
                max-height="700px"
                ref="multipleAdminTable"
                stripe
                @selection-change="handleAdminTableSelectionChange"
                :data="AdminList"
                tooltip-effect="dark"
                class="admintable"
              >
                <el-table-column align="center" type="selection" width="50"></el-table-column>
                <el-table-column align="center" prop="AdminID" label="ID" width="80"></el-table-column>
                <el-table-column
                  align="center"
                  prop="AdminName"
                  label="名称"
                  width="250"
                  show-overflow-tooltip
                ></el-table-column>
                <el-table-column align="center" prop="Phone" label="联系方式" width="150"></el-table-column>
                <el-table-column label="操作" align="center" width="360">
                  <template slot-scope="scope">
                    <el-button type="primary" icon="el-icon-s-tools" size="mini" @click="toUpdatePermissions(scope.row)"
                      >设为系统管理员</el-button
                    >
                    <el-button
                      type="warning"
                      icon="el-icon-office-building"
                      size="mini"
                      @click="toManageList(scope.row)"
                      >管理的影院</el-button
                    >
                    <el-button
                      size="mini"
                      icon="el-icon-delete-solid"
                      type="danger"
                      @click="toDeleteOneAdmin(scope.row)"
                      >删除</el-button
                    >
                  </template>
                </el-table-column>
                <el-table-column></el-table-column>
              </el-table>
            </div>
            <div class="admin-footer">
              <el-button
                @click="toDeleteSomeAdmin"
                :disabled="this.AdminIDs.length == 0"
                type="danger"
                icon="el-icon-delete"
                style="float: left;margin-top: 10px;margin-left: 10px;margin-bottom: 10px"
                >批量删除</el-button
              >
              <el-pagination
                background
                @size-change="handleAdminTableSizeChange"
                @current-change="handleAdminTablePageChange"
                :page-sizes="[5, 10, 20, 30, 50, 100]"
                :page-size="this.AdminTableSize"
                :current-page="this.AdminTablePage"
                layout="total,sizes,prev,pager,next,jumper"
                :total="AdminCount"
              >
              </el-pagination>
            </div>
          </div>
        </el-tab-pane>
        <!-- 管理的影院列表 -->
        <el-tab-pane :disabled="this.AdminInfo.AdminID == null">
          <span slot="label"
            ><i :class="this.AdminInfo.AdminID == null ? 'el-icon-lock' : 'el-icon-office-building'"></i>
            {{ this.AdminInfo.AdminID == null ? ' 管理的影院' : ' 管理员ID:' + this.AdminInfo.AdminID }}</span
          >
          <div class="admin-body">
            <div class="admin-list">
              <el-table
                max-height="700px"
                ref="multipleManageTable"
                stripe
                @selection-change="handleManageTableSelectionChange"
                :data="ManageList"
                tooltip-effect="dark"
                class="admintable"
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
                <el-table-column label="操作" align="center" width="100">
                  <template slot-scope="scope">
                    <el-button
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
            <div class="admin-footer">
              <el-button
                @click="toDeleteSomeCinema"
                :disabled="this.CinemaIDs.length == 0"
                type="danger"
                icon="el-icon-delete"
                style="float: left;margin-top: 10px;margin-left: 10px;margin-bottom: 10px"
                >批量删除</el-button
              >
              <el-button
                @click="openAddManage"
                type="primary"
                icon="el-icon-circle-plus"
                style="float: left;margin-top: 10px;margin-left: 10px;margin-bottom: 10px"
                >添加</el-button
              >
              <el-pagination
                background
                @size-change="handleManageTableSizeChange"
                @current-change="handleManageTablePageChange"
                :page-sizes="[5, 10, 20, 30, 50, 100]"
                :page-size="this.ManageTableSize"
                :current-page="this.ManageTablePage"
                layout="total,sizes,prev,pager,next,jumper"
                :total="ManageCount"
              >
              </el-pagination>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
      <!-- 添加管理 -->
      <el-dialog title="添加管理影院" :visible.sync="showAddManage" class="add-admin-dialog">
        <el-input v-model="this.AdminInfo.AdminID" disabled>
          <template slot="prepend">管理员ID</template>
        </el-input>
        <span style="font-size: 20px;margin-top: 20px;">城市:&nbsp;</span>
        <el-cascader
          v-model="City"
          :props="{ emitPath: false, children: 'city', expandTrigger: 'hover' }"
          placeholder="试试搜索城市"
          :options="cityList"
          filterable
          style="margin-top: 20px;width: 309px;"
        ></el-cascader>
        <span style="font-size: 20px;margin-top: 20px;">影院:&nbsp;</span>
        <el-select
          :disabled="cinemaSelectDisabled"
          style="margin-top: 20px;width: 309px;"
          v-model="CinemaID"
          filterable
          placeholder="试试搜索影院"
        >
          <el-option
            v-for="item in CityCimenaList"
            :key="item.CinemaID"
            :label="item.CinemaName"
            :value="item.CinemaID"
          >
          </el-option>
        </el-select>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showAddManage = false">取 消</el-button>
          <el-button :disabled="addManageDisabled" type="primary" @click="toAddManage" icon="el-icon-finished"
            >确认修改</el-button
          >
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import arealist from '../assets/js/city'
import {
  reqGetCinemaAdmins,
  reqPermissions,
  reqDeleteAdmins,
  reqGetManage,
  reqGetCityCinemas,
  reqAddManage,
  reqDeleteManage,
} from '../api'
export default {
  name: 'Admin',
  data() {
    //这里存放数据
    return {
      AdminList: [], //管理员列表
      ManageList: [], //某个管理员管理的影院
      cityList: arealist,
      AdminTablePage: 1,
      AdminTableSize: 10,
      ManageTablePage: 1,
      ManageTableSize: 10,
      AdminCount: 0,
      ManageCount: 0,
      AdminInfo: {}, //影院管理员对象
      AdminIDs: [], //要删除的管理员ID数组
      CinemaIDs: [], //要删除管理的影院ID数组
      showAddManage: false,
      City: '',
      CityCimenaList: [], //所有的影院列表
      CinemaID: '',
      editableTabsValue: '0',
    }
  },
  //监听属性 类似于data概念
  computed: {
    addManageDisabled() {
      return this.CinemaID == '' || this.City == ''
    },
    cinemaSelectDisabled() {
      return this.City == ''
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/admin') {
        this.getAdminList()
      }
    },
    AdminTablePage() {
      this.getAdminList()
    },
    AdminTableSize() {
      this.getAdminList()
    },
    ManageTablePage() {
      this.getManageList()
    },
    ManageTableSize() {
      this.getManageList()
    },
    City() {
      this.getCityCinemaList()
    },
  },
  //方法集合
  methods: {
    //改变管理员表格每页显示条数
    handleAdminTableSizeChange(val) {
      this.AdminTableSize = val
    },
    //改变管理员表格当前页
    handleAdminTablePageChange(val) {
      this.AdminTablePage = val
    },
    //管理员表格多选发生改变时
    handleAdminTableSelectionChange(val) {
      this.AdminIDs = val.map((item) => {
        return item.AdminID
      })
    },
    //改变管理的影院列表每页显示条数
    handleManageTableSizeChange(val) {
      this.ManageTableSize = val
    },
    //改变管理的影院列表当前页
    handleManageTablePageChange(val) {
      this.ManageTablePage = val
    },
    //多选发生改变时
    handleManageTableSelectionChange(val) {
      this.CinemaIDs = val.map((item) => {
        return item.CinemaID
      })
    },
    openAddManage() {
      this.showAddManage = true
    },
    //清空管理员选择
    toggleSelectionAdmin(rows) {
      if (rows) {
        rows.forEach((row) => {
          this.$refs.multipleAdminTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleAdminTable.clearSelection()
      }
    },
    //清空影院选择
    toggleSelectionManage(rows) {
      if (rows) {
        rows.forEach((row) => {
          this.$refs.multipleManageTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleManageTable.clearSelection()
      }
    },
    //设置为系统管理员
    toUpdatePermissions(admin) {
      const AdminID = admin.AdminID
      this.$confirm('此操作将该管理员设置为系统管理员, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqPermissions({ AdminID })
          if (result.code == 0) {
            this.getAdminList()
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '设置成功!',
              center: true, //文字居中
            })
          } else {
            this.$message({
              showClose: true, //可关闭
              type: 'error',
              message: result.message,
              center: true, //文字居中
            })
          }
        })
        .catch(() => {})
    },
    //去查看管理的影院列表
    toManageList(admin) {
      this.AdminInfo = admin
      this.getManageList()
      this.editableTabsValue = '1'
    },
    //获取影院管理员
    async getAdminList() {
      let result
      const Page = this.AdminTablePage
      const Size = this.AdminTableSize
      result = await reqGetCinemaAdmins({ Page, Size })
      if (result.code == 0) {
        this.AdminCount = result.count
        this.AdminList = result.data
      } else {
        this.AdminCount = 0
        this.AdminTablePage = 1
        this.AdminList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //系统管理员获取影院管理员管理的影院列表
    async getManageList() {
      let result
      const AdminID = this.AdminInfo.AdminID
      const Page = this.ManageTablePage
      const Size = this.ManageTableSize
      result = await reqGetManage({ AdminID, Page, Size })
      if (result.code == 0) {
        this.ManageCount = result.count
        this.ManageList = result.data
      } else {
        this.ManageCount = 0
        this.ManageTablePage = 1
        this.ManageList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    //根据城市获取影院列表
    async getCityCinemaList() {
      let result
      const { City } = this
      result = await reqGetCityCinemas({ City })
      if (result.code == 0) {
        this.CityCimenaList = result.data
      } else {
        this.CityCimenaList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //给影院管理员添加管理的影院
    async toAddManage() {
      let result
      const AdminID = this.AdminInfo.AdminID
      const { CinemaID } = this
      result = await reqAddManage({ AdminID, CinemaID })
      if (result.code == 0) {
        this.ManageList.unshift(result.data)
        this.ManageCount++
        this.showAddManage = false
        this.$message({
          showClose: true, //可关闭
          message: '添加成功',
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
    //删除某个管理员
    toDeleteOneAdmin(admin) {
      const { AdminIDs } = this
      AdminIDs[0] = admin.AdminID
      this.$confirm('此操作将永久删除该管理员, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteAdmins({ AdminIDs })
          if (result.code == 0) {
            this.getAdminList()
            this.AdminIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.AdminIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'error',
              message: result.message,
              center: true, //文字居中
            })
          }
        })
        .catch(() => {
          this.toggleSelectionAdmin()
          this.AdminIDs = []
        })
    },
    //批量删除管理员
    toDeleteSomeAdmin() {
      const { AdminIDs } = this
      this.$confirm('此操作将永久删除已选管理员, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteAdmins({ AdminIDs })
          if (result.code == 0) {
            this.AdminIDs = []
            this.getAdminList()
            this.$message({
              showClose: true, //可关闭
              type: 'success',
              message: '删除成功!',
              center: true, //文字居中
            })
          } else {
            this.AdminIDs = []
            this.$message({
              showClose: true, //可关闭
              type: 'error',
              message: result.message,
              center: true, //文字居中
            })
          }
        })
        .catch(() => {
          this.toggleSelectionAdmin()
          this.AdminIDs = []
        })
    },
    //删除某个管理的影院
    toDeleteOneCinema(manage) {
      const AdminID = this.AdminInfo.AdminID
      const { CinemaIDs } = this
      CinemaIDs[0] = manage.CinemaID
      this.$confirm('此操作将删除该管理的影院, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteManage({ AdminID, CinemaIDs })
          if (result.code == 0) {
            this.getManageList()
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
          this.toggleSelectionManage()
          this.CinemaIDs = []
        })
    },
    //批量删除某些管理的影院
    toDeleteSomeCinema() {
      const AdminID = this.AdminInfo.AdminID
      const { CinemaIDs } = this
      this.$confirm('此操作将删除已选管理的影院, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteManage({ AdminID, CinemaIDs })
          if (result.code == 0) {
            this.getManageList()
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
          this.toggleSelectionManage()
          this.CinemaIDs = []
        })
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getAdminList()
  },
}
</script>
<style>
@import url('../assets/css/admin.css');
</style>
