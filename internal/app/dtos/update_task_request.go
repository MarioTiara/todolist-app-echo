package dtos

type UpdateTaskRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Checked     bool   `json:"checked"`
	ParentID    *uint  `json:"parent_id"`
}
