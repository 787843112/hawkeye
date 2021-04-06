import {
  RECEIVE_ADMIN_INFO,
  RESET_ADMIN_INFO,
  RECEIVE_MOVIE_INFO,
  RECEIVE_CINEMA_INFO,
  RECEIVE_HALL_INFO,
  RECEIVE_ALL_MOVIE_TYPES,
} from './mutation-types'
import { reqAutoLogin, reqLogout, reqGetAllMovieTypes } from '../api'

export default {
  //同步记录管理员信息
  recordAdminInfo({ commit }, adminInfo) {
    commit(RECEIVE_ADMIN_INFO, { adminInfo })
  },

  //异步获取管理员信息
  async toGetAdminInfo({ commit }) {
    const result = await reqAutoLogin()
    if (result.code == 0) {
      const adminInfo = result.data
      commit(RECEIVE_ADMIN_INFO, { adminInfo })
    }
  },

  //异步获取所有的电影类型
  async toGetAllMovieTypes({ commit }) {
    const result = await reqGetAllMovieTypes()
    if (result.code == 0) {
      const allMovieTypes = result.data
      commit(RECEIVE_ALL_MOVIE_TYPES, { allMovieTypes })
    }
  },

  //异步退出登录
  async logout({ commit }) {
    const result = await reqLogout()
    if (result.code == 0) {
      commit(RESET_ADMIN_INFO)
    }
  },

  //记录电影信息
  recordMovieInfo({ commit }, movieInfo) {
    commit(RECEIVE_MOVIE_INFO, { movieInfo })
  },

  //记录影院信息
  recordCinemaInfo({ commit }, cinemaInfo) {
    commit(RECEIVE_CINEMA_INFO, { cinemaInfo })
  },

  //记录影厅信息
  recordHallInfo({ commit }, hallInfo) {
    commit(RECEIVE_HALL_INFO, { hallInfo })
  },
}
