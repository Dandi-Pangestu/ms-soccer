package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTeamTableName(t *testing.T) {
	team := Team{}
	assert.Equal(t, team.TableName(), "teams")
}
