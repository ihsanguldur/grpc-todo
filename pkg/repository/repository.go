package repository

import (
	"gorm.io/gorm"
	"grpc-todo/pkg/api/models"
)

type Storage interface {
	Get(todos *[]models.Todo) error
	GetOne(todo *models.Todo) error
	Create(todo models.Todo) error
	Update(todo models.Todo) error
	Delete(id uint) error
}

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{db: db}
}

func (s *storage) Get(todos *[]models.Todo) error {
	var err error
	if err = s.db.Find(&todos).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) GetOne(todo *models.Todo) error {
	var err error
	if err = s.db.First(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) Create(todo models.Todo) error {
	var err error
	if err = s.db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) Update(todo models.Todo) error {
	var err error
	if err = s.db.Model(&models.Todo{}).Updates(todo).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) Delete(id uint) error {
	var err error
	if err = s.db.Delete(&models.Todo{}, id).Error; err != nil {
		return err
	}

	return nil
}
