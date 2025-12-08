package main

import (
	"fmt"
	"sort"
)

func main() {
	println(" ")
	input := "cat"
	result := []EditDistanceReturn{}
	targets := []string{"cut", "cot", "cast", "at", "acts", "cost", "scat", "cat"}

	for _, target := range targets {
		distance := LevenshteinDistance(input, target)
		result = append(result, distance)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})

	for _, res := range result {
		fmt.Printf("Edit distance between '%s' and '%s' is %d\n", input, res.target, res.distance)
	}
}

type EditDistanceReturn struct {
	target   string
	distance int
}

func LevenshteinDistance(query, target string) EditDistanceReturn {
	row := 1 + len(query)
	col := 1 + len(target)
	slice := make([][]int, row)
	for i := range slice {
		slice[i] = make([]int, col)
	}

	for i := range row {
		slice[i][0] = i
	}

	for j := range col {
		slice[0][j] = j
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if query[i-1] == target[j-1] {
				slice[i][j] = slice[i-1][j-1]
			} else {
				replace := slice[i-1][j-1] + 1
				dekete := slice[i-1][j] + 1
				insert := slice[i][j-1] + 1
				slice[i][j] = min(replace, dekete, insert)
			}
		}
	}
	/* for i := range row {
		fmt.Println(slice[i])
	} */
	return EditDistanceReturn{
		target:   target,
		distance: slice[row-1][col-1],
	}
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}
