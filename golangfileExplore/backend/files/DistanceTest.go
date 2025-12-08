package files

import (
	"fmt"
	"sort"
)

type DistanceTest struct{}

func NewDistanceTest() *DistanceTest {
	return &DistanceTest{}
}

type EditDistanceReturn struct {
	target   string
	distance int
	score    int
}

func (d *DistanceTest) DistanceTests(query string) []string {

	result := []EditDistanceReturn{}
	targets := []string{"moonlight", "starlace", "shadowveil", "crimsonrose", "snowtrace",
		"sunmirror", "bloodpetal", "stormbloom", "neonvortex", "dusthalo", "emberglow", "frostflare", "ghostwhisper", "ironshade", "silvershard"}

	for _, target := range targets {
		distance := LevenshteinDistance(query, target)
		result = append(result, EditDistanceReturn{target: target, distance: distance.distance, score: distance.score})
	}

	sort.Slice(result, func(i, j int) bool {
		/* return result[i].distance < result[j].distance */
		return result[i].distance < result[j].distance
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].score > result[j].score
	})

	returnValues := []string{}
	for _, res := range result {
		returnValues = append(returnValues,
			fmt.Sprintf("Edit distance between '%s' and '%s' is %d", query, res.target, res.distance))
	}

	return returnValues
}

func LevenshteinDistance(query, target string) EditDistanceReturn {
	row := 1 + len(query)
	col := 1 + len(target)
	slice := make([][]int, row)
	score := 0
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
				score++
			} else {
				replace := slice[i-1][j-1] + 1
				dekete := slice[i-1][j] + 1
				insert := slice[i][j-1] + 1
				slice[i][j] = min(replace, dekete, insert)
			}
		}
	}
	/* for i := range row {
		fmt.Println(slice[i], score)
	} */
	if slice[row-1][col-1] > 10 {
		return EditDistanceReturn{
			target:   target,
			distance: 9999,
			score:    score,
		}
	} else {
		return EditDistanceReturn{
			target:   target,
			distance: slice[row-1][col-1],
			score:    score,
		}
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
