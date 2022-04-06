package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/lib/store"
	"github.com/gorilla/mux"
)

func (gr *GorillaHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	userUuid := r.Header.Get("iss")

	var getAll bool
	if len(uuid) == 0 {
		getAll = true
	}
	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	var todos []datamodel.Todo
	if getAll {
		if err := gr.DataGateway.GetUserTodos(ctx, &datamodel.Todo{
			UserUUID: userUuid,
		}, &todos); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				respEncoder.Encode(map[string]interface{}{
					"status":   "todos not found",
					"userUuid": userUuid,
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			respEncoder.Encode(map[string]interface{}{
				"status": "failed to find todo",
				"uuid":   uuid,
				"error":  err.Error(),
			})
			return
		}
	} else {
		var todo datamodel.Todo
		if err := gr.DataGateway.GetOneTodo(ctx, &datamodel.Todo{
			UserUUID: userUuid,
			UUID:     uuid,
		}, &todo); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				respEncoder.Encode(map[string]interface{}{
					"status": "todos not found",
					"uuid":   uuid,
					"error":  err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			respEncoder.Encode(map[string]interface{}{
				"status": "failed to find todo",
				"uuid":   uuid,
				"error":  err.Error(),
			})
			return
		}
		todos = append(todos, todo)
	}
	w.WriteHeader(http.StatusOK)
	respEncoder.Encode(todos)
}
