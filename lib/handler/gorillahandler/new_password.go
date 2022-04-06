package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
	"github.com/gorilla/mux"
)

func (gr *GorillaHandler) NewPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	userUuid := r.Header.Get("iss")

	respEncoder := json.NewEncoder(w)
	ctx := r.Context()

	var newPassReq internal.NewPasswordJson
	if err := json.NewDecoder(r.Body).Decode(&newPassReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(newPassReq.NewPassword) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"error": "blank password received",
		})
		return
	}

	var targetUser datamodel.User
	if err := gr.DataGateway.GetUserByUuid(ctx, userUuid, &targetUser); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			respEncoder.Encode(map[string]interface{}{
				"status":   "users not found",
				"userUuid": uuid,
				"error":    err.Error(),
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
	pw, err := utils.EncodeBcrypt([]byte(newPassReq.NewPassword))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"status": "failed to encode password",
		})
		return
	}
	if err := gr.DataGateway.ChangePassword(ctx, &targetUser, pw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"status": "failed to change password",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	respEncoder.Encode(map[string]interface{}{
		"status":   "todo update sucessful",
		"userUuid": uuid,
	})
}
