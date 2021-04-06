<!-- 登录注册 -->
<template>
  <div class="login-div">
    <el-card shadow="always" class="card-box">
      <img src="../assets/images/login_logo.png" />
      <div class="input-box">
        <div class="login-input">
          <el-input
            style="margin-top: 50px;"
            v-show="!isRegister"
            prefix-icon="el-icon-user-solid"
            clearable
            v-model="AdminName"
            placeholder="请输入管理员名"
          ></el-input>
          <el-input
            v-show="isRegister"
            prefix-icon="el-icon-user-solid"
            maxlength="20"
            show-word-limit
            clearable
            v-model="AdminName"
            placeholder="请输入管理员名"
          ></el-input>
          <el-input
            prefix-icon="el-icon-view"
            clearable
            v-model="Password"
            placeholder="请输入密码"
            show-password
          ></el-input>
        </div>
        <el-collapse-transition>
          <div class="register-input" v-show="isRegister">
            <el-input
              prefix-icon="el-icon-view"
              clearable
              v-model="repassword"
              placeholder="请再次输入密码"
              show-password
            ></el-input>
            <el-input prefix-icon="el-icon-phone" clearable v-model="Phone" placeholder="请输入手机号"></el-input>
          </div>
        </el-collapse-transition>
      </div>
      <div class="button-group" v-show="!isRegister">
        <el-button-group>
          <el-button
            v-show="!isRegister"
            type="primary"
            round
            :disabled="loginDisabled"
            @click="toLogin"
            :loading="isLoading"
            >登录</el-button
          >
          <el-button type="warning" round @click="isRegister = !isRegister">去注册</el-button>
        </el-button-group>
      </div>
      <div class="button-group" v-show="isRegister">
        <el-button-group>
          <el-button
            v-show="isRegister"
            type="primary"
            round
            :disabled="registerDisabled"
            @click="toRegister"
            :loading="isLoading"
            >注册</el-button
          >
          <el-button type="warning" round @click="isRegister = !isRegister">去登录</el-button>
        </el-button-group>
      </div>
    </el-card>
    <el-calendar v-model="calendarDate" class="calendar-box"> </el-calendar>
  </div>
</template>

<script>
import { reqLogin, reqRegister } from '../api'
export default {
  name: 'Login',
  data() {
    //这里存放数据
    return {
      AdminName: '',
      Password: '',
      repassword: '',
      Phone: '',
      isRegister: false,
      calendarDate: new Date(),
      isLoading: false,
    }
  },
  //监听属性 类似于data概念
  computed: {
    rightPhone() {
      // 利用正则对手机号进行匹配，返回布尔值
      return /^1\d{10}$/.test(this.Phone)
    },
    rightAdminName() {
      return /^[a-zA-Z0-9]{1,20}$/.test(this.AdminName)
    },
    rightPassword() {
      return /^[a-zA-Z0-9_]{8,20}$/.test(this.Password)
    },
    loginDisabled() {
      return this.AdminName == '' || this.Password == ''
    },
    registerDisabled() {
      return this.AdminName == '' || this.Password == '' || this.rePassword == '' || this.Phone == ''
    },
  },
  //监控data中的数据变化
  watch: {
    isRegister() {
      this.isLoading = false
    },
    $route(to) {
      if (to.path == '/login') {
        this.AdminName = ''
        this.Password = ''
        this.repassword = ''
        this.Phone = ''
      }
    },
  },
  //方法集合
  methods: {
    //异步登录
    async toLogin() {
      this.isLoading = true
      let result
      const { AdminName, Password } = this
      result = await reqLogin({ AdminName, Password })
      if (result.code == 0) {
        this.isLoading = false
        const admin = result.data
        this.$store.dispatch('recordAdminInfo', admin)
        this.$router.replace('/movie')
      } else {
        this.isLoading = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    async toRegister() {
      this.isLoading = true
      let result
      const { AdminName, Password, Phone } = this
      if (!this.rightAdminName) {
        this.isLoading = false
        this.$message({
          showClose: true, //可关闭
          message: '管理员名只能是20位以内的数字和字母的组合!',
          type: 'warning', //提示类型
          center: true, //文字居中
        })
        return
      }
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
      if (!this.rightPhone) {
        this.isLoading = false
        this.$message({
          showClose: true, //可关闭
          message: '请输入正确的手机号!',
          type: 'warning', //提示类型
          center: true, //文字居中
        })
        return
      }
      if (this.Password != this.repassword) {
        this.isLoading = false
        this.$message({
          showClose: true, //可关闭
          message: '两次密码不一样!',
          type: 'error', //提示类型
          center: true, //文字居中
        })
        return
      }
      result = await reqRegister({ AdminName, Password, Phone })
      if (result.code == 0) {
        this.isLoading = false
        const admin = result.data
        this.$store.dispatch('recordAdminInfo', admin)
        this.$router.replace('/movie')
        this.$message({
          showClose: true, //可关闭
          message: '注册成功!',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.isLoading = false
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
  },
}
</script>
<style>
@import '../assets/css/login.css';
</style>
