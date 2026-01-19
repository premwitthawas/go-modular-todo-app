-- name: InsertTodo :one
INSERT INTO todos (title,email)
VALUES ($1,$2)
RETURNING *;

-- name: GetTodos :many
SELECT *
FROM todos;

-- name: UpdateTodoCompletedById :execrows
UPDATE todos
SET done = true
WHERE id = $1;

-- name: DeleteTodoById :execrows
DELETE FROM todos
WHERE id = $1;
