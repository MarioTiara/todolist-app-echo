package dtos

type AddTaskRequest struct {
	Title       string           `json:"title" binding:"required"`
	Description string           `json:"description"`
	Priority    int              `json:"priority"`
	Childrens   []AddTaskRequest `json:"childrens"`
}
