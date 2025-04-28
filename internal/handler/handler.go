package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"temp-email-service/internal/model"
	"temp-email-service/internal/service"
	"temp-email-service/internal/utils"
)

func CreateEmailHandler(w http.ResponseWriter, r *http.Request) {
	ttlParam := r.URL.Query().Get("ttl")
	ttl := 10 * time.Minute // default

	if ttlParam != "" {
		parsedTTL, err := time.ParseDuration(ttlParam)
		if err == nil && parsedTTL > 0 {
			ttl = parsedTTL
		}
	}

	email := service.CreateTempEmail(ttl)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(email)
}

func GetInboxHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "emailID")
	email, exists := service.GetTempEmail(id)
	if !exists {
		http.Error(w, "Email not found or expired", http.StatusNotFound)
		return
	}

	response := struct {
		Address  string          `json:"address"`
		TTL      int64           `json:"ttl_seconds"`
		Messages []model.Message `json:"messages"`
	}{
		Address:  email.Address,
		TTL:      int64(email.ExpiresAt.Sub(time.Now()).Seconds()),
		Messages: email.Messages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "emailID")

	var req struct {
		From    string `json:"from"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	msg := model.Message{
		ID:      utils.GenerateID(),
		From:    req.From,
		Subject: req.Subject,
		Body:    req.Body,
		SentAt:  time.Now(),
	}

	success := service.AddMessageToEmail(id, msg)
	if !success {
		http.Error(w, "Email not found or expired", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
