package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(data []string) int {
	var cols = make(map[int]string)
	for _, v := range data {
		for k, a := range v {
			cols[k] += string(a)
		}
	}

	var gamma string
	var epsilon string

	for i := 0; i < len(data[0]); i++ {
		c := countCol(cols[i])

		if c == "1" {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
		if c == "0" {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaDec) * int(epsilonDec)
}

func countCol(col string) string {
	a := strings.Count(col, "0")
	b := strings.Count(col, "1")

	if a < b {
		return "1"
	}
	if a > b {
		return "0"
	}
	return ""
}

func part2(data []string) int {
	oxygen := OxygenRating(data)
	scrubber := ScrubberRating(data)
	return oxygen * scrubber
}

func ScrubberRating(data []string) int {
	scrubberrating := calcScrubberRating(data, 0)
	scrubberratingDec, _ := strconv.ParseInt(scrubberrating[0], 2, 64)
	return int(scrubberratingDec)
}

func OxygenRating(data []string) int {
	oxyrating := calcOxygenRating(data, 0)
	oxyratingDec, _ := strconv.ParseInt(oxyrating[0], 2, 64)
	return int(oxyratingDec)
}

func calcScrubberRating(data []string, depth int) []string {
	if len(data) == 1 {
		return data
	}
	newdata := make([]string, 0)
	cols := convertToCols(data)
	common := countCol(cols[depth])
	if common == "" {
		common = "1"
	}
	for _, bin := range data {
		if bin[depth:depth+1] != common {
			newdata = append(newdata, bin)
		}
	}
	return calcScrubberRating(newdata, depth+1)
}

func calcOxygenRating(data []string, depth int) []string {
	if len(data) == 1 {
		return data
	}
	newdata := make([]string, 0)
	cols := convertToCols(data)
	common := countCol(cols[depth])
	if common == "" {
		common = "1"
	}
	for _, bin := range data {
		if bin[depth:depth+1] == common {
			newdata = append(newdata, bin)
		}
	}
	return calcOxygenRating(newdata, depth+1)
}

func convertToCols(data []string) map[int]string {
	var cols = make(map[int]string)
	for _, v := range data {
		for k, a := range v {
			cols[k] += string(a)
		}
	}
	return cols
}
