package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	positions := newSet()
	guard := "^"

	grid := fromInput()
	row, col := findGuardPosition(grid, guard)

	positions.add(fmt.Sprintf("%d,%d", row, col))

	for {
		col, row = findGuardPosition(grid, guard)
		direction := getDirection(guard)
		right := getRight(guard)

		positions.add(fmt.Sprintf("%d,%d", row, col))

		if nextLeavingGrid(row, direction, col, grid) {
			break
		}

		next := grid[row+direction[0]][col+direction[1]]
		if next == "#" {
			grid[row][col] = right
			guard = right
		} else {
			grid[row][col] = "o"
			grid[row+direction[0]][col+direction[1]] = guard
		}
	}

	fmt.Println(positions.size())
}

func fromInput() [][]string {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	grid := make([][]string, len(rows))
	for i, row := range rows {
		grid[i] = strings.Split(row, "")
	}
	return grid
}

func findGuardPosition(grid [][]string, guard string) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == guard {
				return i, j
			}
		}
	}
	return -1, -1
}

type set struct {
	m map[string]bool
}

func newSet() *set {
	return &set{m: make(map[string]bool)}
}

func (s *set) add(v string) {
	s.m[v] = true
}

func (s *set) contains(v string) bool {
	_, ok := s.m[v]
	return ok
}

func (s *set) size() int {
	return len(s.m)
}

func getDirection(s string) [2]int {
	switch s {
	case "^":
		return [2]int{-1, 0}
	case "v":
		return [2]int{1, 0}
	case "<":
		return [2]int{0, -1}
	case ">":
		return [2]int{0, 1}
	default:
		panic("Invalid direction")
	}
}

func getRight(c string) string {
	switch c {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	default:
		return "^"
	}
}

func nextLeavingGrid(row int, direction [2]int, col int, grid [][]string) bool {
	return row+direction[0] < 0 || row+direction[0] >= len(grid) || col+direction[1] < 0 || col+direction[1] >= len(grid[row])
}
