package gorillahandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yimsoijoi/todong/datamodel"
)

func (gr *GorillaHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]

	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	if err := gr.DataGateway.DeleteUser(ctx, &datamodel.User{
		UUID: uuid,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"error": "failed to delete user",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	respEncoder.Encode(map[string]interface{}{
		"status":   "delete user sucessful",
		"userUuid": uuid,
	})
}
