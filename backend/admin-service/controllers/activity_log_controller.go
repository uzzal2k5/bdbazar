package controllers

import (
	"admin-service/models"
	"admin-service/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityLogController struct {
	Repo repository.ActivityLogRepository
}

func NewActivityLogController(repo repository.ActivityLogRepository) *ActivityLogController {
	return &ActivityLogController{Repo: repo}
}

func (c *ActivityLogController) CreateActivityLog(ctx *gin.Context) {
	var log models.ActivityLog
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Repo.Create(&log); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity log"})
		return
	}

	ctx.JSON(http.StatusCreated, log)
}

func (c *ActivityLogController) GetActivityLogs(ctx *gin.Context) {
	logs, err := c.Repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activity logs"})
		return
	}
	ctx.JSON(http.StatusOK, logs)
}

func (c *ActivityLogController) GetActivityLogByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	log, err := c.Repo.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Activity log not found"})
		return
	}

	ctx.JSON(http.StatusOK, log)
}
