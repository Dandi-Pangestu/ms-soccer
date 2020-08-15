package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPlayerTableName(t *testing.T) {
	p := Player{}
	assert.Equal(t, p.TableName(), "players")
}
