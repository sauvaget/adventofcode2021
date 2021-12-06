package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	res := part1(data)
	assert.Equal(t, 198, res)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	res := part2(data)
	assert.Equal(t, 230, res)
}

func TestOxygenRating(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	res := OxygenRating(data)
	assert.Equal(t, 23, res)
}

func TestScrubberRating(t *testing.T) {
	t.Parallel()
	var data []string = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	res := ScrubberRating(data)
	assert.Equal(t, 10, res)
}
