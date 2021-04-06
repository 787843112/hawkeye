<!-- 我的页面 -->
<template>
  <div id="myCenter">
    <div class="myCenter-header">
      <h1>我 的</h1>
    </div>
    <div class="UserAvatar">
      <van-uploader :after-read="uploadAvatar" :before-read="beforeUpload">
        <van-image
          round
          height="7rem"
          style="margin-top: 8px;"
          width="7rem"
          :src="'static/avatars/' + this.UserInfo.Avatar"
        ></van-image>
      </van-uploader>
      <p>
        {{ UserInfo.Nickname != '' ? UserInfo.Nickname : UserInfo.UserName }}
      </p>
    </div>
    <van-divider :style="{ borderColor: '#ff0000', padding: '0 16px' }"></van-divider>
    <div class="grid">
      <van-grid :column-num="2" :gutter="30">
        <van-grid-item icon="coupon" text="电影票" @click="toOrder"></van-grid-item>
        <van-grid-item icon="eye" text="看过" @click="toSeen"></van-grid-item>
        <van-grid-item icon="like-o" text="想看" @click="toWant"></van-grid-item>
        <van-grid-item icon="setting-o" text="个人信息" @click="toInfo"></van-grid-item>
      </van-grid>
    </div>
    <van-divider :style="{ borderColor: '#ff0000', padding: '0 16px' }"></van-divider>
    <div class="logout">
      <van-button round loading-text="注销中..." size="large" type="danger" :loading="logoutLoading" @click="toLogout"
        >注 销
      </van-button>
    </div>
  </div>
</template>

<script>
import { reqUploadAvatar } from '../api'
export default {
  name: 'MyCenter',
  data() {
    return {
      logoutLoading: false,
    }
  },
  //监听属性 类似于data概念
  computed: {
    UserInfo() {
      return this.$store.state.userInfo
    },
  },
  //方法集合
  methods: {
    //上传前校验
    beforeUpload(file) {
      const isIMG = file.type === 'image/jpeg' || file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isIMG) {
        this.$toast('头像只能是JPG或者PNG格式!')
      }
      if (!isLt2M) {
        this.$toast('头像大小不能超过2MB!')
      }
      return isIMG && isLt2M
    },
    //上传头像
    async uploadAvatar(file) {
      let result
      var formData = new FormData()
      formData.append('file', file.file)
      result = await reqUploadAvatar(formData)
      if (result.code == 0) {
        this.$store.dispatch('recordUserInfo', result.data)
        this.$toast('修改头像成功')
      } else {
        this.$toast(result.message)
      }
    },
    //去电影票页面
    toOrder() {
      this.$router.replace('/myOrder')
    },
    //去看过页面
    toSeen() {
      this.$router.replace('/mySeen')
    },
    //去想看页面
    toWant() {
      this.$router.replace('/myWant')
    },
    //去设置页面
    toInfo() {
      this.$router.replace('/myInfo')
    },
    //注销
    toLogout() {
      this.logoutLoading = true
      this.$dialog
        .confirm({
          title: '确定注销吗?',
        })
        .then(() => {
          this.$store.dispatch('logout')
          this.$router.replace('/movie')
          this.logoutLoading = false
        })
        .catch(() => {
          this.logoutLoading = false
        })
    },
  },
}
</script>
<style>
@import url('../assets/css/myCenter.css');
</style>
