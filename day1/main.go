package main

import "log"

func main() {
	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(data []int) int {
	var counter int = 0
	for k, v := range data {
		if k != 0 {
			if v > data[k-1] {
				counter++
			}
		}
	}
	return counter
}

func part2(data []int) int {
	newinput := make([]int, 0)
	for k, v := range data {
		if k+2 < len(data) {
			newinput = append(newinput, (v + data[k+1] + data[k+2]))
		}
	}

	return part1(newinput)
}
