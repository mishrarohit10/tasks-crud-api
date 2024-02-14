package controllers

import (
	"example/webservices/initializers"
	"example/webservices/models"
	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	var task struct {
		ID          uint   `json:"id" gorm:"primary_key"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Due_date    string `json:"due_date" binding:"required"`
		Status      string `json:"status" binding:"required"`
	}

	c.BindJSON(&task)

	tasks := models.Task{Title: task.Title, Description: task.Description, Due_date: task.Due_date, Status: task.Status}

	result := initializers.DB.Create(&tasks)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, tasks)
}

func GetOne(c *gin.Context) {
	ID := c.Param("id")
	var task models.Task

	if err := initializers.DB.First(&task, ID).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, task)
}

func Update(c *gin.Context) {
	ID := c.Param("id")

	var task models.Task
	result := initializers.DB.First(&task, ID)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(204, gin.H{"error": err.Error()})
        return
    }

    task.Title = updatedTask.Title
    task.Description = updatedTask.Description
    task.Due_date = updatedTask.Due_date
    task.Status = updatedTask.Status

	initializers.DB.Save(&task)

	c.JSON(200, task)
}

func Delete(c *gin.Context) {
	ID := c.Param("id")
	var task models.Task

	result := initializers.DB.Delete(&task, ID)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}
	err := initializers.DB.Delete(&task, ID)

	if err.Error == nil {
		c.JSON(200, gin.H{"message": "Record deleted successfully!"})
	} else {
		c.JSON(400, gin.H{"error": "Record not found!"})
	}

}

func GetAll(c *gin.Context) {
	var tasks []models.Task
	result := initializers.DB.Find(&tasks)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, tasks)
}
