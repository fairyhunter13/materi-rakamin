package oop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTryOOP(t *testing.T) {
	// Don't modify this
	doc := NewDocument()
	doc.SetState(doc.GetStateFromMap("draft"))

	// initial check
	require.IsType(t, new(Draft), doc.GetState())

	doc.Progress()
	require.IsType(t, new(Completed), doc.GetState())

	doc.Submit()
	require.IsType(t, new(Submitted), doc.GetState())

	doc.Reject()
	require.IsType(t, new(Rejected), doc.GetState())

	doc.Progress()
	require.IsType(t, new(Completed), doc.GetState())

	doc.Submit()
	require.IsType(t, new(Submitted), doc.GetState())

	doc.Approve()
	require.IsType(t, new(Approved), doc.GetState())

	doc.Publish()
	require.IsType(t, new(Published), doc.GetState())
}
