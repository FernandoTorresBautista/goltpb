package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	mApi := NewMockApiHandle(t)
	mCtx, _ := context.WithCancel(context.Background())
	err := mApi.Run(mCtx)
	assert.NoError(t, err, "should be nil")

	err = mApi.TurnOff()
	assert.NoError(t, err, "should be nil")
}
