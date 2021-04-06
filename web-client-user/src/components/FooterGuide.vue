<!-- 底部导航 -->
<template>
  <div id="footerGuide">
    <van-tabbar route>
      <van-tabbar-item replace icon="wap-home" to="/movie">
        电影
      </van-tabbar-item>
      <van-tabbar-item replace icon="shop" to="/cinema">
        影院
      </van-tabbar-item>
      <van-tabbar-item replace icon="manager" :to="this.$store.state.userInfo.UserID != null ? '/myCenter' : '/login'">
        我的
      </van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
export default {
  name: 'FooterGuide',
}
</script>
