package service

import (
	"sync"
	"time"

	"temp-email-service/internal/model"
	"temp-email-service/internal/utils"
)

var (
	emails = make(map[string]*model.TempEmail)
	mu     sync.RWMutex
)

func CreateTempEmail(ttl time.Duration) *model.TempEmail {
	id := utils.GenerateID()
	email := &model.TempEmail{
		ID:        id,
		Address:   id + "@tempdomain.com",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
		Messages:  []model.Message{},
	}

	mu.Lock()
	emails[id] = email
	mu.Unlock()

	return email
}

func GetTempEmail(id string) (*model.TempEmail, bool) {
	mu.RLock()
	defer mu.RUnlock()

	email, exists := emails[id]
	if !exists || time.Now().After(email.ExpiresAt) {
		return nil, false
	}
	return email, true
}

func AddMessageToEmail(id string, msg model.Message) bool {
	mu.Lock()
	defer mu.Unlock()

	email, exists := emails[id]
	if !exists || time.Now().After(email.ExpiresAt) {
		return false
	}

	email.Messages = append(email.Messages, msg)
	return true
}

func CleanExpiredEmails() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	for id, email := range emails {
		if now.After(email.ExpiresAt) {
			delete(emails, id)
		}
	}
}
