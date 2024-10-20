package storage

import (
	"errors"
	"sync"
)

// InMemoryRepository is an in-memory implementation of the Repository interface.
type InMemoryRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewInMemoryRepository creates a new instance of InMemoryRepository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

// Get retrieves the value for a given key.
func (repo *InMemoryRepository) Get(key string) (string, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if value, exists := repo.data[key]; exists {
		return value, nil
	}
	return "", errors.New("key not found")
}

// Set stores the key-value pair in the repository.
func (repo *InMemoryRepository) Set(key string, value string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.data[key] = value
	return nil
}

// Delete removes the key-value pair from the repository.
func (repo *InMemoryRepository) Delete(key string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.data[key]; exists {
		delete(repo.data, key)
		return nil
	}
	return errors.New("key not found")
}
