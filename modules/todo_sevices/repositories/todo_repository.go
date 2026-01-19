package todoRepository

import (
	"context"
	"fmt"
	todoModels "recap-golang-basic/modules/todo_sevices/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type (
	todoResitory struct {
		*Queries
	}
)

func (t todoResitory) DeleteTodoById(ctx context.Context, id pgtype.UUID) (int64, error) {
	nctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	rows, err := t.Queries.DeleteTodoById(nctx, id)
	if err != nil {
		return 0, fmt.Errorf("todoResitory.DeleteTodoById: %w", err)
	}
	if rows == 0 {
		return 0, todoModels.ErrNotFoundTodo
	}
	return rows, nil
}

func (t todoResitory) GetTodos(ctx context.Context) ([]Todo, error) {
	nctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	todos, err := t.Queries.GetTodos(nctx)
	if err != nil {
		return nil, fmt.Errorf("todoResitory.GetTodos: %w", err)
	}
	return todos, nil
}

func (t todoResitory) InsertTodo(ctx context.Context, arg InsertTodoParams) (Todo, error) {
	nctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	todo, err := t.Queries.InsertTodo(nctx, arg)
	var empty Todo
	if err != nil {
		return empty, fmt.Errorf("todoResitory.InsertTodo: %w", err)
	}
	return todo, nil
}

func (t todoResitory) UpdateTodoCompletedById(ctx context.Context, id pgtype.UUID) (int64, error) {
	nctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	rows, err := t.Queries.UpdateTodoCompletedById(nctx, id)
	if err != nil {
		return 0, fmt.Errorf("todoResitory.UpdateTodoCompletedById: %w", err)
	}
	if rows == 0 {
		return 0, todoModels.ErrNotFoundTodo
	}
	return rows, nil
}

func NewTodoRository(db DBTX) Querier {
	return &todoResitory{
		Queries: New(db),
	}
}
