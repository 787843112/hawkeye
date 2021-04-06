package controller

import (
	"web-server/middleware"
	"web-server/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const key = "MY KEY"

var srv *service.Service

//Init ...
func Init(s *service.Service) {
	srv = s
	initRouter()
}

func initRouter() {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(key))
	if err != nil {
		panic("Redis数据库连接失败!")
	}
	router := gin.Default()
	router.Static("/static", "../static")
	router.Use(sessions.Sessions("session_id", store))
	user := router.Group("/user")
	{
		user.POST("/login", userLogin)
		user.GET("/autologin", userAutoLogin)
		user.GET("/logout", userLogout)
		user.POST("/register", userRegister)
		user.POST("/update", userUpdate)
		user.POST("/upload", updateUserAvatar)
	}
	want := router.Group("/want")
	{
		want.GET("/get", getWant)
		want.GET("/getmovieid", getWantMovieID)
		want.POST("/add", addWant)
		want.POST("/delete", deleteWant)
	}
	order := router.Group("/order")
	{
		order.GET("/get", getOrder)
		order.POST("/add", addOrder)
		order.POST("/delete", deleteOrder)
	}
	seen := router.Group("/seen")
	{
		seen.GET("/get", getSeen)
		seen.GET("/getbymovieid", getSeenByMovieID)
		seen.POST("/delete", deleteSeen)
	}
	evaluation := router.Group("/evaluation")
	{
		evaluation.POST("/update", updateEvaluation)
		evaluation.POST("/delete", deleteEvaluation)
	}
	admin := router.Group("/admin")
	{
		admin.POST("/login", adminLogin)
		admin.GET("/autologin", adminAutoLogin)
		admin.GET("/logout", adminLogout)
		admin.POST("/register", adminRegister)
		admin.POST("/upload", updateAdminAvatar)
		admin.POST("/update", adminUpdate)
		admin.POST("/permissions", middleware.VerifyAdmin(), permissions)
		admin.GET("/get", middleware.VerifyAdmin(), getAdmin)
		admin.POST("/delete", middleware.VerifyAdmin(), deleteAdmin)
	}
	manage := router.Group("manage", middleware.VerifyAdmin())
	{
		manage.GET("/get", getManage)
		manage.POST("/add", addManage)
		manage.POST("/delete", deleteManage)
	}
	movie := router.Group("/movie")
	{
		movie.POST("/upload", middleware.VerifyAdmin(), uploadPoster)
		movie.POST("/add", middleware.VerifyAdmin(), addMovie)
		movie.POST("/addtype", middleware.VerifyAdmin(), addMovieType)
		movie.POST("/delete", middleware.VerifyAdmin(), deleteMovie)
		movie.POST("/update", middleware.VerifyAdmin(), updateMovie)
		movie.GET("/get", getMovies)
		movie.GET("/getbyid", getMovieByID)
		movie.GET("/getreleased", getReleasedMovies)
		movie.GET("/getexpect", getExpectMovies)
		movie.GET("/gettype", getMovieType)
	}
	cinema := router.Group("/cinema")
	{
		cinema.POST("/add", middleware.VerifyAdmin(), addCinema)
		cinema.POST("/delete", middleware.VerifyAdmin(), deleteCinema)
		cinema.POST("/update", updateCinema)
		cinema.GET("/get", getCinema)
		cinema.GET("/getbycity", middleware.VerifyAdmin(), getCinemaByCity)
		cinema.GET("/getmanage", adminManage)
	}
	hall := router.Group("/hall")
	{
		hall.POST("/add", addHall)
		hall.POST("/delete", deleteHall)
		hall.POST("/update", updateHall)
		hall.GET("/get", getHall)
		hall.GET("/getbyid", getHallByID)
	}
	seat := router.Group("/seat")
	{
		seat.POST("/add", addSeats)
		seat.POST("/delete", deleteSeats)
		seat.GET("/get", getSeat)
	}
	status := router.Group("/status")
	{
		status.POST("/update", updateStatus)
		status.GET("/get", getStatus)
	}
	screening := router.Group("/screening")
	{
		screening.POST("/add", addScreening)
		screening.POST("/delete", deleteScreeings)
		screening.GET("/adminget", adminGetScreeings)
		screening.GET("/userget", userGetScreeings)
	}
	router.Run(":8080")
}
