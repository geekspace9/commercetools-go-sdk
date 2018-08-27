package cterrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestError(t *testing.T) {
	err := newRequestError("Forbidden", 403)
	assert.Equal(t, "403: Forbidden", err.Error())
}
