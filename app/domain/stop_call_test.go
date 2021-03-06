package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStopCall(t *testing.T) {
	n := time.Now()

	tables := []StopCall{
		{"R1", n, "C1"},
		{"R1123", n, "asasdasd"},
		{"999", n, "23123132@#@!@@"},
		{"aaAAA321312", n, "C1"},
	}

	for _, table := range tables {
		r, _ := NewStopCall(table.RecordID, table.Timestamp, table.CallID)
		assert.NotNil(t, r, "Must be a valid object!")
		assert.Equal(t, r.RecordID, table.RecordID)
		assert.Equal(t, r.Timestamp, table.Timestamp)
		assert.Equal(t, r.CallID, table.CallID)
	}
}

func TestInvalidStopCallWhenSomeFieldIsEmpty(t *testing.T) {
	n := time.Now()

	tables := []StopCall{
		{"", n, "C1"},
		{"R1", time.Time{}, "C1"},
		{"R1", n, ""},
	}

	for _, table := range tables {
		r, err := NewStopCall(table.RecordID, table.Timestamp, table.CallID)
		assert.Nil(t, r, "Must be a invalid object!")
		assert.EqualError(t, err, "empty fields")
	}
}
