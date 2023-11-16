package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(equalPairs([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

func equalPairs(grid [][]int) int {
	rowMap := make(map[string]int)
	columnMap := make(map[string]int)
	var result int
	var row string
	var column = make([]string, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {

			if i == len(grid)-1 {
				column[j] += strconv.Itoa(grid[i][j])
			} else {
				column[j] += strconv.Itoa(grid[i][j]) + "-"
			}

			if j == len(grid[i])-1 {
				row += strconv.Itoa(grid[i][j])
			} else {
				row += strconv.Itoa(grid[i][j]) + "-"
			}
		}
		rowMap[row]++
		row = ""
	}
	for i := range column {
		columnMap[column[i]]++
	}
	for key, row := range rowMap {
		if col, ok := columnMap[key]; ok {
			result += row * col
		}
	}
	return result
}
