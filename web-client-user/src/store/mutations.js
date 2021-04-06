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

export default {
  //记录定位到的城市
  // [RECEIVE_CITY_INFO](state, { city }) {
  //   state.city = city
  // },

  //记录所选城市
  [RECEIVE_CITY](state, { city }) {
    state.city = city
  },

  //记录用户信息
  [RECEIVE_USER_INFO](state, { userInfo }) {
    state.userInfo = userInfo
  },

  //重置用户信息
  [RESET_USER_INFO](state) {
    state.userInfo = {}
  },

  //记录所选的电影信息
  [RECEIVE_MOVIE_INFO](state, { movieInfo }) {
    state.movieInfo = movieInfo
  },

  //记录所选的订单信息
  [RECEIVE_ORDER_INFO](state, { orderInfo }) {
    state.orderInfo = orderInfo
  },

  //记录所选的看过信息
  [RECEIVE_SEEN_INFO](state, { seenInfo }) {
    state.seenInfo = seenInfo
  },

  //记录获取到的正在上映的电影信息
  [RECEIVE_MOVIE_LIST](state, { movieList }) {
    state.movieList = movieList
  },

  //记录所选的影院信息
  [RECEIVE_CINEMA_INFO](state, { cinemaInfo }) {
    state.cinemaInfo = cinemaInfo
  },

  //记录上一步的地址
  [RECEIVE_BACK_PATH](state, { backPath }) {
    state.backPath = backPath
  },

  //记录所选的日期
  [RECEIVE_SELECT_DATE](state, { selectDate }) {
    state.selectDate = selectDate
  },

  //记录所选场次
  [RECEIVE_SCREENING](state, { screening }) {
    state.screening = screening
  },

  //记录所选场次的影厅信息
  [RECEIVE_HALL_INFO](state, { hallInfo }) {
    state.hallInfo = hallInfo
  },

  //记录所选场次的座位状态
  [RECEIVE_SEAT_STATUS](state, { seatStatus }) {
    state.seatStatus = seatStatus
  },

  //记录所选座位
  [RECEIVE_SELECT_SEAT](state, { selectSeat }) {
    state.selectSeat = selectSeat
  },
}
