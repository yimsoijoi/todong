package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/lib/store"
	"github.com/gorilla/mux"
)

func (gr *GorillaHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	userUuid := r.Header.Get("iss")

	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	var targetTodo datamodel.Todo
	if err := gr.DataGateway.GetOneTodo(ctx, &datamodel.Todo{
		UserUUID: userUuid,
		UUID:     uuid,
	}, &targetTodo); err != nil {
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
	if err := gr.DataGateway.DeleteTodo(ctx, &datamodel.Todo{
		UserUUID: userUuid,
		UUID:     uuid,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"error": "failed to delete todo",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	respEncoder.Encode(map[string]interface{}{
		"status":   "delete todo sucessful",
		"userUuid": uuid,
	})
}
