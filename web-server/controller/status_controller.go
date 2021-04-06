package controller

import (
	"strconv"
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	orderSeat string
	state     []*model.StatusInfo
)

//更新座位状态
func updateStatus(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		return
	}
	var (
		v = new(struct {
			Status []*model.StatusInfo `json:"Status" form:"Status"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.Status == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.Status) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交信息为空",
		})
		return
	}
	orderSeat = ""
	for _, v := range v.Status {
		if srv.UpdateStatusInfo(v) != nil {
			srv.RestoreStatusInfo(state)
			c.JSON(200, gin.H{
				"code":    1,
				"message": "修改失败",
			})
			orderSeat = ""
			return
		}
		state = append(state, v)
		orderSeat += strconv.Itoa(v.SeatRow) + "排" + strconv.Itoa(v.SeatColumn) + "座 "
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改成功",
	})
}

//查询座位状态
func getStatus(c *gin.Context) {
	var (
		v = new(struct {
			ScreeningID int `form:"ScreeningID"`
		})
		status []*model.StatusInfo
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.ScreeningID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	hall, status, err := srv.QueryStatus(v.ScreeningID)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(status) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "没有该场次",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"hall":    hall,
		"data":    status,
	})
}
