package todoUsecase

import (
	"context"
	todoRepository "recap-golang-basic/modules/todo_sevices/repositories"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type mockTodoRepo struct {
	mock.Mock
}

func NewMockTodoRepo() *mockTodoRepo {
	return &mockTodoRepo{}
}

func (m *mockTodoRepo) GetTodos(ctx context.Context) ([]todoRepository.Todo, error) {
	args := m.Called(ctx)
	var todos []todoRepository.Todo
	if args.Get(0) != nil {
		todos = args.Get(0).([]todoRepository.Todo)
	}
	return todos, args.Error(1)
}

func (m *mockTodoRepo) DeleteTodoById(ctx context.Context, id pgtype.UUID) (int64, error) {
	args := m.Called(ctx, id)
	return int64(args.Int(0)), args.Error(1)
}

func (m *mockTodoRepo) UpdateTodoCompletedById(ctx context.Context, id pgtype.UUID) (int64, error) {
	args := m.Called(ctx, id)
	return int64(args.Int(0)), args.Error(1)
}

func (m *mockTodoRepo) InsertTodo(ctx context.Context, arg todoRepository.InsertTodoParams) (todoRepository.Todo, error) {
	args := m.Called(ctx, arg)
	var todo todoRepository.Todo
	if args.Get(0) != nil {
		todo = args.Get(0).(todoRepository.Todo)
	}
	return todo, args.Error(1)
}
