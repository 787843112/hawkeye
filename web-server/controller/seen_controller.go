package controller

import (
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//获取用户看过的电影
func getSeen(c *gin.Context) {
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
	count, seens, err := srv.QueryUserSeen(userID.(int), v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if len(seens) == 0 {
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
		"data":    seens,
	})
}

//根据电影ID获取看过信息
func getSeenByMovieID(c *gin.Context) {
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
			MovieID int `form:"MovieID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.MovieID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	seen, err := srv.QueryUserSeenByMovieID(userID.(int), v.MovieID)
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
		"data":    seen,
	})
}

//删除用户看过的电影
func deleteSeen(c *gin.Context) {
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
			MovieID int `json:"MovieID" form:"MovieID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.MovieID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求体错误",
		})
		return
	}
	if v.MovieID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	if err = srv.DeleteSeenMovie(userID.(int), v.MovieID); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}

//更改评价
func updateEvaluation(c *gin.Context) {
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
		seen = &model.UserSeen{}
		err  error
	)
	if err = c.ShouldBind(&seen); err != nil || seen.MovieID == 0 || seen.Score == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if seen.MovieID < 0 || seen.Score < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	seen, err = srv.UpdateMovieEvaluation(userID.(int), seen.MovieID, seen.Score, seen.Review)
	if seen == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": err.Error(),
		"data":    seen,
	})
}

//删除评价
func deleteEvaluation(c *gin.Context) {
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
			MovieID int `json:"MovieID" form:"MovieID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.MovieID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.MovieID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	seen, err := srv.DeleteMovieEvaluation(userID.(int), v.MovieID)
	if seen == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": err.Error(),
		"data":    seen,
	})
}
