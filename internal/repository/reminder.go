package repository

import (
	"context"
	"example-project/internal/domain"
	"example-project/internal/model/request"

	"github.com/jinzhu/gorm"
)

type reminder struct {
	db *gorm.DB
}

// NewReminderRepository 创建一个新的提醒仓库实例
func NewReminderRepository(db *gorm.DB) ReminderRepository {
	return &reminder{
		db: db,
	}
}

// FindAll 查询满足条件的所有提醒
func (r *reminder) FindAll(ctx context.Context, filter request.ReminderFilter) ([]domain.Reminder, int, error) {
	var reminders []domain.Reminder
	query := r.db.Where(filter)
	err := query.Find(&reminders).Error
	if err != nil {
		return nil, 0, err
	}
	var count int
	query.Count(&count)
	return reminders, count, nil
}

// FindByID 根据ID查询单个提醒
func (r *reminder) FindByID(ctx context.Context, id int64) (*domain.Reminder, error) {
	var reminder domain.Reminder
	err := r.db.First(&reminder, id).Error
	if err != nil {
		return nil, err
	}
	return &reminder, nil
}

// Create 创建新的催租提醒
func (r *reminder) Create(ctx context.Context, reminder domain.Reminder) (*domain.Reminder, error) {
	err := r.db.Create(&reminder).Error
	if err != nil {
		return nil, err
	}
	return &reminder, nil
}

// Update 更新提醒信息
func (r *reminder) Update(ctx context.Context, reminder domain.Reminder) error {
	return r.db.Save(&reminder).Error
}
