<!-- 电影列表 -->
<template>
  <div id="movielist">
    <el-table
      max-height="600px"
      :row-class-name="tableRowClassName"
      ref="multipleTable"
      @selection-change="handleSelectionChange"
      :data="MovieList"
      tooltip-effect="dark"
      class="movietable"
    >
      <el-table-column align="center" type="selection" width="50"></el-table-column>
      <el-table-column align="center" prop="MovieID" label="ID" width="80"></el-table-column>
      <el-table-column align="center" prop="MovieName" label="名称" width="170" show-overflow-tooltip></el-table-column>
      <el-table-column align="center" prop="Director" label="导演" width="180" show-overflow-tooltip></el-table-column>
      <el-table-column align="center" prop="MovieTypes" label="类型" width="120"></el-table-column>
      <el-table-column align="center" prop="Region" label="地区" width="120"></el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template slot-scope="scope">{{ MovieStatus[scope.row.Status] }}</template></el-table-column
      >
      <el-table-column label="更多" width="150" align="center">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <p>主演: {{ scope.row.Starring }}</p>
            <p>上映时间: {{ scope.row.ReleaseTime }}</p>
            <p>时长: {{ scope.row.MovieLength }}分钟</p>
            <p>评分: {{ scope.row.Score }}分</p>
            <p>票房: {{ scope.row.BoxOffice }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">详细信息</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200" v-if="this.$store.state.adminInfo.Permissions == 1">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-edit" size="mini" @click="toEditMovie(scope.row)">编辑</el-button>
          <el-button size="mini" icon="el-icon-delete-solid" type="danger" @click="toDeleteOneMovie(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
      <el-table-column></el-table-column>
    </el-table>
    <div class="footer">
      <el-button
        v-if="this.$store.state.adminInfo.Permissions == 1"
        @click="toDeleteSomeMovie"
        :disabled="this.MovieIDs.length == 0"
        type="danger"
        icon="el-icon-delete"
        style="float: left;margin-top: 10px;margin-left: 10px;"
        >批量删除</el-button
      >
      <el-pagination
        style="margin-left: 200px;"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[5, 10, 20, 30, 50, 100]"
        :page-size="this.Size"
        :current-page="this.Page"
        layout="total,sizes,prev,pager,next,jumper"
        :total="count"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { reqGetMovies, reqDeleteMovies } from '../api'
export default {
  name: 'MovieList',
  props: ['MovieName', 'MovieTypes', 'Status'],
  data() {
    //这里存放数据
    return {
      Page: 1,
      Size: 10,
      count: 0,
      MovieStatus: ['全部', '已下架', '正在上映', '即将上映'],
      MovieIDs: [],
      MovieList: [],
    }
  },
  //监听属性 类似于data概念
  computed: {},
  //监控data中的数据变化
  watch: {
    MovieName() {
      this.getMovieList()
    },
    MovieTypes() {
      this.getMovieList()
    },
    Status() {
      this.getMovieList()
    },
    Page() {
      this.getMovieList()
    },
    Size() {
      this.getMovieList()
    },
  },
  //方法集合
  methods: {
    //改变每页显示条数
    handleSizeChange(val) {
      this.Size = val
    },
    //改变当前页
    handleCurrentChange(val) {
      this.Page = val
    },
    //去修改电影信息
    toEditMovie(movie) {
      this.$store.dispatch('recordMovieInfo', movie)
      this.$emit('toUpdateMovie')
    },
    //多选发生改变时
    handleSelectionChange(val) {
      this.MovieIDs = val.map((item) => {
        return item.MovieID
      })
    },
    //清空选择
    toggleSelection(rows) {
      if (rows) {
        rows.forEach((row) => {
          this.$refs.multipleTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleTable.clearSelection()
      }
    },
    //获取电影列表
    async getMovieList() {
      let result
      const { MovieName, MovieTypes, Status, Page, Size } = this
      result = await reqGetMovies({ MovieName, MovieTypes, Status, Page, Size })
      if (result.code == 0) {
        this.count = result.count
        this.MovieList = result.data
      } else {
        this.count = 0
        this.Page = 1
        this.MovieList = []
        this.$message({
          showClose: true, //可关闭
          message: result.message,
          type: 'warning', //提示类型
          center: true, //文字居中
        })
      }
    },
    //删除某个电影
    toDeleteOneMovie(movie) {
      const { MovieIDs } = this
      MovieIDs[0] = movie.MovieID
      this.$confirm('此操作将永久删除该电影, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteMovies({ MovieIDs })
          if (result.code == 0) {
            this.getMovieList()
            this.MovieIDs = []
            this.$message({
              type: 'success',
              message: '删除成功!',
            })
          } else {
            this.MovieIDs = []
            this.$message({
              type: 'error',
              message: result.message,
            })
          }
        })
        .catch(() => {
          this.toggleSelection()
          this.MovieIDs = []
        })
    },
    //批量删除某些电影
    toDeleteSomeMovie() {
      const { MovieIDs } = this
      this.$confirm('此操作将永久删除已选电影, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          let result
          result = await reqDeleteMovies({ MovieIDs })
          if (result.code == 0) {
            this.MovieIDs = []
            this.getMovieList()
            this.$message({
              type: 'success',
              message: '删除成功!',
            })
          } else {
            this.MovieIDs = []
            this.$message({
              type: 'error',
              message: result.message,
            })
          }
        })
        .catch(() => {
          this.toggleSelection()
          this.MovieIDs = []
        })
    },
    //修改列表样式
    tableRowClassName({ row }) {
      if (row.Status === 1) {
        return 'shelves-row'
      } else if (row.Status === 2) {
        return 'release-row'
      } else if (row.Status === 3) {
        return 'upcoming-row'
      }
      return ''
    },
  },
  //生命周期 - 创建完成（可以访问当前this实例）
  created() {
    this.getMovieList()
  },
}
</script>
<style>
@import url('../assets/css/movie.css');
</style>
