/**
 * 获取state中的各种数据
 */

export default {
  getCity: (state) => state.city,
  getUserInfo: (state) => state.userInfo,
  getMovieInfo: (state) => state.movieInfo,
  getMovieList: (state) => state.movieList,
  getCinemaInfo: (state) => state.cinemaInfo,
  getBackPath: (state) => state.backPath,
  getSelectDate: (state) => state.selectDate,
  getSelectSeat: (state) => state.selectSeat,
}
