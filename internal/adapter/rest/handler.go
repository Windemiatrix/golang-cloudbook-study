package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Windemiatrix/golang-cloudbook-study/internal/domain"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repo domain.Repository
}

func NewHandler(repo domain.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetKeyValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := h.Repo.Get(key)
	if err != nil {
		logrus.WithError(err).Error("Failed to get value")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	logrus.WithFields(logrus.Fields{"key": key, "value": value}).Info("Value retrieved")
	if err := json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value}); err != nil {
		logrus.WithError(err).Error("Failed to encode response")
	}
}

func (h *Handler) SetKeyValue(w http.ResponseWriter, r *http.Request) {
	var kv domain.KeyValue
	if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
		logrus.WithError(err).Error("Failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.Set(kv.Key, kv.Value); err != nil {
		logrus.WithError(err).Error("Failed to set key-value")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.WithFields(logrus.Fields{"key": kv.Key, "value": kv.Value}).Info("Key-value set")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteKeyValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if err := h.Repo.Delete(key); err != nil {
		logrus.WithError(err).Error("Failed to delete key")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	logrus.WithField("key", key).Info("Key deleted")
	w.WriteHeader(http.StatusNoContent)
}
