package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// 2. [METHOD:GET] Menampilkan data detail list by list id.
// 4. [METHOD:GET] Menampilkan data detail sub list by sub list id.
func (h *Handler) GetTaskByIDHandler(c echo.Context) error {

	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	//preload flag
	var preloadFlag bool
	preloadSubTaskParam := c.QueryParam("preloadSubTasks")

	if preloadSubTaskParam != "" {
		preloadFlag, _ = strconv.ParseBool(preloadSubTaskParam)
	}

	task, err := h.service.TaskService().FindByID(uint(id), preloadFlag)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		return c.JSON(500, map[string]interface{}{"error": "Failed to load task"})
	}
	return c.JSON(200, map[string]interface{}{"status": "success", "data": task})
}

// 1. [METHOD:GET] Menampilkan data all list ( include pagination, filter[Search By: title, description] ) dengan atau tanpa preload sub list (dynamic)
func (h *Handler) GetAllList(c echo.Context) error {
	//pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	//Search filters
	title := c.QueryParam("title")
	description := c.QueryParam("description")

	//preload flag
	var preloadFlag bool
	preloadSubTaskParam := c.QueryParam("preloadSubTasks")

	if preloadSubTaskParam != "" {
		preloadFlag, _ = strconv.ParseBool(preloadSubTaskParam)
	}

	tasks, err := h.service.TaskService().FilterTask(title, description, page, pageSize, preloadFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to load tasks"})
	}
	return c.JSON(200, map[string]interface{}{"status": "success", "data": tasks})
}

// 3.[METHOD:GET] Menampilkan data all sub list by list id ( include pagination, filter[Search By: title, description] )
func (h *Handler) GetAllSubListByParentID(c echo.Context) error {
	strID := c.Param("parentID")
	parentID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"error": "Invalid input"})
	}

	//pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	//Search filters
	title := c.QueryParam("title")
	description := c.QueryParam("description")

	subTasks, err := h.service.TaskService().FindSubTaskByTaskID(title, description, uint(parentID), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to load subtasks"})
	}
	return c.JSON(200, map[string]interface{}{"status": "success", "data": subTasks})
}

func (h *Handler) DownloadFile(c echo.Context) error {
	fileName := c.QueryParam("fileName")
	path, err := h.service.FileService().Download(fileName)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.Inline(path, fileName)
}
