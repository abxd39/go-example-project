package repository

import (
	"context"
	"example-project/internal/domain"
	"example-project/internal/model/request"
)

// ReminderRepository 定义数据访问接口
type ReminderRepository interface {
	// FindAll 查询满足条件的所有提醒
	FindAll(ctx context.Context, filter request.ReminderFilter) ([]domain.Reminder, int, error)

	// FindByID 根据ID查询单个提醒
	FindByID(ctx context.Context, id int64) (*domain.Reminder, error)

	// Create 创建新的催租提醒
	Create(ctx context.Context, reminder domain.Reminder) (*domain.Reminder, error)

	// Update 更新提醒信息
	Update(ctx context.Context, reminder domain.Reminder) error
}
