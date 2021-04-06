<!-- 日期选择器 -->
<template>
  <div id="selectDate">
    <van-tabs animated sticky swipeable v-model="selectDate" :ellipsis="false" :line-width="90">
      <van-tab v-for="(item, index) in tabTitle" :key="item" :title="item" :name="dateArr[index]"></van-tab>
    </van-tabs>
  </div>
</template>

<script>
export default {
  name: 'SelectDate',
  data() {
    //这里存放数据
    return {
      tabTitle: [],
      dateArr: [],
    }
  },
  //监听属性 类似于data概念
  computed: {
    selectDate: {
      get() {
        return this.$store.state.selectDate
      },
      set(val) {
        this.$store.state.selectDate = val
      },
    },
  },
  //监控data中的数据变化
  watch: {},
  //方法集合
  methods: {
    //获取日期
    getDate(day) {
      var date = new Date()
      date.setDate(date.getDate() + day)
      var YY = date.getFullYear() + '-'
      let MD = date.getMonth() + 1 + '月' + date.getDate() + '日'
      var MM = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
      var DD = date.getDate() < 10 ? '0' + date.getDate() : date.getDate()
      if (day == 0) {
        this.tabTitle[0] = '今天' + MD
      } else if (day == 1) {
        this.tabTitle[1] = '明天' + MD
      } else if (day == 2) {
        this.tabTitle[2] = '后天' + MD
      } else {
        this.tabTitle[day] =
          (date.getDay() == 0
            ? '周日'
            : date.getDay() == 1
            ? '周一'
            : date.getDay() == 2
            ? '周二'
            : date.getDay() == 3
            ? '周三'
            : date.getDay() == 4
            ? '周四'
            : date.getDay() == 5
            ? '周五'
            : '周六') + MD
      }
      return YY + MM + DD
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    for (let i = 0; i < 7; i++) {
      this.dateArr[i] = this.getDate(i)
    }
    if (this.$store.state.selectDate == '') {
      this.$store.dispatch('recordSelectDate', this.dateArr[0])
    }
  },
}
</script>
<style>
@import url('../assets/css/selectDate.css');
</style>
