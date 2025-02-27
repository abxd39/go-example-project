package domain

import (
	"fmt"
	"time"
)

type Reminder struct {
	ID               int64     `json:"id"`
	MainOrderID      int64     `json:"mainOrderId"`
	MainOrderNo      string    `json:"mainOrderNo"`
	Status           int       `json:"status"`
	LastEndDate      time.Time `json:"lastEndDate"`
	RentReminderDate time.Time `json:"rentReminderDate"`
	OverdueDays      int       `json:"overdueDays"`
	CustomerName     string    `json:"customerName"`
	CustomerPhone    string    `json:"customerPhone"`
}

// 定义领域内的错误类型
type ErrNotFound struct {
	ID int64
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("reminder with ID %d not found", e.ID)
}
