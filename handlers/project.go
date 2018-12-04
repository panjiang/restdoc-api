package handlers

import (
	"restdoc-api/app"
	"restdoc-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type newProjectReq struct {
	Name string `json:"name" binding:"required,max=20"`
}

// NewProject 创建新项目
func NewProject(c *gin.Context) {
	var err error
	var json newProjectReq

	if err := c.ShouldBindJSON(&json); err != nil {
		BadRequestBind(c, err)
		return
	}

	err = app.DB.Where("name = ?", json.Name).First(&models.Project{}).Error
	if !RecordNotFound(err) {
		BadRequest(c, "name", BadOccupied)
		return
	}

	project := models.Project{
		Name: json.Name,
	}
	err = app.DB.Create(&project).Error
	CheckError(err)

	c.JSON(200, gin.H{
		"project": project,
	})
}

type projectsItem struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetProjects 获取所有项目
func GetProjects(c *gin.Context) {
	var err error
	var items []projectsItem
	err = app.DB.Table("projects").Order("id desc").Find(&items).Error
	CheckError(err)

	Success(c, gin.H{
		"projects": items,
	})
}

// GetProject 获取项目数据
func GetProject(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var project models.Project
	err := app.DB.First(&project, id).Error
	if RecordNotFound(err) {
		BadRequest(c, "id", BadNotFound)
		return
	}

	Success(c, gin.H{
		"project": project,
	})
}
