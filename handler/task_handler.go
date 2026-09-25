package handler

import (
	"encoding/json"
	"net/http"

	"github.com/miyamo2/go-tebanare-sample/domain"
	"github.com/miyamo2/go-tebanare-sample/usecase"
)

// TaskHandler adapts TaskUseCase to net/http.
type TaskHandler struct {
	uc *usecase.TaskUseCase
}

// NewTaskHandler creates a TaskHandler backed by uc.
func NewTaskHandler(uc *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{uc: uc}
}

// Get handles GET /tasks/{id}.
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	t, err := h.uc.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeTask(w, t)
}

// Complete handles POST /tasks/{id}/complete.
func (h *TaskHandler) Complete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.uc.CompleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Delete handles DELETE /tasks/{id}.
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.uc.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeTask(w http.ResponseWriter, t *domain.Task) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"id":        t.ID(),
		"projectId": t.ProjectID(),
		"title":     t.Title(),
		"status":    t.Status(),
		"assignee":  t.AssigneeID(),
		"dueAt":     t.DueAt(),
	})
}
