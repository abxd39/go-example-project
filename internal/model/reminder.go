package model

import "time"

// Reminder 表示催租提醒实体
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
