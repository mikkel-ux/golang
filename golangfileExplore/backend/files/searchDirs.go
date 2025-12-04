package files

import (
	"strings"
)

type SearchResult struct {
	Path  string
	Score int
}

func SearchForDirs(query string, dir string) {
	query = strings.ToLower(query)
	/* var results []SearchResult */
	/* entries, err := os.ReadDir(dir)
	if err != nil {
		println("Error reading directory:", err.Error())
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			name := strings.ToLower(entry.Name())
			if name == query {
				println("Found exact match:", dir+"/"+entry.Name())
			} else if strings.Contains(name, query) {
				editDistance(query, name)
			}
		}
	} */
	name := strings.ToLower(dir)
	if query == name {
		println("Found exact match:", dir)
	} else {
		editDistance(query, name)
	}

}

func editDistance(query, dir string) int {
	rows := len(query) + 1
	cols := len(dir) + 1
	dp := make([][]int, rows)
	for i := range rows {
		dp[i] = make([]int, cols)
	}

	for i := range rows {
		dp[i][0] = i
	}

	for j := range cols {
		dp[0][j] = j
	}

	for _, data := range dp {
		println(data)
	}

	return 0
}
