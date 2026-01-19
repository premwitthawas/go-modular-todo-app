package todoUsecase

import (
	"context"
	todoModels "recap-golang-basic/modules/todo_sevices/models"
	todoRepository "recap-golang-basic/modules/todo_sevices/repositories"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestInsertTodo_Usecase(t *testing.T) {
	uuid, _ := uuid.Parse("3a34363e-06e9-4f2c-b3e1-1450fd133351")
	mockRepo := NewMockTodoRepo()
	uc := NewTodousecase(mockRepo)
	t.Run("Success", func(t *testing.T) {
		req := &todoModels.TodoCreateReq{
			Title: "Testing",
			Email: "example@test.com",
		}
		insertPayload := todoRepository.InsertTodoParams{
			Title: "Testing",
			Email: "example@test.com",
		}
		exectDataTodo := todoRepository.Todo{
			ID: pgtype.UUID{
				Bytes: uuid,
				Valid: true,
			},
			Title: "Testing",
			Email: "example@test.com",
			Done: pgtype.Bool{
				Bool:  false,
				Valid: true,
			},
			CreatedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
		}
		mockRepo.On("InsertTodo", context.Background(), insertPayload).Return(exectDataTodo, nil).Once()
		res, err := uc.CreateTodo(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, uuid.String(), res.ID)
		assert.Equal(t, "Testing", res.Title)
		assert.Equal(t, "example@test.com", res.Email)
		mockRepo.AssertExpectations(t)
	})
}
