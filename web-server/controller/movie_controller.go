package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"sync"
	"web-server/model"

	"github.com/gin-gonic/gin"
)

var poster string //电影海报地址

//添加电影类型
func addMovieType(c *gin.Context) {
	//权限校验
	var (
		movieType = &model.MovieType{}
		err       error
	)
	if err = c.ShouldBind(&movieType); err != nil || movieType.TypeName == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(movieType.TypeName) > 6 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "类型值太长",
		})
		return
	}
	if err = srv.AddMovieType(movieType.TypeName); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
	})
}

//添加电影
func addMovie(c *gin.Context) {
	//权限校验
	if poster == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请先上传电影海报",
		})
		return
	}
	var (
		movie = &model.MovieInfo{}
		err   error
	)
	if err = c.ShouldBind(&movie); err != nil || movie.Status == 0 || movie.MovieLength == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if movie.Status < 0 || movie.Status != 1 && movie.Status != 2 &&
		movie.Status != 3 || movie.MovieLength < 0 ||
		movie.MovieName == "" || movie.Director == "" ||
		movie.Starring == "" || movie.Region == "" ||
		movie.MovieTypes == "" ||
		movie.Introduction == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "电影信息不完整",
		})
		return
	}
	movie, err = srv.AddMovie(movie.MovieLength, movie.Status,
		movie.MovieName, movie.Director, movie.Starring, movie.Region,
		movie.MovieTypes, movie.ReleaseTime, movie.Introduction, poster)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    movie,
	})
	poster = ""
	defer func() {
		poster = ""
	}()
}

//删除某些电影
func deleteMovie(c *gin.Context) {
	var (
		v = new(struct {
			MovieIDs []int `json:"MovieIDs" form:"MovieIDs"`
		})
		failID = []int{} //删除失败的电影ID数组
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)
	if err = c.ShouldBind(&v); err != nil || v.MovieIDs == nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if len(v.MovieIDs) == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "提交信息为空",
		})
		return
	}
	for _, v := range v.MovieIDs {
		wg.Add(1)
		go func(v int) {
			if err = srv.DeleteMovieInfo(v); err != nil {
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
		"message": "删除电影成功",
		"data":    failID,
	})
}

//上传电影海报
func uploadPoster(c *gin.Context) {
	poster = ""
	var (
		path     = "../static/posters/"
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
			"message": "无法获取图片文件类型",
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
	poster = fileName + fileType
	c.JSON(200, gin.H{
		"code":    0,
		"message": "上传文件成功",
		"data":    poster,
	})
}

//修改电影信息
func updateMovie(c *gin.Context) {
	//权限校验
	var (
		movie = &model.MovieInfo{}
		err   error
	)
	if err = c.ShouldBind(&movie); err != nil || movie.Status != 1 && movie.Status != 2 && movie.Status != 3 || movie.MovieLength < 1 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "请求参数错误",
		})
		return
	}
	if movie.MovieName == "" || movie.Director == "" ||
		movie.Starring == "" || movie.Region == "" ||
		movie.Introduction == "" {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "电影信息不完善",
		})
		return
	}
	movie, err = srv.UpdateMovieInfo(movie.MovieID, movie.MovieLength, movie.Status,
		movie.MovieName, movie.Director, movie.Starring, movie.Region, movie.MovieTypes,
		movie.ReleaseTime, movie.Introduction, poster)
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
		"data":    movie,
	})
	poster = ""
}

//根据条件获取电影列表
func getMovies(c *gin.Context) {
	var (
		v = new(struct {
			MovieName  string `form:"MovieName"`
			MovieTypes string `form:"MovieTypes"`
			Status     int    `form:"Status"`
			Page       int    `form:"Page"`
			Size       int    `form:"Size"`
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
	if v.Page < 0 || v.Size < 0 || v.Status < 0 || v.Status != 0 && v.Status != 1 && v.Status != 2 && v.Status != 3 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "数值不合法",
		})
		return
	}
	count, movies, err := srv.QueryMovies(v.MovieName, v.MovieTypes, v.Status, v.Page, v.Size)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	if len(movies) == 0 {
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
		"data":    movies,
	})
}

//根据电影ID获取电影信息
func getMovieByID(c *gin.Context) {
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
	movie, err := srv.QueryMovieInfoByMovieID(v.MovieID)
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
		"data":    movie,
	})
}

//获取正在上映的电影列表
func getReleasedMovies(c *gin.Context) {
	movies, err := srv.QueryReleasedMovies()
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
		"data":    movies,
	})
}

//获取最受期待的10部电影
func getExpectMovies(c *gin.Context) {
	movies, err := srv.QueryExpectMovies()
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
		"data":    movies,
	})
}

//获取所有的电影类型
func getMovieType(c *gin.Context) {
	//权限校验
	movieTypes, err := srv.QueryAllMovieType()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"messqge": "获取成功",
		"data":    movieTypes,
	})
}
