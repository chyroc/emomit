package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Key  string
	Text string
	In   bool
}

func Test_InText(t *testing.T) {
	cases := []Case{
		{
			Key:  "a",
			Text: "aa",
			In:   true,
		},
		{
			Key:  "abc",
			Text: "aabbcc",
			In:   true,
		},
		{
			Key:  "ab",
			Text: "bbaa",
			In:   false,
		},
		{
			Key:  "你好",
			Text: "我叫你，请问好不好？",
			In:   true,
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.In, InText(v.Key, v.Text), fmt.Sprintf("%#v", v))
	}
}
