package gorillahandler

import (
	"encoding/json"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/utils"
	"github.com/google/uuid"
)

func (gr *GorillaHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req internal.AuthJson
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respEncoder := json.NewEncoder(w)
	// Also, return after calling this func
	registerFailed := func() {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"error": "register failed",
		})
	}

	var user datamodel.User
	ctx := r.Context()
	_ = gr.DataGateway.GetUserByUsername(ctx, req.Username, &user)

	if user.Username != "" {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"duplicate username": req.Username,
		})
		return
	}
	pw, err := utils.EncodeBcrypt([]byte(req.Password))
	if err != nil {
		registerFailed()
		return
	}
	user = datamodel.User{
		UUID:     uuid.NewString(),
		Username: req.Username,
		Password: pw,
	}
	if err := gr.DataGateway.CreateUser(ctx, &user); err != nil {
		registerFailed()
		return
	}
	w.WriteHeader(http.StatusCreated)
	respEncoder.Encode(map[string]interface{}{
		"username": user.Username,
		"userUuid": user.UUID,
	})
}
