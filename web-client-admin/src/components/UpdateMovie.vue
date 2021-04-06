<!-- 修改电影信息 -->
<template>
  <div id="updatemovie">
    <div class="movie-info-form">
      <el-form ref="form" :model="MovieInfo" label-width="100px">
        <el-form-item label="ID">
          <el-input v-model="MovieInfo.MovieID" disabled></el-input>
        </el-form-item>
        <el-form-item label="名称">
          <el-input
            placeholder="输入电影名称"
            clearable
            v-model="MovieInfo.MovieName"
            maxlength="20"
            show-word-limit
          ></el-input>
        </el-form-item>
        <el-form-item label="导演"
          ><el-input
            placeholder="请输入导演名称"
            clearable
            v-model="MovieInfo.Director"
            maxlength="20"
            show-word-limit
          ></el-input
        ></el-form-item>
        <el-form-item label="主演">
          <el-input
            placeholder="请输入主演"
            clearable
            v-model="MovieInfo.Starring"
            maxlength="255"
            show-word-limit
          ></el-input
        ></el-form-item>
        <el-form-item label="地区">
          <el-input
            placeholder="请输入地区"
            clearable
            prefix-icon="el-icon-place"
            v-model="MovieInfo.Region"
            maxlength="10"
            show-word-limit
          ></el-input
        ></el-form-item>
        <el-form-item label="类型">
          <el-select
            v-model="newMovieTypes"
            multiple
            :multiple-limit="3"
            :placeholder="MovieTypes"
            style="width: 300px;"
          >
            <el-option
              v-for="(item, index) in this.$store.state.allMovieTypes"
              :key="index"
              :label="item.TypeName"
              :value="item.TypeName"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="时长">
          <el-input-number
            :min="1"
            :max="999"
            placeholder="请输入时长(分钟)"
            v-model="MovieInfo.MovieLength"
          ></el-input-number
        ></el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="MovieInfo.Status">
            <el-radio-button :label="1">已下架</el-radio-button>
            <el-radio-button :label="2">正在上映</el-radio-button>
            <el-radio-button :label="3">即将上映</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="上映日期">
          <el-date-picker
            :editable="false"
            style="width: 300px;"
            type="date"
            placeholder="请选择日期"
            value-format="yyyy-MM-dd"
            v-model="MovieInfo.ReleaseTime"
          ></el-date-picker>
        </el-form-item>
      </el-form>
    </div>
    <div class="movie-introduction">
      <el-input
        placeholder="请输入简介"
        type="textarea"
        maxlength="1000"
        show-word-limit
        v-model="MovieInfo.Introduction"
      ></el-input>
    </div>
    <el-upload
      action
      class="movie-info-poster"
      :show-file-list="false"
      :http-request="uploadPoster"
      :before-upload="beforePosterUpload"
    >
      <img :src="'static/posters/' + MovieInfo.Poster" />
    </el-upload>
    <el-button
      icon="el-icon-finished"
      :disabled="updateMovieDisabled"
      @click="toUpdateMovieInfo"
      style="margin-top: 10px;margin-left: 300px"
      type="primary"
      >保存修改</el-button
    >
  </div>
</template>

<script>
import { reqUploadPoster, reqUpdateMovie } from '../api'
export default {
  name: 'UpdataMovie',
  data() {
    //这里存放数据
    return {
      newMovieTypes: [],
      MovieTypes: '',
    }
  },
  //监听属性 类似于data概念
  computed: {
    MovieInfo() {
      return this.$store.state.movieInfo
    },
    updateMovieDisabled() {
      return (
        this.MovieInfo.MovieName == '' ||
        this.MovieInfo.Director == '' ||
        this.MovieInfo.Starring == '' ||
        this.MovieInfo.Region == '' ||
        this.MovieInfo.Introduction == ''
      )
    },
  },
  //监控data中的数据变化
  watch: {
    newMovieTypes() {
      this.MovieInfo.MovieTypes = this.newMovieTypes.join(' ')
    },
    MovieInfo() {
      this.newMovieTypes = []
      this.MovieTypes = this.MovieInfo.MovieTypes
    },
  },
  //方法集合
  methods: {
    //验证海报是否是图片
    beforePosterUpload(file) {
      const isIMG = file.type === 'image/jpeg' || file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isIMG) {
        this.$message({
          showClose: true, //可关闭
          type: 'error', //提示类型
          message: '电影海报只能是JPG或者PNG格式!',
          center: true, //文字居中
        })
      }
      if (!isLt2M) {
        this.$message({
          showClose: true, //可关闭
          type: 'error', //提示类型
          message: '电影海报大小不能超过2MB!',
          center: true, //文字居中
        })
      }
      return isIMG && isLt2M
    },
    //上传电影海报
    async uploadPoster(file) {
      let result
      var formData = new FormData()
      formData.append('file', file.file)
      result = await reqUploadPoster(formData)
      if (result.code == 0) {
        this.MovieInfo.Poster = result.data
        this.$message({
          showClose: true, //可关闭
          message: '海报上传成功',
          type: 'success', //提示类型
          center: true, //文字居中
        })
      } else {
        this.$message.error({
          showClose: true, //可关闭
          message: result.message,
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
    //提交修改
    async toUpdateMovieInfo() {
      let result
      const MovieID = this.MovieInfo.MovieID
      const MovieName = this.MovieInfo.MovieName
      const Director = this.MovieInfo.Director
      const Starring = this.MovieInfo.Starring
      const Region = this.MovieInfo.Region
      const MovieTypes = this.MovieInfo.MovieTypes
      const MovieLength = this.MovieInfo.MovieLength
      const ReleaseTime = this.MovieInfo.ReleaseTime
      const Status = this.MovieInfo.Status
      const Introduction = this.MovieInfo.Introduction
      result = await reqUpdateMovie({
        MovieID,
        MovieName,
        Director,
        Starring,
        Region,
        MovieTypes,
        MovieLength,
        ReleaseTime,
        Status,
        Introduction,
      })
      if (result.code == 0) {
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
          type: 'error', //提示类型
          center: true, //文字居中
        })
      }
    },
  },
}
</script>
<style>
@import url('../assets/css/movie.css');
</style>
