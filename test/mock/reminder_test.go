package test

import (
	"context"
	"example-project/internal/domain"
	"example-project/internal/model/request"
	"example-project/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository 是 repository.ReminderRepository 的模拟实现
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAll(ctx context.Context, filter request.ReminderFilter) ([]domain.Reminder, int, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]domain.Reminder), args.Int(1), args.Error(2)
}

func (m *MockRepository) FindByID(ctx context.Context, id int64) (*domain.Reminder, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Reminder), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, reminder domain.Reminder) (*domain.Reminder, error) {
	args := m.Called(ctx, reminder)
	return args.Get(0).(*domain.Reminder), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, reminder domain.Reminder) error {
	args := m.Called(ctx, reminder)
	return args.Error(0)
}

func TestGetReminders(t *testing.T) {
	// 创建模拟仓库
	mockRepo := new(MockRepository)

	// 预设模拟行为
	expectedReminders := []domain.Reminder{
		{
			ID:          1,
			MainOrderID: 1001,
			Status:      1,
		},
	}

	filter := request.ReminderFilter{
		Page:     1,
		PageSize: 10,
	}

	mockRepo.On("FindAll", mock.Anything, filter).Return(expectedReminders, 1, nil)

	// 创建服务
	reminderService := service.NewReminderService(mockRepo)

	// 调用被测试方法
	reminders, total, err := reminderService.GetReminders(context.Background(), filter)

	// 断言结果
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Equal(t, expectedReminders, reminders)

	// 验证模拟对象的调用
	mockRepo.AssertExpectations(t)
}

func TestCreateReminder(t *testing.T) {
	mockRepo := new(MockRepository)
	reminderService := service.NewReminderService(mockRepo)

	// 预设模拟行为
	newReminder := domain.Reminder{
		ID:          2,
		MainOrderID: 1002,
		Status:      1,
	}

	mockRepo.On("Create", mock.Anything, newReminder).Return(&newReminder, nil)

	// 调用被测试方法
	createdReminder, err := reminderService.CreateReminder(context.Background(), newReminder)

	// 断言结果
	assert.NoError(t, err)
	assert.Equal(t, &newReminder, createdReminder)

	// 验证模拟对象的调用
	mockRepo.AssertExpectations(t)
}
