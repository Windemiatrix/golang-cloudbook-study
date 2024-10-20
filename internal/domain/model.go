package domain

// KeyValue is a simple domain model representing a key-value pair.
type KeyValue struct {
	Key   string
	Value string
}

// Repository is an interface for a storage service.
type Repository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
