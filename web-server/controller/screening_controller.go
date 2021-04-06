package controller

import (
	"sync"
	"time"
	"web-server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//添加场次
func addScreening(c *gin.Context) {
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
		screening *model.ScreeningInfo
		v         = new(struct {
			CinemaID  int    `json:"CinemaID" form:"CinemaID"`
			HallID    int    `json:"HallID" form:"HallID"`
			MovieID   int    `json:"MovieID" form:"MovieID"`
			Price     int    `json:"Price" form:"Price"`
			StartTime string `json:"StartTime" form:"StartTime"`
			EndTime   string `json:"EndTime" form:"EndTime"`
		})
		timeLayout = "2006-01-02 15:04:05"
		startTime  time.Time
		endTime    time.Time
		err        error
	)
	if err = c.ShouldBind(&v); err != nil || v.CinemaID == 0 || v.HallID == 0 || v.MovieID == 0 || v.Price == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.CinemaID < 0 || v.HallID < 0 || v.MovieID < 0 || v.Price < 0 || v.StartTime == "" || v.EndTime == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	//权限校验
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
	startTime, err = time.ParseInLocation(timeLayout, v.StartTime, time.Local) //string转time
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "开始时间值不合法",
		})
		return
	}
	endTime, err = time.ParseInLocation(timeLayout, v.EndTime, time.Local) //string转time
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "结束时间值不合法",
		})
		return
	}
	if !startTime.After(time.Now()) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "开始时间不在现在时间之后",
		})
		return
	}
	if !startTime.Before(endTime) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "开始时间不在结束时间之前",
		})
		return
	}
	screening, err = srv.AddScreening(v.HallID, v.MovieID, v.Price, startTime, endTime)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if err = srv.InitializeSeatStatus(screening.HallID, screening.ScreeningID); err != nil {
		if err = srv.DeleteScreeningInfo(screening.ScreeningID); err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "删除场次失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    1,
			"message": "初始化座位状态失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    screening,
	})
}

//删除某些场次
func deleteScreeings(c *gin.Context) {
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
			CinemaID     int   `json:"CinemaID" form:"CinemaID"`
			ScreeningIDs []int `json:"ScreeningIDs" form:"ScreeningIDs"`
		})
		failID = []int{} //删除失败的影院ID数组
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.ScreeningIDs == nil || v.CinemaID < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.ScreeningIDs) == 0 {
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
	for _, v := range v.ScreeningIDs {
		wg.Add(1)
		go func(v int) {
			if err = srv.DeleteScreeningInfo(v); err != nil {
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

//管理员查询场次
func adminGetScreeings(c *gin.Context) {
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
		screenings []*model.ScreeningInfo
		v          = new(struct {
			CinemaID  int    `form:"CinemaID"`
			HallID    int    `form:"HallID"`
			StartDate string `form:"StartDate"`
			EndDate   string `form:"EndDate"`
			Page      int    `form:"Page"`
			Size      int    `form:"Size"`
		})
		timeLayout = "2006-01-02"
		//today, _   = time.ParseInLocation(timeLayout, time.Now().Format(timeLayout), time.Local)
		startDate time.Time
		endDate   time.Time
		err       error
	)
	if err = c.ShouldBind(&v); err != nil || v.CinemaID == 0 || v.HallID == 0 || v.Page == 0 || v.Size == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.CinemaID < 0 || v.HallID < 0 || v.Page < 0 || v.Size < 0 {
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
	if v.StartDate != "" && v.EndDate != "" {
		startDate, err = time.ParseInLocation(timeLayout, v.StartDate, time.Local) //string转time
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "开始日期值不合法",
			})
			return
		}
		endDate, err = time.ParseInLocation(timeLayout, v.EndDate, time.Local) //string转time
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "结束日期值不合法",
			})
			return
		}
		// if startDate.Before(today) {
		// 	c.JSON(200, gin.H{
		// 		"code":    1,
		// 		"message": "开始日期不能在今天之前",
		// 	})
		// 	return
		// }
		if startDate.After(endDate) {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "开始日期不在结束日期之前",
			})
			return
		}
		count, screenings, err := srv.AdminQueryScreenings(v.HallID, v.Page, v.Size, startDate, endDate)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "获取失败",
			})
			return
		}
		if len(screenings) == 0 {
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
			"data":    screenings,
		})
		return
	}
	count, screenings, err := srv.AdminQueryAllScreenings(v.HallID, v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(screenings) == 0 {
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
		"data":    screenings,
	})
}

//用户查询场次
func userGetScreeings(c *gin.Context) {
	var (
		screenings []*model.ScreeningInfo
		v          = new(struct {
			CinemaID   int    `form:"CinemaID"`
			MovieID    int    `form:"MovieID"`
			SelectDate string `form:"SelectDate"`
			Page       int    `form:"Page"`
			Size       int    `form:"Size"`
		})
		timeLayout = "2006-01-02"
		today, _   = time.ParseInLocation(timeLayout, time.Now().Format(timeLayout), time.Local)
		selectDate time.Time
		err        error
	)
	if err = c.ShouldBind(&v); err != nil || v.Page == 0 || v.Size == 0 || v.CinemaID == 0 || v.MovieID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if v.CinemaID < 0 || v.MovieID < 0 || v.Page < 0 || v.Size < 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	selectDate, err = time.ParseInLocation(timeLayout, v.SelectDate, time.Local) //string转time
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "选择的日期值不合法",
		})
		return
	}
	if selectDate.Before(today) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "选择的日期不能在今天之前",
		})
		return
	}
	count, screenings, err := srv.UserQueryScreenings(v.CinemaID, v.MovieID, v.Page, v.Size, selectDate)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(screenings) == 0 {
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
		"data":    screenings,
	})
}
