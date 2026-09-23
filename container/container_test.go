package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsWatchtowerLabel_ShouldReturnTrueIfTheWatchtowerLabelExistsOnTheContainer(t *testing.T) {
	labelMap := map[string]string{
		"com.centurylinklabs.watchtower": "true",
	}
	assert.True(t, ContainsWatchtowerLabel(labelMap))
}

func TestContainsWatchtowerLabel_ShouldReturnFalseIfTheWatchtowerLabelDoesntExistOnTheContainer(t *testing.T) {
	labelMap := map[string]string{
		"com.containrrr.watchtower": "lost",
	}
	assert.False(t, ContainsWatchtowerLabel(labelMap))
}

func TestContainsWatchtowerLabel_ShouldReturnFalseIfLabelsIsEmpty(t *testing.T) {
	labelMap := map[string]string{}
	assert.False(t, ContainsWatchtowerLabel(labelMap))
}