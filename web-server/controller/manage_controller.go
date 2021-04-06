package controller

import (
	"sync"
	"web-server/model"

	"github.com/gin-gonic/gin"
)

//系统管理员给影院管理员添加管理的影院
func addManage(c *gin.Context) {
	var (
		manage = &model.ManageInfo{}
		err    error
	)
	if err = c.ShouldBind(&manage); err != nil || manage.AdminID == 0 || manage.CinemaID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if manage.AdminID < 0 || manage.CinemaID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	cinema, err := srv.AddManage(manage.AdminID, manage.CinemaID)
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
		"data":    cinema,
	})
}

//系统管理员删除某个影院管理员管理的某些影院
func deleteManage(c *gin.Context) {
	var (
		v = new(struct {
			AdminID   int   `json:"AdminID" form:"AdminID"`
			CinemaIDs []int `json:"CinemaIDs" form:"CinemaIDs"`
		})
		failID = []int{} //删除失败的管理的影院ID数组
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.AdminID == 0 || v.CinemaIDs == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.AdminID < 0 || len(v.CinemaIDs) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	for _, cinemaID := range v.CinemaIDs {
		wg.Add(1)
		go func(cinemaID int) {
			if err = srv.DeleteManage(v.AdminID, cinemaID); err != nil {
				mu.Lock()
				failID = append(failID, cinemaID)
				mu.Unlock()
			}
			wg.Done()
		}(cinemaID)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
		"data":    failID,
	})
}

//系统管理员获取影院管理员管理的影院
func getManage(c *gin.Context) {
	var (
		v = new(struct {
			AdminID int `form:"AdminID"`
			Page    int `form:"Page"`
			Size    int `form:"Size"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Page == 0 || v.Size == 0 || v.AdminID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.Page < 0 || v.Size < 0 || v.AdminID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	count, cinemas, err := srv.QueryManageCinema(v.AdminID, v.Page, v.Size, "", "")
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
