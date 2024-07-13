package main

import "fmt"

func getIslands(grid [][]int) int {
	result := 0
	visited := make([][]bool, len(grid), len(grid[0]))
	for i := range len(grid) {
		visited[i] = make([]bool, len(grid[0]))
	}

	var bfs func(i, j int, vidited [][]bool)
	bfs = func(i, j int, visited [][]bool) {
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 || visited[i][j] == true {
			return
		}
		visited[i][j] = true
		if grid[i][j] == 1 {
			result += 1
		}
		bfs(i+1, j, visited)
		bfs(i-1, j, visited)
		bfs(i, j+1, visited)
		bfs(i, j-1, visited)
	}
	bfs(0, 0, visited)
	return result
}

func main() {
	grid := [][]int{
zsh:1: command not found: :w
		{0, 1, 1, 1},
		{0, 0, 1, 0},
	}
	fmt.Println("Result : ", getIslands(grid))
}
