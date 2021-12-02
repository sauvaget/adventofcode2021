package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	var data []int = []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	res := part1(data)
	assert.Equal(t, 7, res)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	var data []int = []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	res := part2(data)
	assert.Equal(t, 5, res)
}
