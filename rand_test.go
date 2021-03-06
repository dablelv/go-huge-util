package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandIntRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"only get random number 0", args{0, 1}},
		{"get random number in range 0-2", args{0, 3}},
		{"get random number in range 10-100", args{10, 101}},
	}
	for _, tt := range tests {
		got := RandIntRange(tt.args.min, tt.args.max)
		valid := got >= tt.args.min && got < tt.args.max
		assert.Equal(t, true, valid, tt.name)
	}
}
