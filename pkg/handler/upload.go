package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/khusainnov/sbercloud/internal/entity"
)

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	m, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var d entity.UploadData
	json.Unmarshal(m, &d)

	if err = h.services.UploadConfig(d); err != nil {
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": errors.New("Cannot upload config: " + err.Error()),
		})
		return
	}

	json.NewEncoder(w).Encode(&map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Uploaded",
	})
	return
}

func (h *Handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query()
	sv := req["service"]

	rsp, err := h.services.GetConfig(sv[0])
	if err != nil {
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&map[string]interface{}{
		"code": http.StatusOK,
		"body": rsp,
	})

	return
}

func (h *Handler) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query()
	sv := req["service"]

	if err := h.services.DeleteConfig(sv[0]); err != nil {
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Deleted",
	})
	return
}
