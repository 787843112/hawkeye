package controller

import (
	"sync"
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//添加某些座位
func addSeats(c *gin.Context) {
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
			CinemaID int               `json:"CinemaID" form:"CinemaID"`
			Seats    []*model.SeatInfo `json:"Seats" form:"Seats"`
		})
		sum = 0 //添加失败的个数
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Seats == nil || v.CinemaID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.Seats) == 0 {
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
	if srv.QueryScreeningByHallID(v.Seats[0].HallID) != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	for _, v := range v.Seats {
		wg.Add(1)
		go func(v *model.SeatInfo) {
			if err = srv.QuerySeat(v); err == nil {
				mu.Lock()
				sum++
				mu.Unlock()
			} else {
				if err = srv.AddSeat(v); err != nil {
					mu.Lock()
					sum++
					mu.Unlock()
				}
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    sum,
	})
}

//删除某些座位
func deleteSeats(c *gin.Context) {
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
			CinemaID int               `json:"CinemaID" form:"CinemaID"`
			Seats    []*model.SeatInfo `json:"Seats" form:"Seats"`
		})
		sum = 0 //删除失败的个数
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Seats == nil || v.CinemaID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.Seats) == 0 {
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
	if srv.QueryScreeningByHallID(v.Seats[0].HallID) != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	for _, v := range v.Seats {
		wg.Add(1)
		go func(v *model.SeatInfo) {
			if err = srv.QuerySeat(v); err != nil {
				mu.Lock()
				sum++
				mu.Unlock()
			} else {
				if err = srv.DeleteSeatInfo(v); err != nil {
					mu.Lock()
					sum++
					mu.Unlock()
				}
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
		"data":    sum,
	})
}

//查询座位信息
func getSeat(c *gin.Context) {
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
			CinemaID int `form:"CinemaID"`
			HallID   int `form:"HallID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.HallID == 0 || v.CinemaID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.HallID < 0 || v.CinemaID < 0 {
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
	seats, err := srv.QuerySeatInfo(v.HallID)
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
		"data":    seats,
	})
}
