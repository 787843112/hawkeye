package controller

import (
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//生成订单
func addOrder(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		srv.RestoreStatusInfo(state)
		state = nil
		orderSeat = ""
		return
	}
	if orderSeat == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未选择座位",
		})
		srv.RestoreStatusInfo(state)
		state = nil
		orderSeat = ""
		return
	}
	var (
		order = &model.UserOrder{}
		err   error
	)
	if err = c.ShouldBind(&order); err != nil || order.MovieID == 0 || order.Number == 0 || order.Price == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		srv.RestoreStatusInfo(state)
		state = nil
		orderSeat = ""
		return
	}
	if order.MovieID < 0 || order.Number < 0 || order.Price < 0 ||
		order.MovieName == "" || order.Date == "" || order.StartTime == "" ||
		order.EndTime == "" || order.CinemaName == "" || order.Address == "" ||
		order.HallName == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "订单信息不完整",
		})
		srv.RestoreStatusInfo(state)
		state = nil
		orderSeat = ""
		return
	}
	order, err = srv.AddUserOrder(userID.(int), order.MovieID, order.Number, order.Price, order.MovieName, order.Date, order.StartTime, order.EndTime, order.CinemaName, order.Address, order.HallName, orderSeat)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "下单失败",
		})
		srv.RestoreStatusInfo(state)
		state = nil
		orderSeat = ""
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "下单成功",
		"data":    order,
	})
	state = nil
	orderSeat = ""
}

//删除订单
func deleteOrder(c *gin.Context) {
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
			OrderID int `json:"OrderID" form:"OrderID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.OrderID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.OrderID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	if err = srv.DeleteUserOrder(v.OrderID, userID.(int)); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"success": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"success": "删除成功",
	})
}

//获取用户的订单
func getOrder(c *gin.Context) {
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
			Page int `form:"Page"`
			Size int `form:"Size"`
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
	count, orders, err := srv.QueryUserOrder(userID.(int), v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if len(orders) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "没有更多了",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"success": "查询成功",
		"count":   count,
		"data":    orders,
	})
}
