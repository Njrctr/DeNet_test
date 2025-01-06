package handlers

import (
	"net/http"
	"strconv"

	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create task
// @Security ApiKeyAuth
// @Tags Tasks
// @Description create task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param input body models.TaskCreate true "task data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /task/create [post]
func (h *Handler) taskCreate(c *gin.Context) {
	var task models.TaskCreate

	if err := c.BindJSON(&task); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	err := task.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newTaskId, err := h.services.CreateTask(task)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"new_task_id": newTaskId,
	})
}

type taskComplete struct {
	TaskId int `json:"task_id"`
}

// @Summary Complete task
// @Security ApiKeyAuth
// @Tags Tasks
// @Description complete task
// @ID complete-task
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param task_id body taskComplete true "task id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/task/complete [post]
func (h *Handler) taskComplete(c *gin.Context) {
	var taskComplete taskComplete

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param")
		return
	}

	if err := c.BindJSON(&taskComplete); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return

	}

	err = h.services.CompleteTask(userId, taskComplete.TaskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"succes": "ok",
	})
}

type getAllTasksResponce struct {
	Data []models.Task `json:"tasks"`
}

// @Summary Get All Tasks
// @Tags Tasks
// @Description get all tasks
// @ID get-all-tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllTasksResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /task/all [get]
func (h *Handler) taskGetAll(c *gin.Context) {

	tasks, err := h.services.GetAllTasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponce{
		Data: tasks,
	})
}
