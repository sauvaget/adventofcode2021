package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	res := part1(data)
	assert.Equal(t, 150, res)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	res := part2(data)
	assert.Equal(t, 900, res)
}
