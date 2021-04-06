<!-- 菜单栏 -->
<template>
  <div id="menu">
    <div class="admin-info-box">
      <el-upload action :http-request="uploadAvatar" :before-upload="beforeAvatarUpload" :show-file-list="false">
        <el-avatar :size="100" :src="'static/avatars/' + adminInfo.Avatar" style="margin-top: 20px"></el-avatar
      ></el-upload>
      <span style="display: block">{{ adminInfo.AdminName }}</span>
      <el-dropdown @command="handleCommand" style="margin-top:10px">
        <span style="font-size: 20px"
          >{{ adminInfo.Permissions == 1 ? '系统管理员' : '影院管理员'
          }}<i class="el-icon-arrow-down el-icon--right"></i
        ></span>
        <el-dropdown-menu slot="dropdown" style="text-align: center">
          <el-dropdown-item command="1">修改信息</el-dropdown-item>
          <el-dropdown-item command="2">注销</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <div class="router-menu">
      <el-menu
        :router="true"
        :default-active="this.$route.path"
        class="el-menu-vertical-demo"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <el-menu-item index="/movie">
          <i class="el-icon-film"></i>
          <span slot="title">电影列表</span>
        </el-menu-item>
        <el-menu-item index="/cinema">
          <i class="el-icon-school"></i>
          <span slot="title">{{ adminInfo.Permissions == 1 ? '影院列表' : '我管理的影院' }}</span>
        </el-menu-item>
        <el-menu-item index="/hall" :disabled="this.$store.state.cinemaInfo.CinemaID == null">
          <i class="el-icon-data-analysis"></i>
          <span slot="title">影厅列表</span>
        </el-menu-item>
        <el-menu-item index="/seat" :disabled="this.$store.state.hallInfo.HallID == null">
          <i class="el-icon-s-grid"></i>
          <span slot="title">座位信息</span>
        </el-menu-item>
        <el-menu-item index="/screening" :disabled="this.$store.state.hallInfo.HallID == null">
          <i class="el-icon-time"></i>
          <span slot="title">场次信息</span>
        </el-menu-item>
        <el-menu-item index="/admin" v-if="adminInfo.Permissions == 1">
          <i class="el-icon-s-custom"></i>
          <span slot="title">影院管理员</span>
        </el-menu-item>
      </el-menu>
    </div>
    <el-dialog title="修改信息" :visible.sync="dialogVisible">
      <el-input type="password" v-model="Password" placeholder="请输入要修改的密码"></el-input>
      <el-input v-model="Phone" placeholder="请输入要修改的手机号"></el-input>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button :disabled="updateAdminInfo" type="primary" @click="toUpdateAdminInfo" icon="el-icon-finished"
          >保存修改</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { reqUploadAvatar, reqUpdate } from '../api'
export default {
  name: 'Menu',
  data() {
    //这里存放数据
    return {
      Password: '',
      Phone: '',
      dialogVisible: false,
    }
  },
  //监听属性 类似于data概念
  computed: {
    rightPhone() {
      // 利用正则对手机号进行匹配，返回布尔值
      return /^1\d{10}$/.test(this.Phone)
    },
    rightPassword() {
      return /^[a-zA-Z0-9_]{8,20}$/.test(this.Password)
    },
    adminInfo() {
      return this.$store.state.adminInfo
    },
    updateAdminInfo() {
      return this.Password == '' && this.Phone == ''
    },
  },
  //监控data中的数据变化
  watch: {},
  //方法集合
  methods: {
    //上传头像
    async uploadAvatar(file) {
      let result
      var formData = new FormData()
      formData.append('file', file.file)
      result = await reqUploadAvatar(formData)
      if (result.code == 0) {
        this.$store.dispatch('recordAdminInfo', result.data)
        this.$message({
          showClose: true, //可关闭
          message: '修改管理员头像成功',
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
    //校验头像图片格式及大小
    beforeAvatarUpload(file) {
      const isIMG = file.type === 'image/jpeg' || file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isIMG) {
        this.$message({
          showClose: true, //可关闭
          type: 'error', //提示类型
          message: '头像图片只能是JPG或者PNG格式!',
          center: true, //文字居中
        })
      }
      if (!isLt2M) {
        this.$message({
          showClose: true, //可关闭
          type: 'error', //提示类型
          message: '头像图片大小不能超过2MB!',
          center: true, //文字居中
        })
      }
      return isIMG && isLt2M
    },
    handleCommand(command) {
      if (command == 1) {
        this.openUpdateBox()
      }
      if (command == 2) {
        this.toLogout()
      }
    },
    openUpdateBox() {
      this.Password = ''
      this.Phone = ''
      this.dialogVisible = true
    },
    async toUpdateAdminInfo() {
      if (this.Password != '') {
        if (!this.rightPassword) {
          this.isLoading = false
          this.$message({
            showClose: true, //可关闭
            message: '密码只能是8到20位的数字,字母和下划线的组合!',
            type: 'warning', //提示类型
            center: true, //文字居中
          })
          return
        }
      }
      if (this.Phone != '') {
        if (!this.rightPhone) {
          this.$message({
            showClose: true, //可关闭
            message: '请输入正确的手机号!',
            type: 'warning', //提示类型
            center: true, //文字居中
          })
          return
        }
      }
      const { Password, Phone } = this
      let result = await reqUpdate({ Password, Phone })
      if (result.code == 0) {
        this.dialogVisible = false
        this.$message({
          showClose: true, //可关闭
          message: '修改成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    toLogout() {
      this.$store.dispatch('logout')
      this.$router.replace('/login')
    },
  },
}
</script>
<style>
@import url('../assets/css/menu.css');
</style>
