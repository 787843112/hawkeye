package dao

import (
	"errors"
	"web-server/model"
)

//CreateMovie ...
func (d *Dao) CreateMovie(movieLength, status int, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster string) (movie *model.MovieInfo, err error) {
	movie = &model.MovieInfo{
		MovieName:    movieName,
		Director:     director,
		Starring:     starring,
		Region:       region,
		MovieTypes:   movieTypes,
		MovieLength:  movieLength,
		ReleaseTime:  releaseTime,
		Status:       status,
		Introduction: introduction,
		Poster:       poster,
	}
	return movie, d.db.Create(movie).Error
}

//CreateMovieType ...
func (d *Dao) CreateMovieType(typeName string) error {
	movieType := &model.MovieType{}
	if d.db.Where("type_name = ?", typeName).First(movieType).Error == nil {
		return errors.New("类型已存在")
	}
	movieType = &model.MovieType{TypeName: typeName}
	if d.db.Create(movieType).Error != nil {
		return errors.New("添加失败")
	}
	return nil
}

//DeleteMovie ...
func (d *Dao) DeleteMovie(movie *model.MovieInfo) error {
	return d.db.Delete(movie).Error
}

//UpdateMovie ...
func (d *Dao) UpdateMovie(movieID, movieLength, status int, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster string) (movie *model.MovieInfo, err error) {
	movie = &model.MovieInfo{}
	if err = d.db.Where("movie_id = ?", movieID).First(movie).Error; err != nil {
		return nil, errors.New("电影不存在")
	}
	movie.MovieLength = movieLength
	movie.Status = status
	if movieName != "" {
		movie.MovieName = movieName
	}
	if director != "" {
		movie.Director = director
	}
	if starring != "" {
		movie.Starring = starring
	}
	if region != "" {
		movie.Region = region
	}
	if movieTypes != "" {
		movie.MovieTypes = movieTypes
	}
	if releaseTime != "" {
		movie.ReleaseTime = releaseTime
	}
	if introduction != "" {
		movie.Introduction = introduction
	}
	if poster != "" {
		movie.Poster = poster
	}
	if d.db.Save(movie).Error != nil {
		return nil, errors.New("修改失败")
	}
	return
}

//FindMovieByMovieID ...
func (d *Dao) FindMovieByMovieID(movieID int) (movie *model.MovieInfo, err error) {
	movie = &model.MovieInfo{}
	return movie, d.db.Where("movie_id = ?", movieID).First(movie).Error
}

//FindMovies ...
func (d *Dao) FindMovies(movieName, movieType string, status, page, size int) (count int, movies []*model.MovieInfo, err error) {
	if status == 0 {
		d.db.Model(model.MovieInfo{}).Where("movie_name LIKE ? AND movie_types LIKE ?", "%"+movieName+"%", "%"+movieType+"%").Count(&count)
		return count, movies, d.db.Where("movie_name LIKE ? AND movie_types LIKE ?", "%"+movieName+"%", "%"+movieType+"%").Limit(size).Offset((page - 1) * size).Find(&movies).Error
	}
	d.db.Model(model.MovieInfo{}).Where("movie_name LIKE ? AND movie_types LIKE ? AND status = ?", "%"+movieName+"%", "%"+movieType+"%", status).Count(&count)
	return count, movies, d.db.Where("movie_name LIKE ? AND movie_types LIKE ? AND status = ?", "%"+movieName+"%", "%"+movieType+"%", status).Limit(size).Offset((page - 1) * size).Find(&movies).Error
}

//FindReleasedMovies ...
func (d *Dao) FindReleasedMovies() (movies []*model.MovieInfo, err error) {
	return movies, d.db.Where("status = 2").Find(&movies).Error
}

//FindExpectMovies ...
func (d *Dao) FindExpectMovies() (movies []*model.MovieInfo, err error) {
	return movies, d.db.Where("status = 3").Limit(10).Order("want_number desc").Find(&movies).Error
}

//FindAllMovieType ...
func (d *Dao) FindAllMovieType() (movieTypes []*model.MovieType, err error) {
	return movieTypes, d.db.Find(&movieTypes).Error
}
