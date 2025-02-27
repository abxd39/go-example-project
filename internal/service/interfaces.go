package service

import (
	"context"
	"example-project/internal/domain"
	"example-project/internal/model/request"
	"example-project/internal/repository"
	"time"
)

// reminderService 实现 ReminderService 接口
type reminderService struct {
	repo repository.ReminderRepository
}

// NewReminderService 创建催租服务实例
func NewReminderService(repo repository.ReminderRepository) ReminderService {
	return &reminderService{repo: repo}
}

// GetReminders 实现接口方法
func (s *reminderService) GetReminders(ctx context.Context, filter request.ReminderFilter) ([]domain.Reminder, int, error) {
	return s.repo.FindAll(ctx, filter)
}

// GetReminderByID 实现接口方法
func (s *reminderService) GetReminderByID(ctx context.Context, id int64) (*domain.Reminder, error) {
	return s.repo.FindByID(ctx, id)
}

// CreateReminder 实现接口方法
func (s *reminderService) CreateReminder(ctx context.Context, reminder domain.Reminder) (*domain.Reminder, error) {
	// 计算逾期天数
	if !reminder.LastEndDate.IsZero() {
		endDateStr := reminder.LastEndDate.Format("2006-01-02")
		reminder.OverdueDays = s.CalculateOverdueDays(endDateStr)
	}

	return s.repo.Create(ctx, reminder)
}

// UpdateReminder 实现接口方法
func (s *reminderService) UpdateReminder(ctx context.Context, reminder domain.Reminder) error {
	return s.repo.Update(ctx, reminder)
}

// CalculateOverdueDays 计算逾期天数
func (s *reminderService) CalculateOverdueDays(endDate string) int {
	if len(endDate) != 10 {
		return 0
	}

	endDate = endDate + " 23:59:59"
	parsedDate, err := time.ParseInLocation("2006-01-02 15:04:05", endDate, time.Local)
	if err != nil {
		return 0
	}

	diff := time.Now().Unix() - parsedDate.Unix()
	return int(diff/86400) + 1
}
