package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//VerifyAdmin ...验证是否是系统管理员
func VerifyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		adminID := session.Get("admin_id")
		if adminID == nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "未登录",
			})
			c.Abort()
			return
		}
		adminType := session.Get("admin_type")
		if adminType.(int) != 1 {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "权限不足",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
