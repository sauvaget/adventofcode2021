package main

import (
	"log"
	"strconv"
	"strings"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(data []string) int {
	p := &position{
		horizontal: 0,
		depth:      0,
	}
	for _, v := range data {
		step := strings.Split(v, " ")
		amount, err := strconv.Atoi(step[1])
		if err != nil {
			log.Println("something went horrible wrong")
		}
		switch step[0] {
		case "forward":
			p.horizontal = p.horizontal + amount
		case "down":
			p.depth = p.depth + amount
		case "up":
			p.depth = p.depth - amount
		}
	}

	return p.depth * p.horizontal
}

func part2(data []string) int {
	p := &position{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}
	for _, v := range data {
		step := strings.Split(v, " ")
		amount, err := strconv.Atoi(step[1])
		if err != nil {
			log.Println("something went horrible wrong")
		}
		switch step[0] {
		case "forward":
			p.horizontal = p.horizontal + amount
			p.depth = p.depth + (p.aim * amount)
		case "down":
			p.aim = p.aim + amount
		case "up":
			p.aim = p.aim - amount
		}
	}

	return p.depth * p.horizontal
}
