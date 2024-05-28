package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	mockR := NewMockRepository(t)
	err := mockR.Start()
	assert.Equal(t, nil, err)
	err = mockR.Stop()
	assert.Equal(t, nil, err)
}
