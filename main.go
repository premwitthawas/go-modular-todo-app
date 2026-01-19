package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	todoHandler "recap-golang-basic/modules/todo_sevices/handlers"
	todoRepository "recap-golang-basic/modules/todo_sevices/repositories"
	todoUsecase "recap-golang-basic/modules/todo_sevices/usecases"
	databasePkg "recap-golang-basic/pkgs/database"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	pool := databasePkg.GetPoolDb()
	defer func() {
		log.Println("Database pool is closing...")
		pool.Close()
	}()
	validate := validator.New()
	app := gin.New()
	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.Handler(),
	}

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	api := app.Group("/api/v1")

	todoRepo := todoRepository.NewTodoRository(pool)
	todoUsecase := todoUsecase.NewTodousecase(todoRepo)
	todoHanler := todoHandler.NewTodoHandler(todoUsecase, validate)

	todos := api.Group("/todos")
	todos.POST("", todoHanler.CreatTodo)
	todos.GET("", todoHanler.GetTodos)
	todos.PATCH("/:id", todoHanler.UpdatedCompleteTodoById)
	todos.DELETE("/:id", todoHanler.DeleteTodoById)

	ch := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error: Server Listening: %s \n", err.Error())
		}
	}()
	log.Println("info: Server Listening :4000")
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Println("info: Server Shutting down ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("error: Server Shutdown error: ", err)
	}
	log.Println("info: Server Shutdown completed.")
}
