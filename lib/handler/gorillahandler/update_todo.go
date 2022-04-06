package gorillahandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
	"github.com/yimsoijoi/todong/lib/store"
)

func (gr *GorillaHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
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

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("here 2 " + err.Error()))
		return
	}
	defer file.Close()
	imgBuf := new(bytes.Buffer)
	if _, err := io.Copy(imgBuf, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"status": "failed to read file",
			"error":  err.Error(),
		})
		return
	}
	jsonBody := r.FormValue("data")
	if len(jsonBody) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"status": "empty multipart/form-data key \"data\"",
		})
		return
	}

	var updatesReq internal.TodoReqBody
	if err := json.Unmarshal([]byte(jsonBody), &updatesReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"status": "bad JSON body in multipart/form-data key \"data\"",
		})
		return
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"status": "image file to large",
		})
		return
	}

	compareVal := func(old, new string, target *string) {
		var nullString string
		if new == nullString {
			*target = old
			return
		}
		*target = new
	}
	compareStatus := func(old, new enums.Status, target *enums.Status) {
		if new.IsValid() {
			*target = new
		}
		*target = old
	}
	var u datamodel.Todo // Updated to-do
	compareVal(targetTodo.UUID, uuid, &u.UUID)
	compareVal(targetTodo.UserUUID, "", &u.UserUUID)
	compareVal(targetTodo.Title, updatesReq.Title, &u.Title)
	compareVal(targetTodo.Description, updatesReq.Description, &u.Description)
	compareVal(targetTodo.TodoDate, updatesReq.TodoDate, &u.TodoDate)
	compareVal(targetTodo.Image, imgBase64Str, &u.Image)
	compareStatus(targetTodo.Status, enums.Status(updatesReq.Status), &u.Status)

	if err := gr.DataGateway.UpdateTodo(ctx, &datamodel.Todo{
		UserUUID: userUuid,
		UUID:     uuid,
	}, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"status": "failed to update todo",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	respEncoder.Encode(map[string]interface{}{
		"status":   "todo update sucessful",
		"uuid":     u.UUID,
		"userUuid": u.UserUUID,
	})
}
