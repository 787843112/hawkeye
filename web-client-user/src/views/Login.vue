<!-- 登录注册页面 -->
<template>
  <div id="login">
    <van-nav-bar left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1 slot="title">鹰 眼 电 影</h1>
    </van-nav-bar>
    <div class="headLogo">
      <van-image round height="9.5rem" src="logo.jpg" width="9.5rem"></van-image>
    </div>
    <!-- 登录 -->
    <div v-show="isLogin" class="login-box">
      <van-cell-group>
        <van-field
          clearable
          v-model="UserName"
          label="用户名"
          left-icon="contact"
          placeholder="请输入用户名"
          type="text"
          maxlength="20"
        >
        </van-field>
        <van-field
          clearable
          v-model="Password"
          label="密 码"
          left-icon="warning-o"
          placeholder="请输入密码"
          :type="showloginPassword ? 'text' : 'password'"
          :right-icon="showloginPassword ? 'eye-o' : 'closed-eye'"
          maxlength="20"
          @click-right-icon="showLoginPassword"
        >
        </van-field>
      </van-cell-group>
      <p @click="switchBox">免费注册</p>
      <van-button
        round
        loading-text="登录中..."
        loading-type="spinner"
        size="large"
        style="margin-top: 20px"
        type="primary"
        :loading="loginLoading"
        :disabled="loginDisabled"
        @click="login"
        >登 录
      </van-button>
    </div>
    <!-- 注册 -->
    <div v-show="!isLogin" class="register-box">
      <van-cell-group>
        <van-field
          clearable
          required
          v-model="UserName"
          label="用户名"
          left-icon="contact"
          placeholder="请输入用户名"
          type="text"
          maxlength="20"
        ></van-field>
        <van-field
          clearable
          required
          v-model="Password"
          label="密 码"
          left-icon="warning-o"
          placeholder="请输入密码"
          :type="showregisterPassword ? 'text' : 'password'"
          :right-icon="showregisterPassword ? 'eye-o' : 'closed-eye'"
          @click-right-icon="showRegisterPassword"
          maxlength="20"
        ></van-field>
        <van-field
          clearable
          required
          v-model="rePassword"
          label="确认密码"
          left-icon="passed"
          placeholder="请再次输入密码"
          :type="showrePassword ? 'text' : 'password'"
          :right-icon="showrePassword ? 'eye-o' : 'closed-eye'"
          @click-right-icon="showRePassword"
          maxlength="20"
        ></van-field>
        <van-field
          clearable
          required
          type="digit"
          v-model="Phone"
          label="手机号"
          left-icon="phone-o"
          placeholder="请输入手机号"
        ></van-field>
      </van-cell-group>
      <p @click="switchBox">已有账号</p>
      <van-button
        plain
        round
        loading-text="注册中..."
        loading-type="spinner"
        size="large"
        style="margin-top: 20px"
        type="warning"
        :disabled="registerDisabled"
        :loading="registerLoading"
        @click="register"
        >注 册
      </van-button>
    </div>
  </div>
</template>

<script>
import { reqLogin, reqRegister } from '../api'
export default {
  name: 'Login',
  data() {
    return {
      UserName: '',
      Password: '',
      rePassword: '',
      Phone: '',
      showloginPassword: false,
      showregisterPassword: false,
      showrePassword: false,
      isLogin: true,
      loginLoading: false,
      registerLoading: false,
    }
  },
  computed: {
    rightUserName() {
      return /^[a-zA-Z0-9]{1,20}$/.test(this.UserName)
    },
    rightPassword() {
      return /^[a-zA-Z0-9_]{8,20}$/.test(this.Password)
    },
    rightPhone() {
      // 利用正则对手机号进行匹配，返回布尔值
      return /^1\d{10}$/.test(this.Phone)
    },
    loginDisabled() {
      return this.UserName == '' || this.Password == ''
    },
    registerDisabled() {
      return this.UserName == '' || this.Password == '' || this.rePassword == '' || this.Phone == ''
    },
  },
  watch: {
    $route() {
      this.isLogin = true
      this.UserName = ''
      this.Password = ''
      this.rePassword = ''
      this.Phone = ''
    },
  },
  //方法集合
  methods: {
    //返回按键
    onClickLeft() {
      this.$router.replace('/movie')
    },
    //切换登录和注册
    switchBox() {
      this.isLogin = !this.isLogin
    },
    //显示登录密码
    showLoginPassword() {
      this.showloginPassword = !this.showloginPassword
    },
    //显示注册密码
    showRegisterPassword() {
      this.showregisterPassword = !this.showregisterPassword
    },
    //显示确认密码
    showRePassword() {
      this.showrePassword = !this.showrePassword
    },
    //登录
    async login() {
      this.loginLoading = true
      let result
      const { UserName, Password } = this
      //发送ajax请求密码登陆
      result = await reqLogin({ UserName, Password })
      // 根据结果数据处理
      if (result.code == 0) {
        // 将user保存到vuex的state
        this.loginLoading = false
        this.$toast.success('登录成功!')
        this.$store.dispatch('recordUserInfo', result.data)
        this.$router.replace('/myCenter')
      } else {
        this.loginLoading = false
        this.$toast(result.message)
      }
    },
    //注册
    async register() {
      this.registerLoading = true
      if (!this.rightUserName) {
        this.registerLoading = false
        this.$toast('用户名只能是20位以内的数字和字母的组合!')
        return
      }
      if (!this.rightPassword) {
        this.registerLoading = false
        this.$toast('密码只能是8到20位的数字,字母,下划线的组合!')
        return
      }
      if (this.Password !== this.rePassword) {
        this.registerLoading = false
        this.$toast('两次密码不一样!')
        return
      }
      if (!this.rightPhone) {
        this.registerLoading = false
        this.$toast('请输入正确的手机号!')
        return
      }
      let result
      const { UserName, Password, Phone } = this
      result = await reqRegister({ UserName, Password, Phone })
      if (result.code == 0) {
        this.registerLoading = false
        this.$store.dispatch('recordUserInfo', result.data)
        this.$router.replace('/myCenter')
      } else {
        this.registerLoading = false
        this.$toast(result.message)
      }
    },
  },
}
</script>
<style>
@import url('../assets/css/login.css');
</style>
