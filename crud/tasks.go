package crud

import (
	"fmt"

	"github.com/models"
)

func CreateTask(task *models.Task) error {
	err := models.DB.Model(&models.Task{}).Create(task).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(id int32, task *models.Task) error {
	err := models.DB.Model(&models.Task{}).Where("id = ?", id).Updates(&models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      models.TaskStatus(task.Status),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int32) error {
	err := models.DB.Model(&models.Task{}).Delete(&models.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func ListTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	err := models.DB.Model(&models.Task{}).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SearchTasks(search string)([]models.Task, error){
	tasks := make([]models.Task, 0)
	searchQuery := fmt.Sprintf("SELECT * FROM tasks WHERE (title LIKE '%%%s%%' OR status LIKE '%%%s%%' OR description LIKE '%%%s%%') AND (deleted_at is NULL)", search, search, search)
	fmt.Println(searchQuery)
	err := models.DB.Raw(searchQuery).Find(&tasks).Error
	if err!=nil{
		return nil, err
	}
	return tasks, nil
}
