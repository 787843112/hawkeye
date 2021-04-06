package service

import (
	"web-server/config"
	"web-server/dao"
)

// Service struct
type Service struct {
	c   *config.Config
	dao *dao.Dao
}

//New init
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}

//Close service
func (s *Service) Close() {
	s.dao.Close()
}
