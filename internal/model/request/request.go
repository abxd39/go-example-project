package request

// ReminderFilter 定义查询条件
type ReminderFilter struct {
	Status       *int    `json:"status"`
	StoreIDs     []int64 `json:"storeIds"`
	CustomerName *string `json:"customerName"`
	Page         int     `json:"page"`
	PageSize     int     `json:"pageSize"`
}
