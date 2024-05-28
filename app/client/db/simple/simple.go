package simple

import (
	"sync"
)

// SimpleRepo simulates a repository storing data in memory
type SimpleRepo struct {
	mu     sync.RWMutex
	nextID uint64
}

// NewSimpleRepo creates a new instance of SimpleRepo
func NewSimpleRepo() *SimpleRepo {
	return &SimpleRepo{
		nextID: 1,
	}
}

// Start function
func (m *SimpleRepo) Start() error {
	return nil
}

// Stop function
func (m *SimpleRepo) Stop() error {
	return nil
}

// implement functions
