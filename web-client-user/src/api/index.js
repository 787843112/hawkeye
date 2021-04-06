/*
  与后端交互模块
**/
import ajax from './ajax'

//BASE_URL 后端 IP地址+映射的端口号，可前往修改
const BASE_URL = '/api'

//定位城市
// export const reqCityInfo = ({ location, key }) =>
//   ajax('https://apis.map.qq.com/ws/geocoder/v1/', { location, key })

//登录
export const reqLogin = ({ UserName, Password }) => ajax(BASE_URL + '/user/login', { UserName, Password }, 'POST')

//自动登录
export const reqAutoLogin = () => ajax(BASE_URL + '/user/autologin')

//注册
export const reqRegister = ({ UserName, Password, Phone }) => ajax(BASE_URL + '/user/register', { UserName, Password, Phone }, 'POST')

//注销
export const reqLogout = () => ajax(BASE_URL + '/user/logout')

//修改信息
export const reqUpdate = ({ Nickname, Password, Phone }) => ajax(BASE_URL + '/user/update', { Nickname, Password, Phone }, 'POST')

//上传头像
export const reqUploadAvatar = (file) =>
  ajax(BASE_URL + '/user/upload', file, 'POST', {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })

//获取用户想看的电影
export const reqGetWant = ({ Page, Size }) => ajax(BASE_URL + '/want/get', { Page, Size })

//获取用户所有想看的电影ID
export const reqGetWantMovieID = () => ajax(BASE_URL + '/want/getmovieid')

//添加想看的电影
export const reqAddWant = ({ MovieID }) => ajax(BASE_URL + '/want/add', { MovieID }, 'POST')

//删除用户某个想看的电影
export const reqDeleteWant = ({ MovieID }) => ajax(BASE_URL + '/want/delete', { MovieID }, 'POST')

//获取用户看过的电影
export const reqGetSeen = ({ Page, Size }) => ajax(BASE_URL + '/seen/get', { Page, Size })

//获取某个看过的信息
export const reqGetSeenByMovieID = ({ MovieID }) => ajax(BASE_URL + '/seen/getbymovieid', { MovieID })

//删除用户某个看过的电影
export const reqDeleteSeen = ({ MovieID }) => ajax(BASE_URL + '/seen/delete', { MovieID }, 'POST')

//更新评价
export const reqUpdateEvaluation = ({ MovieID, Score, Review }) => ajax(BASE_URL + '/evaluation/update', { MovieID, Score, Review }, 'POST')

//删除评价
export const reqDeleteEvaluation = ({ MovieID }) => ajax(BASE_URL + '/evaluation/delete', { MovieID }, 'POST')

//生成订单
export const reqAddOrder = ({ MovieID, MovieName, Date, StartTime, EndTime, CinemaName, Address, HallName, Number, Price }) =>
  ajax(
    BASE_URL + '/order/add',
    {
      MovieID,
      MovieName,
      Date,
      StartTime,
      EndTime,
      CinemaName,
      Address,
      HallName,
      Number,
      Price,
    },
    'POST'
  )

//删除订单
export const reqDeleteOrder = ({ OrderID }) => ajax(BASE_URL + '/order/delete', { OrderID }, 'POST')

//获取用户订单
export const reqGetOrders = ({ Page, Size }) => ajax(BASE_URL + '/order/get', { Page, Size })

//根据条件获取电影列表
export const reqGetMovies = ({ MovieName, MovieType, Status, Page, Size }) =>
  ajax(BASE_URL + '/movie/get', {
    MovieName,
    MovieType,
    Status,
    Page,
    Size,
  })

//根据电影ID获取电影信息
export const reqGetMovieInfo = ({ MovieID }) => ajax(BASE_URL + '/movie/getbyid', { MovieID })

//获取所有正在上映的电影
export const reqReleaseMovies = () => ajax(BASE_URL + '/movie/getreleased')

//获取最受期待的10部电影
export const reqExpectMovies = () => ajax(BASE_URL + '/movie/getexpect')

//获取所有的电影类型
export const reqGetAllMovieTypes = () => ajax(BASE_URL + '/movie/gettype')

//根据条件获取影院列表
export const reqGetCinemas = ({ CinemaName, City, Page, Size }) => ajax(BASE_URL + '/cinema/get', { CinemaName, City, Page, Size })

//获取某影院某电影的场次
export const reqGetScreenings = ({ CinemaID, MovieID, SelectDate, Page, Size }) =>
  ajax(BASE_URL + '/screening/userget', {
    CinemaID,
    MovieID,
    SelectDate,
    Page,
    Size,
  })

//获取某场次的座位状态
export const reqGetStatus = ({ ScreeningID }) =>
  ajax(BASE_URL + '/status/get', {
    ScreeningID,
  })

//修改座位状态
export const reqUpdateStatus = ({ Status }) => ajax(BASE_URL + '/status/update', { Status }, 'POST')
