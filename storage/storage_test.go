package storage

import (
	"testing"
	"time"
)

func TestStorage_SetAndGet(t *testing.T) {
	// Arrange
	storage := NewStorage()
	key := "username"
	value := "ritesh"
	timestamp := time.Now()

	// Act
	err := storage.Set(key, value, timestamp)
	if err != nil {
		t.Fatalf("Set returned an error: %v", err)
	}

	result := storage.Get(key)

	// Assert
	if result == nil {
		t.Fatalf("Expected value, got nil")
	}

	if result != value {
		t.Fatalf("Expected value %v, got %v", value, result)
	}
}

func TestStorage_Get_NonExistentKey(t *testing.T) {
	// Arrange
	storage := NewStorage()

	// Act
	result := storage.Get("non-existent-key")

	// Assert
	if result != nil {
		t.Fatalf("Expected nil for non-existent key, got %v", result)
	}
}
