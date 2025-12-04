package main

import (
	"fmt"
)

func main() {
	println(" ")
	/* data, err := settings.CheckAndCreateSettings()
	if err != nil {
		println("Error checking/creating settings:", err.Error())
		return
	} else {
		println(data.PinnedDirs[0])
		println(" ")
		for _, dir := range data.PinnedDirs {
			println(dir)
		}
	} */
	/* dirs := files.DefoultDirs
	dir := dirs[files.Pictures] */
	/* data, err := files.SearchForDirs("testa", dir)
	if err != nil {
		println("Error searching dirs:", err.Error())
		return
	}
	for _, d := range data {
		println(d)
	} */
	/* files.SearchForDirs("tasting", "testing") */
	multiDimensionalSlice()
}

func multiDimensionalArray() {
	numbers := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// value := numbers[0][1]
	// println("Value: 1 ", value)
	numbers[0][1] = 10
	for i := range numbers {
		for j := range numbers[i] {
			println("Value:", i, j, numbers[i][j])
		}
	}
}

func multiDimensionalSlice() {
	str1 := "hat"
	str2 := "can"
	rows := 1 + len(str1)
	col := 1 + len(str2)
	slice := make([][]int, rows)
	for i := range slice {
		slice[i] = make([]int, col)
	}

	for i := range rows {
		slice[i][0] = i
	}

	for j := range col {
		slice[0][j] = j
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < col; j++ {
			if str1[i-1] == str2[j-1] {
				slice[i][j] = slice[i-1][j-1]
			} else {
				replace := slice[i-1][j-1] + 1
				dekete := slice[i-1][j] + 1
				insert := slice[i][j-1] + 1
				slice[i][j] = min(replace, dekete, insert)
			}
		}
	}
	println(str1)
	println(str2)
	editDistance := slice[rows-1][col-1]
	fmt.Println("Edit Distance:", editDistance)

	for i := range slice {
		fmt.Println(slice[i])
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
