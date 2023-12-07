package repository

import (
	"github.com/marioTiara/todolistapp/internal/api/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindAll() ([]models.Task, error)
	FindByID(ID uint, preload bool) (models.Task, error)
	Create(task models.Task) (models.Task, error)
	CreateSubTask(task models.Task) (models.Task, error)
	FilterByTitleAndDescription(title, description string, page, limit int, preload bool) ([]models.Task, error)
}

type task_repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &task_repository{db: db}
}

func (r *task_repository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *task_repository) FindByID(ID uint, preload bool) (models.Task, error) {
	var parentTaskWithChildren models.Task
	var err error
	if preload {
		err = r.db.Preload("Children").First(&parentTaskWithChildren, ID).Error
	} else {
		err = r.db.First(&parentTaskWithChildren, ID).Error
	}

	return parentTaskWithChildren, err
}

func (r *task_repository) Create(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *task_repository) FilterByTitleAndDescription(title, description string, page, limit int, preload bool) ([]models.Task, error) {
	var tasks []models.Task
	query := r.db.Model(&models.Task{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	offset := (page - 1) * limit

	if preload {
		query = query.Preload("Children")
	}
	if !preload {
		query = query.Where("parent_id IS NULL")
	}
	if err := query.Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *task_repository) CreateSubTask(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}
