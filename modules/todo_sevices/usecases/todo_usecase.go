package todoUsecase

import (
	"context"
	todoModels "recap-golang-basic/modules/todo_sevices/models"
	todoRepository "recap-golang-basic/modules/todo_sevices/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func todoMapTodoRes(todo todoRepository.Todo) todoModels.TodoRes {
	return todoModels.TodoRes{
		ID:        todo.ID.String(),
		Title:     todo.Title,
		Email:     todo.Email,
		Done:      todo.Done.Bool,
		CreatedAt: todo.CreatedAt.Time,
		UpdatedAt: todo.UpdatedAt.Time,
	}
}

type (
	TodoUsecase interface {
		CreateTodo(ctx context.Context, req *todoModels.TodoCreateReq) (todoModels.TodoRes, error)
		GetTodos(ctx context.Context) ([]todoModels.TodoRes, error)
		DeleteTodoById(ctx context.Context, id string) error
		UpdatedTodoCompletedById(ctx context.Context, id string) error
	}
	todoUsecase struct {
		repo todoRepository.Querier
	}
)

func NewTodousecase(repo todoRepository.Querier) TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}

func (u *todoUsecase) CreateTodo(ctx context.Context, req *todoModels.TodoCreateReq) (todoModels.TodoRes, error) {
	payload := todoRepository.InsertTodoParams{
		Title: req.Title,
		Email: req.Email,
	}
	todo, err := u.repo.InsertTodo(ctx, payload)
	if err != nil {
		return todoModels.TodoRes{}, err
	}
	return todoMapTodoRes(todo), nil
}

func (u *todoUsecase) GetTodos(ctx context.Context) ([]todoModels.TodoRes, error) {
	todos, err := u.repo.GetTodos(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]todoModels.TodoRes, 0, len(todos))
	for i := range todos {
		result = append(result, todoMapTodoRes(todos[i]))
	}
	return result, nil
}

func (u *todoUsecase) DeleteTodoById(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return todoModels.ErrParseUUIDError
	}
	rowAffectd, err := u.repo.DeleteTodoById(ctx, pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	})
	if err != nil {
		return err
	}
	if rowAffectd == 0 {
		return todoModels.ErrNotFoundTodo
	}
	return nil
}

func (u *todoUsecase) UpdatedTodoCompletedById(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return todoModels.ErrParseUUIDError
	}
	rowAffectd, err := u.repo.UpdateTodoCompletedById(ctx, pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	})
	if err != nil {
		return err
	}
	if rowAffectd == 0 {
		return todoModels.ErrNotFoundTodo
	}
	return nil
}
