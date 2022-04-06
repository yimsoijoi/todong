package gorillahandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
)

func (gr *GorillaHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("here 2 " + err.Error()))
		return
	}
	defer file.Close()
	respEncoder := json.NewEncoder(w)
	imgBuf := new(bytes.Buffer)
	if _, err := io.Copy(imgBuf, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to read file",
			"error":  err.Error(),
		})
		return
	}
	jsonBody := r.FormValue("data")
	if len(jsonBody) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "empty multipart/form-data key \"data\"",
		})
		return
	}

	var req internal.TodoReqBody
	if err := json.Unmarshal([]byte(jsonBody), &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "bad JSON body in multipart/form-data key \"data\"",
		})
		return
	}
	userUuid := r.Header.Get("iss")
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"status": "image file to large",
		})
		return
	}
	order, err := datamodel.NewTodo(userUuid, req.Title, req.Description, req.TodoDate, enums.Status(req.Status), imgBase64Str)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to create new todo (1)",
			"error":  err.Error(),
		})
		return
	}
	if err := gr.DataGateway.CreateTodo(r.Context(), order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to create new todo (2)",
			"error":  err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = respEncoder.Encode(order)
}
