package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_UpdateCost(t *testing.T) {
	q := NewStripedQueue[int, int](1, 10, func() int32 { return -1 })
	entry := &Entry[int, int]{cost: -1}
	entry.queued = 1
	q.Push(20, entry, 1, false)
	require.Equal(t, 1, q.qs[0].len)

	q.UpdateCost(20, entry, 5)
	require.Equal(t, 5, q.qs[0].len)

	q.UpdateCost(20, entry, 3)
	require.Equal(t, 3, q.qs[0].len)
}
