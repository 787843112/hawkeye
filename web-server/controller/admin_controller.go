package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"sync"
	"web-server/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//管理员登录
func adminLogin(c *gin.Context) {
	var (
		v = new(struct {
			AdminName string `json:"AdminName" form:"AdminName"`
			Password  string `json:"Password" form:"Password"`
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
	if v.AdminName == "" || v.Password == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "管理员名或密码为空",
		})
		return
	}
	admin, err := srv.QueryAdminInfoByAdminName(v.AdminName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "管理员不存在",
		})
		return
	}
	v.Password = util.MD5V([]byte(v.Password))
	if v.Password != admin.Password {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("admin_id", admin.AdminID)
	session.Set("admin_type", admin.Permissions)
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    admin,
	})
}

//管理员自动登录
func adminAutoLogin(c *gin.Context) {
	session := sessions.Default(c)
	adminID := session.Get("admin_id")
	if adminID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "登录信息已过期",
		})
		return
	}
	admin, err := srv.QueryAdminInfoByAdminID(adminID.(int))
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "管理员不存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    admin,
	})
}

//管理员注销
func adminLogout(c *gin.Context) {
	session := sessions.Default(c)
	adminID := session.Get("admin_id")
	if adminID == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未登录",
		})
		return
	}
	session.Delete("admin_id")
	session.Delete("admin_type")
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注销成功",
	})
}

//注册管理员
func adminRegister(c *gin.Context) {
	var (
		v = new(struct {
			AdminName string `json:"AdminName" form:"AdminName"`
			Password  string `json:"Password" form:"Password"`
			Phone     string `json:"Phone" form:"Phone"`
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
	if v.AdminName == "" || v.Password == "" || v.Phone == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册信息不完整",
		})
		return
	}
	if _, err = srv.QueryAdminInfoByAdminName(v.AdminName); err == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "管理员名已被注册",
		})
		return
	}
	v.Password = util.MD5V([]byte(v.Password))
	admin, err := srv.RegisterAdmin(v.AdminName, v.Password, v.Phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册失败",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("admin_id", admin.AdminID)
	session.Set("admin_type", admin.Permissions)
	session.Save()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注册成功",
		"data":    admin,
	})
}

//更新管理员信息
func adminUpdate(c *gin.Context) {
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
	if v.Password == "" && v.Phone == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "未提交修改内容",
		})
		return
	}
	admin, err := srv.UpdateAdminInfo(adminID.(int), v.Password, v.Phone)
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
		"data":    admin,
	})
}

//将影院管理员设置为系统管理员
func permissions(c *gin.Context) {
	var (
		v = new(struct {
			AdminID int `json:"AdminID" form:"AdminID"`
		})
		err error
	)
	if err = c.ShouldBind(&v); err != nil || v.AdminID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.AdminID < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	if err = srv.UpdateAdminPermissions(v.AdminID); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改成功",
	})
}

//获取所有影院管理员
func getAdmin(c *gin.Context) {
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
	count, admins, err := srv.QueryAllCinemaAdmin(v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(admins) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "没有更多了!",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取成功",
		"count":   count,
		"data":  admins,
	})
}

//删除某些管理员
func deleteAdmin(c *gin.Context) {
	var (
		v = new(struct {
			AdminIDs []int `json:"AdminIDs" form:"AdminIDs"`
		})
		failID = []int{} //删除失败的管理员ID数值
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.AdminIDs == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.AdminIDs) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交数据为空",
		})
		return
	}
	for _, v := range v.AdminIDs {
		wg.Add(1)
		go func(v int) {
			if err = srv.DeleteAdminInfo(v); err != nil {
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

//修改管理员头像
func updateAdminAvatar(c *gin.Context) {
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
		path      = "../static/avatars/"
		avatar = ""
		fileType  string
		err       error
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
	admin, err := srv.UpdateAdminAvatar(adminID.(int), avatar)
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
		"data":    admin,
	})
}
