package services

import "../models"
import "../dao"

type GiftServer interface {
	GetAll() []models.Gift
	CountAll() int64
	Get(id int) *models.Gift
	Delete(id int) error
	Update(data *models.Gift, columns []string) error
	Create(data *models.Gift) error
}

type giftServer struct {
	dao *dao.GiftDao
}

func (s *giftServer) GetAll() []models.Gift {
	return s.dao.GetAll()
}

func (s *giftServer) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *giftServer) Get(id int) *models.Gift {
	return s.dao.Get(id)
}

func (s *giftServer) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *giftServer) Update(data *models.Gift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftServer) Create(data *models.Gift) error {
	return s.dao.Create(data)
}

func newGiftServer(dao *dao.GiftDao) *giftServer {
	return &giftServer{dao: dao}
}
