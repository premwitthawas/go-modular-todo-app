package todoHandler

import (
	"errors"
	"log"
	"net/http"
	todoModels "recap-golang-basic/modules/todo_sevices/models"
	todoUsecase "recap-golang-basic/modules/todo_sevices/usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	TodoHandler interface {
		CreatTodo(c *gin.Context)
		GetTodos(c *gin.Context)
		DeleteTodoById(c *gin.Context)
		UpdatedCompleteTodoById(c *gin.Context)
	}
	todoHandler struct {
		usecase  todoUsecase.TodoUsecase
		validate *validator.Validate
	}
)

func NewTodoHandler(usecase todoUsecase.TodoUsecase, validate *validator.Validate) TodoHandler {
	return &todoHandler{
		validate: validate,
		usecase:  usecase,
	}
}

func (h *todoHandler) CreatTodo(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(todoModels.TodoCreateReq)
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, todoModels.ErrorRespnse{
			Status:  http.StatusBadRequest,
			Message: "body request error",
		})
		return
	}
	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, todoModels.ErrorRespnse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	res, err := h.usecase.CreateTodo(ctx, req)
	if err != nil {
		log.Printf("[ERROR]: %v \n", err)
		c.JSON(http.StatusInternalServerError, todoModels.ErrorRespnse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, todoModels.Response{
		Status: http.StatusCreated,
		Data:   res,
	})
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	ctx := c.Request.Context()
	todos, err := h.usecase.GetTodos(ctx)
	if err != nil {
		log.Printf("[ERROR]: %v \n", err)
		c.JSON(http.StatusInternalServerError, todoModels.ErrorRespnse{
			Status:  http.StatusInternalServerError,
			Message: "InternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, todoModels.Response{
		Status: http.StatusOK,
		Data:   todos,
	})
}

func (h *todoHandler) DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	err := h.usecase.DeleteTodoById(ctx, id)
	if err != nil {
		log.Printf("[ERROR]: %v \n", err)

		switch {
		case errors.Is(err, todoModels.ErrParseUUIDError):
			c.JSON(http.StatusBadRequest, todoModels.ErrorRespnse{
				Status:  http.StatusBadRequest,
				Message: "uuid invalid",
			})
		case errors.Is(err, todoModels.ErrNotFoundTodo):
			c.JSON(http.StatusNotFound, todoModels.ErrorRespnse{
				Status:  http.StatusNotFound,
				Message: "not found",
			})
		default:
			c.JSON(http.StatusInternalServerError, todoModels.ErrorRespnse{
				Status:  http.StatusInternalServerError,
				Message: "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, todoModels.ResponseMessage{
		Status:  http.StatusOK,
		Message: "deleted success.",
	})
}

func (h *todoHandler) UpdatedCompleteTodoById(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	err := h.usecase.UpdatedTodoCompletedById(ctx, id)
	if err != nil {
		log.Printf("[ERROR]: %v \n", err)

		switch {
		case errors.Is(err, todoModels.ErrParseUUIDError):
			c.JSON(http.StatusBadRequest, todoModels.ErrorRespnse{
				Status:  http.StatusBadRequest,
				Message: "uuid invalid",
			})
		case errors.Is(err, todoModels.ErrNotFoundTodo):
			c.JSON(http.StatusNotFound, todoModels.ErrorRespnse{
				Status:  http.StatusNotFound,
				Message: "not found",
			})
		default:
			c.JSON(http.StatusInternalServerError, todoModels.ErrorRespnse{
				Status:  http.StatusInternalServerError,
				Message: "internal server error",
			})
		}
		return
	}
	c.JSON(http.StatusOK, todoModels.ResponseMessage{
		Status:  http.StatusOK,
		Message: "updated success.",
	})
}
