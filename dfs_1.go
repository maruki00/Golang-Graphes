package main

import "fmt"

func getIslands(grid [][]int) int {
	result := 0
	visited := make([][]bool, len(grid), len(grid[0]))
	for i := range len(grid) {
		visited[i] = make([]bool, len(grid[0]))
	}

	X := []int{0, 1, 0, -1}
	Y := []int{-1, 0, 1, 0}

	var dfs func(i, j int, vidited [][]bool)
	dfs = func(i, j int, visited [][]bool) {
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 || visited[i][j] == true {
			return
		}
		visited[i][j] = true
		if grid[i][j] == 1 {
			result += 1
		}
		for k := 0; k < 4; k++ {
			dx := i + X[k]
			dy := j + Y[k]
			dfs(dx, dy, visited)
		}

	}
	dfs(0, 0, visited)
	return result
}

func main() {
	grid := [][]int{
		{0, 1, 0, 1},
		{1, 1, 1, 0},
		{0, 1, 1, 1},
		{0, 0, 1, 0},
	}
	fmt.Println("Result : ", getIslands(grid))
}
