package controller

import (
	"sync"
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//系统管理员添加影院
func addCinema(c *gin.Context) {
	var (
		cinema = &model.CinemaInfo{}
		err    error
	)
	if err = c.ShouldBind(&cinema); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if cinema.CinemaName == "" || cinema.City == "" || cinema.Address == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "影院信息不完整",
		})
		return
	}
	cinema, err = srv.AddCinema(cinema.CinemaName, cinema.City, cinema.Address)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    cinema,
	})
}

//系统管理员删除某些影院
func deleteCinema(c *gin.Context) {
	var (
		v = new(struct {
			CinemaIDs []int `json:"CinemaIDs" form:"CinemaIDs"`
		})
		failID = []int{} //删除失败的影院ID数组
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.CinemaIDs == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.CinemaIDs) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交数据为空",
		})
		return
	}
	for _, v := range v.CinemaIDs {
		wg.Add(1)
		go func(v int) {
			if err = srv.DeleteCinemaInfo(v); err != nil {
				mu.Lock()
				failID = append(failID, v)
				mu.Unlock()
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
		"data":    failID,
	})
}

//系统管理员和管理者修改影院信息
func updateCinema(c *gin.Context) {
	session := sessions.Default(c)
	adminID := session.Get("admin_id")
	if adminID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		return
	}
	var (
		cinema = &model.CinemaInfo{}
		err    error
	)
	if err = c.ShouldBind(&cinema); err != nil || cinema.CinemaID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if cinema.CinemaID < 0 || cinema.CinemaName == "" || cinema.City == "" || cinema.Address == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "修改信息错误",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), cinema.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	cinema, err = srv.UpdateCinemaInfo(cinema.CinemaID, cinema.CinemaName, cinema.City, cinema.Address)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改成功",
		"data":    cinema,
	})
}

//根据条件获取影院列表
func getCinema(c *gin.Context) {
	var (
		v = new(struct {
			City       string `form:"City"`
			CinemaName string `form:"CinemaName"`
			Page       int    `form:"Page"`
			Size       int    `form:"Size"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Page == 0 || v.Size == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.Page < 0 || v.Size < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交信息错误",
		})
		return
	}
	count, cinemas, err := srv.QueryCinemaInfo(v.City, v.CinemaName, v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
	}
	if len(cinemas) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "没有更多了",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"count":   count,
		"data":    cinemas,
	})
}

//获取某个城市的所有影院
func getCinemaByCity(c *gin.Context) {
	var (
		v = new(struct {
			City string `form:"City"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	cinemas, err := srv.QueryCinemaInfoByCity(v.City)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"data":    cinemas,
	})
}

//影院管理员获取管理的影院
func adminManage(c *gin.Context) {
	session := sessions.Default(c)
	adminID := session.Get("admin_id")
	if adminID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) == 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "系统管理员无需获取",
		})
		return
	}
	var (
		v = new(struct {
			City       string `form:"City"`
			CinemaName string `form:"CinemaName"`
			Page       int    `form:"Page"`
			Size       int    `form:"Size"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Page == 0 || v.Size == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.Page < 0 || v.Size < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	count, cinemas, err := srv.QueryManageCinema(adminID.(int), v.Page, v.Size, v.City, v.CinemaName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if len(cinemas) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "没有更多了",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"count":   count,
		"data":    cinemas,
	})
}
