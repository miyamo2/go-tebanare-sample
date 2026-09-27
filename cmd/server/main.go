// Command server wires the domain, usecase, repository, and handler layers
// together and serves the task API over HTTP.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/miyamo2/go-tebanare-sample/handler"
	"github.com/miyamo2/go-tebanare-sample/infra/logger"
	"github.com/miyamo2/go-tebanare-sample/repository"
	"github.com/miyamo2/go-tebanare-sample/usecase"
)

func main() {
	repo := repository.NewInMemoryTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(repo, logger.Noop{})

	seed, err := taskUseCase.CreateTask("task-1", "proj-1", "Write the sample README", time.Now().Add(48*time.Hour))
	if err != nil {
		log.Fatalf("seed task: %v", err)
	}
	log.Printf("seeded %s: %s", seed.ID(), seed.Title())

	recent := usecase.NewRecentlyViewedUseCase()
	recent.Touch(seed.ID())
	log.Printf("recently viewed: %s (count=%d)", recent.Last(), recent.Count())

	taskHandler := handler.NewTaskHandler(taskUseCase)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks/{id}", taskHandler.Get)
	mux.HandleFunc("POST /tasks/{id}/complete", taskHandler.Complete)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
