package service

import "web-server/model"

//AddMovie ...
func (s *Service) AddMovie(movieLength, status int, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster string) (movie *model.MovieInfo, err error) {
	return s.dao.CreateMovie(movieLength, status, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster)
}

//AddMovieType ...
func (s *Service) AddMovieType(typeName string) error {
	return s.dao.CreateMovieType(typeName)
}

//DeleteMovieInfo ...
func (s *Service) DeleteMovieInfo(movieID int) (err error) {
	movie, err := s.dao.FindMovieByMovieID(movieID)
	if err != nil {
		return
	}
	return s.dao.DeleteMovie(movie)
}

//UpdateMovieInfo ...
func (s *Service) UpdateMovieInfo(movieID, movieLength, status int, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster string) (movie *model.MovieInfo, err error) {
	return s.dao.UpdateMovie(movieID, movieLength, status, movieName, director, starring, region, movieTypes, releaseTime, introduction, poster)
}

//QueryMovieInfoByMovieID ...
func (s *Service) QueryMovieInfoByMovieID(movieID int) (movie *model.MovieInfo, err error) {
	return s.dao.FindMovieByMovieID(movieID)
}

//QueryMovies ...
func (s *Service) QueryMovies(movieName, movieType string, status, page, size int) (count int, movies []*model.MovieInfo, err error) {
	return s.dao.FindMovies(movieName, movieType, status, page, size)
}

//QueryReleasedMovies ...
func (s *Service) QueryReleasedMovies() (movies []*model.MovieInfo, err error) {
	return s.dao.FindReleasedMovies()
}

//QueryExpectMovies ...
func (s *Service) QueryExpectMovies() (movies []*model.MovieInfo, err error) {
	return s.dao.FindExpectMovies()
}

//QueryAllMovieType ...
func (s *Service) QueryAllMovieType() (movieTypes []*model.MovieType, err error) {
	return s.dao.FindAllMovieType()
}
