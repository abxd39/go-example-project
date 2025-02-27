package model

// RentReminder 表示数据库中的催租记录实体
type RentReminder struct {
	Id               int64  `gorm:"primaryKey;column:id" json:"Id,string"`
	MainOrderId      int64  `gorm:"column:main_order_id" json:"mainOrderId,string"`
	MainOrderNo      string `gorm:"column:main_order_no" json:"mainOrderNo"`
	Status           int    `gorm:"column:status" json:"status"`
	LastEndDate      string `gorm:"column:last_end_date" json:"lastEndDate"`
	RentReminderDate string `gorm:"column:rent_reminder_date" json:"rentReminderDate"`
	OverdueDays      int    `gorm:"column:overdue_days" json:"overdueDays"`
	CustomerName     string `gorm:"column:customer_name" json:"customerName"`
	CustomerPhone    string `gorm:"column:mobile" json:"customerPhone"`
	// 其他字段...
}

// WorkerMinderOrderListPage 表示催租列表查询参数
type WorkerMinderOrderListPage struct {
	PageNum       int    `json:"pageNum"`
	PageSize      int    `json:"pageSize"`
	PreferStatus  int    `json:"preferStatus"` // 审批状态
	DayNumber     int    `json:"dayNumber"`    // 天数筛选
	StoreId       int64  `json:"storeId,string"`
	CustomerPhone string `json:"customerPhone"`
	CustomerName  string `json:"customerName"`
	// 其他筛选字段...
}
