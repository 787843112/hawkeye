/**
 * 获取state中的各种数据
 */

export default {
  getAdminInfo: (state) => state.adminInfo,
  getMovieInfo: (state) => state.movieInfo,
  getCinemaInfo: (state) => state.cinemaInfo,
  getHallInfo: (state) => state.hallInfo,
  getAllMovieTypes: (state) => state.allMovieTypes,
}
