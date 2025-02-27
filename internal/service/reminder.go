package service

import (
	"context"
	"example-project/internal/domain"
	"example-project/internal/model/request"
)

// ReminderService 定义催租服务接口
type ReminderService interface {
	// GetReminders 获取符合条件的提醒列表
	GetReminders(ctx context.Context, filter request.ReminderFilter) ([]domain.Reminder, int, error)

	// GetReminderByID 根据ID获取单个提醒
	GetReminderByID(ctx context.Context, id int64) (*domain.Reminder, error)

	// CreateReminder 创建新的催租提醒
	CreateReminder(ctx context.Context, reminder domain.Reminder) (*domain.Reminder, error)

	// UpdateReminder 更新提醒信息
	UpdateReminder(ctx context.Context, reminder domain.Reminder) error
}
