<!-- 个人信息页面 -->
<template>
  <div id="myInfo">
    <van-nav-bar fixed left-arrow style="background-color: rgb(84, 157, 240)" @click-left="onClickLeft">
      <van-icon color="#ff0000" name="arrow-left" size="30" slot="left"></van-icon>
      <h1
        slot="title"
        style="overflow: hidden;text-overflow: ellipsis;
                white-space: nowrap;"
      >
        个&nbsp;人&nbsp;信&nbsp;息
      </h1>
    </van-nav-bar>
    <van-cell-group style="overflow: auto;margin-top: 46px;">
      <van-cell icon="card" title="UID" :value="UserInfo.UserID"></van-cell>
      <van-cell icon="manager" title="用户名" :value="UserInfo.UserName"></van-cell>
      <van-field
        v-model="Nickname"
        clearable
        label="昵称"
        input-align="right"
        left-icon="comment-circle"
        right-icon="arrow"
        maxlength="10"
        :placeholder="UserInfo.Nickname"
      />
      <van-field
        v-model="Phone"
        label="手机号"
        type="digit"
        clearable
        input-align="right"
        left-icon="phone"
        right-icon="arrow"
        maxlength="20"
        :placeholder="UserInfo.Phone"
      />
      <van-cell icon="clock" title="注册日期" :value="UserInfo.CreateTime.slice(0, 10)"></van-cell>
      <van-field
        v-model="Password"
        label="修改密码"
        clearable
        input-align="right"
        left-icon="warning"
        :type="showNewPassword ? 'text' : 'password'"
        :right-icon="Password == '' ? 'arrow' : showNewPassword ? 'eye-o' : 'closed-eye'"
        @click-right-icon="showPassword"
        maxlength="20"
        placeholder="请输入需要修改的密码"
      />
    </van-cell-group>
    <div style="width: 70%;margin: 30px auto;">
      <van-button
        round
        loading-text="修改中..."
        size="large"
        type="primary"
        :disabled="updateDisabled"
        :loading="updateLoading"
        @click="updateUserInfo"
        >保&nbsp;存&nbsp;修&nbsp;改
      </van-button>
    </div>
  </div>
</template>

<script>
import { reqUpdate } from '../api'
export default {
  name: 'MyInfo',
  data() {
    //这里存放数据
    return {
      Nickname: '',
      Phone: '',
      Password: '',
      showNewPassword: false,
      updateLoading: false,
    }
  },
  //监听属性 类似于data概念
  computed: {
    UserInfo() {
      return this.$store.state.userInfo
    },
    updateDisabled() {
      return this.Nickname == '' && this.Password == '' && this.Phone == ''
    },
    rightPassword() {
      return /^[a-zA-Z0-9_]{8,20}$/.test(this.Password)
    },
    rightPhone() {
      // 利用正则对手机号进行匹配，返回布尔值
      return /^1\d{10}$/.test(this.Phone)
    },
  },
  //监控data中的数据变化
  watch: {},
  //方法集合
  methods: {
    //返回
    onClickLeft() {
      this.$router.replace('/myCenter')
    },
    //显示密码
    showPassword() {
      this.showNewPassword = !this.showNewPassword
    },
    //修改用户信息
    async updateUserInfo() {
      this.updateLoading = true
      let result
      const { Nickname, Password, Phone } = this
      if (this.Phone != '') {
        if (!this.rightPhone) {
          this.updateLoading = false
          this.$toast('请输入正确的手机号!')
          return
        }
      }
      if (this.Password != '') {
        if (!this.rightPassword) {
          this.updateLoading = false
          this.$toast('密码只能是8到20位的数字,字母,下划线的组合!')
          return
        }
      }
      result = await reqUpdate({ Nickname, Password, Phone })
      if (result.code == 0) {
        this.$toast('修改成功')
        this.$store.dispatch('recordUserInfo', result.data)
        this.Nickname = ''
        this.Password = ''
        this.Phone = ''
        this.showNewPassword = false
        this.updateLoading = false
      } else {
        this.$toast(result.message)
        this.updateLoading = false
      }
    },
  },
}
</script>
