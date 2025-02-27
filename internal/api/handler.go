package api

import (
	"example-project/internal/domain"
	"example-project/internal/model/request"
	"example-project/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReminderHandler struct {
	service service.ReminderService
}

func NewReminderHandler(service service.ReminderService) *ReminderHandler {
	return &ReminderHandler{service: service}
}

// ListReminders 获取催租列表
func (h *ReminderHandler) ListReminders(c *gin.Context) {
	var filter request.ReminderFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 从上下文中获取用户信息
	// userID := getUserIDFromContext(c)

	reminders, total, err := h.service.GetReminders(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  reminders,
		"total": total,
	})
}

// GetReminder 获取单个催租信息
func (h *ReminderHandler) GetReminder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	reminder, err := h.service.GetReminderByID(c.Request.Context(), id)
	if err != nil {
		if _, ok := err.(domain.ErrNotFound); ok {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
		}
		return
	}

	c.JSON(http.StatusOK, reminder)
}
