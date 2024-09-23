package handler

import (
	"net/http"
	"strconv"

	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary      createTask
// @Security ApiKeyAuth
// @Description  create task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        input body domain.TaskInput true "tasks info"
// @Success      201  {integer} integer 1
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	userId, err := h.GetUserID(c)
	if err != nil {
		return
	}

	var input domain.TaskInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	idTask, err := h.service.Tasks.Create(input, userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": idTask,
	})
}

type allTasksResponse struct {
	Data []domain.Task `json:"data"`
}

// @Summary      allTasks
// @Security ApiKeyAuth
// @Description  get all tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200  {object} allTasksResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/tasks [get]
func (h *Handler) GetAllTasks(c *gin.Context) {
	userId, err := h.GetUserID(c)
	if err != nil {
		return
	}
	tasks, err := h.service.Tasks.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTasksResponse{Data: tasks})
}

// @Summary      getTaskById
// @Security ApiKeyAuth
// @Description  get task by id
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  domain.Task
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/tasks/{id} [get]
func (h *Handler) GetTaskById(c *gin.Context) {
	userId, err := h.GetUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id in param")
		return
	}
	task, err := h.service.GetById(id, userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

// @Summary      deleteTask
// @Security ApiKeyAuth
// @Description  delete task by id
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      204  {object}  StatusResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	userId, err := h.GetUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id in param")
		return
	}
	if err := h.service.Delete(id, userId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, StatusResponse{Status: "Ok"})
}

// @Summary      updateTask
// @Security ApiKeyAuth
// @Description  update task by id
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Param        input body domain.TaskUpdate true "tasks update info"
// @Success      204  {object}  StatusResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/tasks/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	userId, err := h.GetUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id in param")
		return
	}
	var input domain.TaskUpdate
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.service.Update(id, input, userId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{Status: "Ok"})
}
