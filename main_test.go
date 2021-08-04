package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, World!!!!"
	if want != hello() {
		t.Error("失敗しました。")
	}
}

func TestSquere(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  []int
		output int
	}{
		{
			title:  "2×3の答えが6になる",
			input:  []int{2, 3},
			output: 6,
		},
		{
			title:  "3×5の答えが15になる",
			input:  []int{3, 5},
			output: 15,
		},
		{
			title:  "0×1の答えが0になる",
			input:  []int{0, 1},
			output: 0,
		},
	} {
		t.Run("Square:"+td.title, func(t *testing.T) {
			result := square(td.input[0], td.input[1])
			asserts.Equal(td.output, result)
		})
	}
}
