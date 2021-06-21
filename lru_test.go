package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutFirstNode(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	assert.Len(t, lruN.bucket, 1)
}
