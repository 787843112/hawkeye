<!-- 电影页面 -->
<template>
  <div id="movie">
    <div class="movie-div">
      <!-- 头部 -->
      <div class="movie-header">
        <el-input
          clearable
          placeholder="按电影名称搜索"
          prefix-icon="el-icon-search"
          v-model="MovieName"
          class="search-movie-name"
        ></el-input>
        <el-select
          v-model="searchMovieTypes"
          multiple
          :multiple-limit="3"
          placeholder="请选择类型"
          class="select-movie-type"
        >
          <el-option v-for="(item, index) in allMovieTypes" :key="index" :label="item.TypeName" :value="item.TypeName">
          </el-option>
        </el-select>
        <el-select v-model="Status" class="selsect-movie-status">
          <el-option v-for="(item, index) in MovieStatus" :key="index" :label="item" :value="index"> </el-option>
        </el-select>
        <el-button
          v-if="this.$store.state.adminInfo.Permissions == 1"
          @click="toAddMovie"
          class="add-movie-button"
          type="primary"
          round
          icon="el-icon-plus"
          >添加电影</el-button
        >
        <el-button
          v-if="this.$store.state.adminInfo.Permissions == 1"
          @click="toAddMovieType"
          class="add-type-button"
          type="success"
          icon="el-icon-circle-plus"
          >添加类型</el-button
        >
      </div>
      <!-- 主体 -->
      <div class="movie-body">
        <el-tabs type="border-card" v-model="editableTabsValue">
          <el-tab-pane>
            <span slot="label" @click="clickMovieList"><i class="el-icon-film"></i> 电影列表</span>
            <!-- 电影列表组件 -->
            <MovieList
              ref="MovieList"
              :MovieName="MovieName"
              :MovieTypes="MovieTypes"
              :Status="Status"
              @toUpdateMovie="toUpdateMovie"
            />
          </el-tab-pane>
          <el-tab-pane
            :disabled="this.$store.state.movieInfo.MovieID == null"
            v-if="this.$store.state.adminInfo.Permissions == 1"
          >
            <span slot="label"
              ><i :class="this.$store.state.movieInfo.MovieID == null ? 'el-icon-lock' : 'el-icon-edit'"></i>
              {{
                this.$store.state.movieInfo.MovieID == null
                  ? ' 修改电影信息'
                  : ' 电影ID:' + this.$store.state.movieInfo.MovieID
              }}</span
            >
            <!-- 更新电影信息组件 -->
            <UpdateMovie />
          </el-tab-pane>
          <el-tab-pane v-if="this.$store.state.adminInfo.Permissions == 1">
            <span slot="label"><i class="el-icon-circle-plus"></i> 添加电影</span>
            <!-- 添加电影组件 -->
            <AddMovie />
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>

<script>
import MovieList from '../components/MovieList'
import UpdateMovie from '../components/UpdateMovie'
import AddMovie from '../components/AddMovie'
import { reqAddMovieType } from '../api'
import { mapActions } from 'vuex'
export default {
  name: 'Movie',
  components: { MovieList, UpdateMovie, AddMovie },
  data() {
    //这里存放数据
    return {
      MovieName: '',
      MovieTypes: '',
      Status: 0,
      searchMovieTypes: [],
      MovieStatus: ['全部', '已下架', '正在上映', '即将上映'],
      editableTabsValue: '0',
    }
  },
  //监听属性 类似于data概念
  computed: {
    allMovieTypes() {
      return this.$store.state.allMovieTypes
    },
  },
  //监控data中的数据变化
  watch: {
    $route(to) {
      if (to.path == '/movie') {
        this.$refs.MovieList.getMovieList()
      }
    },
    MovieName() {
      this.$refs.MovieList.getMovieList()
      this.editableTabsValue = '0'
    },
    MovieTypes() {
      this.$refs.MovieList.getMovieList()
      this.editableTabsValue = '0'
    },
    Status() {
      this.$refs.MovieList.getMovieList()
      this.editableTabsValue = '0'
    },
    searchMovieTypes() {
      this.MovieTypes = this.searchMovieTypes.join(' ')
    },
  },
  //方法集合
  methods: {
    clickMovieList() {
      this.$refs.MovieList.getMovieList()
    },
    ...mapActions(['toGetAllMovieTypes']),
    //去修改电影信息
    toUpdateMovie() {
      this.editableTabsValue = '1'
    },
    //去添加电影
    toAddMovie() {
      this.editableTabsValue = '2'
    },
    //添加类型
    toAddMovieType() {
      this.$prompt('请输入类型名称', {
        confirmButtonText: '保存',
        cancelButtonText: '取消',
        inputPattern: /^[\u4e00-\u9fa5]{2,2}$/,
        inputErrorMessage: '类型名称只能是2个中文',
      })
        .then(async ({ value }) => {
          const TypeName = value
          let result = await reqAddMovieType({ TypeName })
          if (result.code == 0) {
            this.$message({
              showClose: true, //可关闭
              message: '添加成功',
              type: 'success', //提示类型
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
  },
  //生命周期 - 挂载完成（可以访问DOM元素）
  mounted() {
    this.toGetAllMovieTypes()
  },
}
</script>
<style>
@import url('../assets/css/movie.css');
</style>
