package controller

import (
	"sync"
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//添加影厅
func addHall(c *gin.Context) {
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
		hall = &model.HallInfo{}
		err  error
	)
	if err = c.ShouldBind(&hall); err != nil || hall.CinemaID == 0 || hall.HallRow == 0 || hall.HallColumn == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if hall.HallName == "" || hall.CinemaID < 0 || hall.HallRow < 0 || hall.HallColumn < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "影厅信息不完整",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), hall.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	hall, err = srv.AddHall(hall.HallName, hall.CinemaID, hall.HallRow, hall.HallColumn)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    hall,
	})
}

//删除某些影厅
func deleteHall(c *gin.Context) {
	//权限校验
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
		v = new(struct {
			CinemaID int   `json:"CinemaID" form:"CinemaID"`
			HallIDs  []int `json:"HallIDs" form:"HallIDs"`
		})
		failID = []int{} //删除失败的影厅ID数组
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.HallIDs == nil || v.CinemaID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.HallIDs) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交数据为空",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), v.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	for _, v := range v.HallIDs {
		wg.Add(1)
		go func(v int) {
			if err = srv.DeleteHallInfo(v); err != nil {
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

//修改影厅
func updateHall(c *gin.Context) {
	//权限校验
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
		hall = &model.HallInfo{}
		err  error
	)
	if err = c.ShouldBind(&hall); err != nil || hall.HallID == 0 || hall.CinemaID == 0 || hall.HallRow == 0 || hall.HallColumn == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if hall.HallName == "" || hall.HallID < 0 || hall.CinemaID < 0 || hall.HallRow < 0 || hall.HallColumn < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), hall.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	hall, err = srv.UpdateHallInfo(hall.HallName, hall.HallID, hall.HallRow, hall.HallColumn)
	if hall == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": err.Error(),
		"data":    hall,
	})
}

//查询影厅
func getHall(c *gin.Context) {
	//权限校验
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
		v = new(struct {
			HallName string `form:"HallName"`
			CinemaID int    `form:"CinemaID"`
			Page     int    `form:"Page"`
			Size     int    `form:"Size"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.CinemaID == 0 || v.Page == 0 || v.Size == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.CinemaID < 0 || v.Page < 0 || v.Size < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), v.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	count, halls, err := srv.QueryHallInfo(v.HallName, v.CinemaID, v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(halls) == 0 {
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
		"data":    halls,
	})
}

//获取某个影厅的信息
func getHallByID(c *gin.Context) {
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
		v = new(struct {
			CinemaID int `form:"CinemaID"`
			HallID   int `form:"HallID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.CinemaID < 1 || v.HallID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	adminType := session.Get("admin_type")
	if adminType.(int) != 1 {
		if srv.FindManageInfo(adminID.(int), v.CinemaID) != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "此影院不是你管理",
			})
			return
		}
	}
	hall, err := srv.QueryHallByHallID(v.HallID)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"data":    hall,
	})
}
