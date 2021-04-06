import {
  RECEIVE_BACK_PATH,
  RECEIVE_CINEMA_INFO,
  RECEIVE_CITY,
  RECEIVE_MOVIE_INFO,
  RECEIVE_MOVIE_LIST,
  RECEIVE_USER_INFO,
  RESET_USER_INFO,
  RECEIVE_SELECT_DATE,
  RECEIVE_SCREENING,
  RECEIVE_HALL_INFO,
  RECEIVE_SEAT_STATUS,
  RECEIVE_SELECT_SEAT,
  RECEIVE_ORDER_INFO,
  RECEIVE_SEEN_INFO,
} from './mutation-types'
import { reqLogout, reqAutoLogin, reqReleaseMovies } from '../api'

export default {
  //同步记录所选城市
  recordCity({ commit }, city) {
    commit(RECEIVE_CITY, { city })
  },

  //同步记录用户信息
  recordUserInfo({ commit }, userInfo) {
    commit(RECEIVE_USER_INFO, { userInfo })
  },

  //异步获取用户信息
  async toGetUserInfo({ commit }) {
    const result = await reqAutoLogin()
    if (result.code == 0) {
      const userInfo = result.data
      commit(RECEIVE_USER_INFO, { userInfo })
    }
  },

  //异步获取正在上映的电影
  async toGetMovieList({ commit }) {
    const result = await reqReleaseMovies()
    if (result.code == 0) {
      const movieList = result.data
      commit(RECEIVE_MOVIE_LIST, { movieList })
    }
  },

  //同步记录所选的电影信息
  recordMovieInfo({ commit }, movieInfo) {
    commit(RECEIVE_MOVIE_INFO, { movieInfo })
  },

  // //同步记录获取到的正在上映的电影信息
  // recordMovieList({ commit }, movieList) {
  //   commit(RECEIVE_MOVIE_LIST, { movieList })
  // },

  //同步记录所选的影院信息
  recordCinemaInfo({ commit }, cinemaInfo) {
    commit(RECEIVE_CINEMA_INFO, { cinemaInfo })
  },

  //异步退出登录
  async logout({ commit }) {
    const result = await reqLogout()
    if (result.code == 0) {
      commit(RESET_USER_INFO)
    }
  },

  //同步记录所选城市的影院列表
  // recordCinemaList({ commit }, cinemaList) {
  //   commit(RECEIVE_CINEMA_LIST, { cinemaList })
  // },

  //记录上一步的地址
  recordBackPath({ commit }, backPath) {
    commit(RECEIVE_BACK_PATH, { backPath })
  },

  //同步记录所选日期
  recordSelectDate({ commit }, selectDate) {
    commit(RECEIVE_SELECT_DATE, { selectDate })
  },

  //同步记录所选场次
  recordSelectScreening({ commit }, screening) {
    commit(RECEIVE_SCREENING, { screening })
  },

  //同步记录所选场次影厅信息
  recordHallInfo({ commit }, hallInfo) {
    commit(RECEIVE_HALL_INFO, { hallInfo })
  },

  //同步记录所选订单信息
  recordOrderInfo({ commit }, orderInfo) {
    commit(RECEIVE_ORDER_INFO, { orderInfo })
  },

  //同步记录所选看过信息
  recordSeenInfo({ commit }, seenInfo) {
    commit(RECEIVE_SEEN_INFO, { seenInfo })
  },

  //同步记录所选场次座位状态
  recordSeatStatus({ commit }, seatStatus) {
    commit(RECEIVE_SEAT_STATUS, { seatStatus })
  },

  //记录所选座位
  recordSelectSeat({ commit }, selectSeat) {
    commit(RECEIVE_SELECT_SEAT, { selectSeat })
  },

  //异步获取当前城市
  // async toGetCity({ commit }) {
  //   var _this = this
  //   if (navigator.geolocation) {
  //     navigator.geolocation.getCurrentPosition(
  //       //locationSuccess 获取成功的话
  //       function(position) {
  //         _this.getLongitude = position.coords.longitude
  //         _this.getLatitude = position.coords.latitude
  //         // alert(_this.getLongitude) //弹出经度测试
  //         // alert(_this.getLatitude)
  //       }
  //       //locationError  获取失败的话
  //       // function(error) {
  //       //   var errorType = [
  //       //     '您拒绝共享位置信息',
  //       //     '获取不到位置信息',
  //       //     '获取位置信息超时',
  //       //   ]
  //       //   alert(errorType[error.code - 1])
  //       // }
  //     )
  //   }
  //   const key = 'XMNBZ-VBURK-IL2JC-AGRAW-4Y6I5-66BWH'
  //   const location = _this.getLatitude + ',' + _this.getLatitude
  //   const result = await reqCityInfo({ location, key })
  //   if (result.status == 0) {
  //     const city = result.result.ad_info.city
  //     commit(RECEIVE_CITY_INFO, { city })
  //   } else {
  //     const city = '北京市'
  //     commit(RECEIVE_CITY_INFO, { city })
  //   }
  // },
}
