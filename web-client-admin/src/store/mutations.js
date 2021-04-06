import {
  RECEIVE_ADMIN_INFO,
  RESET_ADMIN_INFO,
  RECEIVE_MOVIE_INFO,
  RECEIVE_CINEMA_INFO,
  RECEIVE_HALL_INFO,
  RECEIVE_ALL_MOVIE_TYPES,
} from './mutation-types'
export default {
  //记录管理员信息
  [RECEIVE_ADMIN_INFO](state, { adminInfo }) {
    state.adminInfo = adminInfo
  },

  //重置管理员信息
  [RESET_ADMIN_INFO](state) {
    state.adminInfo = {}
  },

  //记录电影信息
  [RECEIVE_MOVIE_INFO](state, { movieInfo }) {
    state.movieInfo = movieInfo
  },

  //记录影院信息
  [RECEIVE_CINEMA_INFO](state, { cinemaInfo }) {
    state.cinemaInfo = cinemaInfo
  },

  //记录影厅信息
  [RECEIVE_HALL_INFO](state, { hallInfo }) {
    state.hallInfo = hallInfo
  },

  //记录所有的电影类型
  [RECEIVE_ALL_MOVIE_TYPES](state, { allMovieTypes }) {
    state.allMovieTypes = allMovieTypes
  },
}
