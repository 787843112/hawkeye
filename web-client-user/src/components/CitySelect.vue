<!-- 城市选择器 -->
<template>
  <div id="citySelect">
    <p @click="showAddress" style="margin-top: 11px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;">
      <van-icon name="location"></van-icon>
      {{ this.city }}
    </p>
    <van-popup round v-model="citySelectShow" position="top" style="overflow: hidden" :style="{ height: '30%' }">
      <van-area
        title="请选择城市"
        value="110101"
        :area-list="cityList"
        :columns-num="2"
        :visible-item-count="4"
        @cancel="onAddressCancel"
        @confirm="onAddressConfirm"
      >
      </van-area>
    </van-popup>
  </div>
</template>

<script>
import AddressInfo from '../assets/js/area'
export default {
  name: 'CitySelect',
  data() {
    //这里存放数据
    return {
      citySelectShow: false,
      cityList: AddressInfo,
    }
  },
  //监听属性 类似于data概念
  computed: {
    city() {
      return this.$store.state.city
    },
  },
  //监控data中的数据变化
  watch: {},
  //方法集合
  methods: {
    showAddress() {
      this.citySelectShow = true
    },
    onAddressConfirm(val) {
      this.citySelectShow = false
      this.$store.dispatch('recordCity', val[1].name)
    },
    onAddressCancel() {
      this.citySelectShow = false
    },
  },
}
</script>
