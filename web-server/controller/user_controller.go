package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"web-server/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//用户登录
func userLogin(c *gin.Context) {
	var (
		v = new(struct {
			UserName string `json:"UserName" form:"UserName"`
			Password string `json:"Password" form:"Password"`
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
	if v.UserName == "" || v.Password == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户名或密码为空",
		})
		return
	}
	user, err := srv.QueryUserInfoByUserName(v.UserName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户不存在",
		})
		return
	}
	v.Password = util.MD5V([]byte(v.Password))
	if v.Password != user.Password {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.UserID)
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    user,
	})
}

//用户自动登录
func userAutoLogin(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "登录信息已过期",
		})
		return
	}
	user, err := srv.QueryUserInfoByUserID(userID.(int))
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户已不存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "自动登录成功",
		"data":    user,
	})
}

//用户注销
func userLogout(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		return
	}
	session.Delete("user_id")
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注销成功",
	})
}

//用户注册
func userRegister(c *gin.Context) {
	var (
		v = new(struct {
			UserName string `json:"UserName" form:"UserName"`
			Password string `json:"Password" form:"Password"`
			Phone    string `json:"Phone" form:"Phone"`
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
	if v.UserName == "" || v.Password == "" || v.Phone == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户信息不完整",
		})
		return
	}
	if _, err = srv.QueryUserInfoByUserName(v.UserName); err == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户名已被注册",
		})
		return
	}
	v.Password = util.MD5V([]byte(v.Password))
	user, err := srv.RegisterUser(v.UserName, v.Password, v.Phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册失败",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.UserID)
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注册成功",
		"data":    user,
	})
}

//修改用户信息
func userUpdate(c *gin.Context) {
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
			Nickname string `json:"Nickname" form:"Nickname"`
			Password string `json:"Password" form:"Password"`
			Phone    string `json:"Phone" form:"Phone"`
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
	if v.Nickname == "" && v.Password == "" && v.Phone == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "至少修改一项",
		})
		return
	}
	user, err := srv.UpdateUserInfo(userID.(int), v.Nickname, v.Password, v.Phone)
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
		"data":    user,
	})
}

//修改用户头像
func updateUserAvatar(c *gin.Context) {
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
		path     = "../static/avatars/"
		avatar   = ""
		fileType string
		err      error
	)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	for k, v := range file.Filename {
		if v == '.' {
			fileType = file.Filename[k:]
		}
	}
	if fileType == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "无效的图片文件类型",
		})
		return
	}
	files, err := file.Open()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "解析文件失败",
		})
		return
	}
	defer files.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, files); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "生成文件名失败",
		})
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	fileName := hex.EncodeToString(hashInBytes)
	if err = c.SaveUploadedFile(file, path+fileName+fileType); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "保存文件失败",
		})
		return
	}
	avatar = fileName + fileType
	user, err := srv.UpdateUserAvatar(userID.(int), avatar)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "修改头像失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改头像成功",
		"data":    user,
	})
}
