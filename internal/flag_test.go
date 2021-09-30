package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Flag(t *testing.T) {
	as := assert.New(t)

	for _, tt := range []struct {
		name string
		args []string
		key  string
		left []string
		res  string
	}{
		{name: "1", args: []string{}, key: "hi", left: []string{}, res: ""},
		{name: "1", args: []string{""}, key: "hi", left: []string{""}, res: ""},
		{name: "2", args: []string{"-hi"}, key: "hi", left: []string{}, res: "true"},
		{name: "2", args: []string{"-hi", ""}, key: "hi", left: []string{}, res: ""},
		{name: "3", args: []string{"-hi", "world"}, key: "hi", left: []string{}, res: "world"},
		{name: "3", args: []string{"-hi", "world", ""}, key: "hi", left: []string{""}, res: "world"},
		{name: "3", args: []string{"", "-hi", "world"}, key: "hi", left: []string{""}, res: "world"},
		{name: "4", args: []string{"no-key", "-hi", "world"}, key: "hi", left: []string{"no-key"}, res: "world"},
		{name: "5", args: []string{"-hi", "world", "no-key"}, key: "hi", left: []string{"no-key"}, res: "world"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			left, res := ReadFromArgs(tt.args, tt.key)
			as.Equal(tt.left, left)
			as.Equal(tt.res, res)
		})
	}
}
