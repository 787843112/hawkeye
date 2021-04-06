/*
  与后端交互模块
*/
import ajax from './ajax'

//BASE_URL 后端 IP地址+映射的端口号，可前往修改
const BASE_URL = '/api'

//登录
export const reqLogin = ({ AdminName, Password }) => ajax(BASE_URL + '/admin/login', { AdminName, Password }, 'POST')

//自动登录
export const reqAutoLogin = () => ajax(BASE_URL + '/admin/autologin')

//注册
export const reqRegister = ({ AdminName, Password, Phone }) =>
  ajax(BASE_URL + '/admin/register', { AdminName, Password, Phone }, 'POST')

//注销
export const reqLogout = () => ajax(BASE_URL + '/admin/logout')

//修改信息
export const reqUpdate = ({ Password, Phone }) => ajax(BASE_URL + '/admin/update', { Password, Phone }, 'POST')

//把影院管理员设置为系统管理员
export const reqPermissions = ({ AdminID }) => ajax(BASE_URL + '/admin/permissions', { AdminID }, 'POST')

//上传头像
export const reqUploadAvatar = (file) =>
  ajax(BASE_URL + '/admin/upload', file, 'POST', {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })

//获取影院管理员
export const reqGetCinemaAdmins = ({ Page, Size }) => ajax(BASE_URL + '/admin/get', { Page, Size })

//删除某些管理员
export const reqDeleteAdmins = ({ AdminIDs }) => ajax(BASE_URL + '/admin/delete', { AdminIDs }, 'POST')

//系统管理员获取影院管理员管理的影院
export const reqGetManage = ({ AdminID, Page, Size }) => ajax(BASE_URL + '/manage/get', { AdminID, Page, Size })

//获取某城市的所有影院
export const reqGetCityCinemas = ({ City }) => ajax(BASE_URL + '/cinema/getbycity', { City })

//系统管理员给影院管理员添加管理的影院
export const reqAddManage = ({ AdminID, CinemaID }) => ajax(BASE_URL + '/manage/add', { AdminID, CinemaID }, 'POST')

//系统管理员删除某个影院管理员管理的某些影院
export const reqDeleteManage = ({ AdminID, CinemaIDs }) =>
  ajax(BASE_URL + '/manage/delete', { AdminID, CinemaIDs }, 'POST')

//上传电影海报
export const reqUploadPoster = (file) =>
  ajax(BASE_URL + '/movie/upload', file, 'POST', {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })

//添加电影
export const reqAddMovie = ({
  MovieName,
  Director,
  Starring,
  Region,
  MovieTypes,
  MovieLength,
  ReleaseTime,
  Status,
  Introduction,
}) =>
  ajax(
    BASE_URL + '/movie/add',
    {
      MovieName,
      Director,
      Starring,
      Region,
      MovieTypes,
      MovieLength,
      ReleaseTime,
      Status,
      Introduction,
    },
    'POST'
  )

//添加电影类型
export const reqAddMovieType = ({ TypeName }) => ajax(BASE_URL + '/movie/addtype', { TypeName }, 'POST')

//删除某些电影
export const reqDeleteMovies = ({ MovieIDs }) => ajax(BASE_URL + '/movie/delete', { MovieIDs }, 'POST')

//修改电影信息
export const reqUpdateMovie = ({
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
}) =>
  ajax(
    BASE_URL + '/movie/update',
    {
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
    },
    'POST'
  )

//根据条件获取电影列表
export const reqGetMovies = ({ MovieName, MovieTypes, Status, Page, Size }) =>
  ajax(BASE_URL + '/movie/get', { MovieName, MovieTypes, Status, Page, Size })

//获取所有正在上映的电影列表
export const reqReleasedMovies = () => ajax(BASE_URL + '/movie/getreleased')

//获取所有的电影类型
export const reqGetAllMovieTypes = () => ajax(BASE_URL + '/movie/gettype')

//添加影院
export const reqAddCinema = ({ CinemaName, City, Address }) =>
  ajax(BASE_URL + '/cinema/add', { CinemaName, City, Address }, 'POST')

//删除某些影院
export const reqDeleteCinemas = ({ CinemaIDs }) => ajax(BASE_URL + '/cinema/delete', { CinemaIDs }, 'POST')

//修改影院信息
export const reqUpdateCinema = ({ CinemaID, CinemaName, City, Address }) =>
  ajax(BASE_URL + '/cinema/update', { CinemaID, CinemaName, City, Address }, 'POST')

//根据条件获取影院列表
export const reqGetCinemas = ({ CinemaName, City, Page, Size }) =>
  ajax(BASE_URL + '/cinema/get', { CinemaName, City, Page, Size })

//影院管理员获取管理的影院
export const reqGetAdminManage = ({ City, CinemaName, Page, Size }) =>
  ajax(BASE_URL + '/cinema/getmanage', { City, CinemaName, Page, Size })

//添加影厅
export const reqAddHall = ({ HallName, CinemaID, HallRow, HallColumn }) =>
  ajax(BASE_URL + '/hall/add', { HallName, CinemaID, HallRow, HallColumn }, 'POST')

//删除某些影厅
export const reqDeleteHalls = ({ CinemaID, HallIDs }) => ajax(BASE_URL + '/hall/delete', { CinemaID, HallIDs }, 'POST')

//修改影厅信息
export const reqUpdateHall = ({ HallID, CinemaID, HallName, HallRow, HallColumn }) =>
  ajax(BASE_URL + '/hall/update', { HallID, CinemaID, HallName, HallRow, HallColumn }, 'POST')

//查询影厅列表
export const reqGetHalls = ({ HallName, CinemaID, Page, Size }) =>
  ajax(BASE_URL + '/hall/get', { HallName, CinemaID, Page, Size })

//查询某个影厅的信息
export const reqGetHall = ({ CinemaID, HallID }) => ajax(BASE_URL + '/hall/getbyid', { CinemaID, HallID })

//添加某些座位
export const reqAddSeats = ({ CinemaID, Seats }) => ajax(BASE_URL + '/seat/add', { CinemaID, Seats }, 'POST')

//删除某些座位
export const reqDeleteSeats = ({ CinemaID, Seats }) => ajax(BASE_URL + '/seat/delete', { CinemaID, Seats }, 'POST')

//获取某个影厅的座位信息
export const reqGetSeats = ({ CinemaID, HallID }) => ajax(BASE_URL + '/seat/get', { CinemaID, HallID })

//添加场次
export const reqAddScreening = ({ CinemaID, HallID, MovieID, StartTime, EndTime, Price }) =>
  ajax(BASE_URL + '/screening/add', { CinemaID, HallID, MovieID, StartTime, EndTime, Price }, 'POST')

//删除某些场次
export const reqDeleteScreenings = ({ CinemaID, ScreeningIDs }) =>
  ajax(BASE_URL + '/screening/delete', { CinemaID, ScreeningIDs }, 'POST')

//获取场次列表
export const reqGetScreenings = ({ CinemaID, HallID, StartDate, EndDate, Page, Size }) =>
  ajax(BASE_URL + '/screening/adminget', {
    CinemaID,
    HallID,
    StartDate,
    EndDate,
    Page,
    Size,
  })
