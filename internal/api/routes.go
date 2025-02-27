package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(route gin.IRouter, handle *ReminderHandler) {

	group := route.Group("/api/v1")
	{
		group.GET("/reminders", handle.ListReminders)
	}
}
