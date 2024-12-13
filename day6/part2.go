package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var Reset = "\033[0m"
var Green = "\033[32m"

func main() {
	grid := newGrid(input, "^")
	fmt.Println(grid.uniqueVisits())
}

func (g *Grid) uniqueVisits() int {
	for {
		currentPosition := g.findGuard()
		right := g.getNextGuardTurn()
		moveDirection := g.getNextGuardMove()

		g.setVisited(currentPosition)

		if g.nextMoveOutOfBound(currentPosition, moveDirection) {
			break
		}

		next := g.get(currentPosition, moveDirection)
		if next == "#" {
			g.set(currentPosition, right)
			g.guard = right
		} else {
			g.set(currentPosition, fmt.Sprintf("%s%s%s", Green, "o", Reset))
			g.set(g.next(currentPosition, moveDirection), g.guard)
		}
	}

	return len(g.visited)
}

type Grid struct {
	area    [][]string
	visited map[string]int
	guard   string
}

func newGrid(input string, guard string) *Grid {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	area := make([][]string, len(rows))
	for i, row := range rows {
		area[i] = strings.Split(row, "")
	}

	return &Grid{area, make(map[string]int), guard}
}

func (g *Grid) print() {
	fmt.Printf("\033[0;0H")
	fmt.Println(g.string())
}

func (g *Grid) string() string {
	var sb strings.Builder
	for _, row := range g.area {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g *Grid) findGuard() coordinate {
	for i, row := range g.area {
		for j, cell := range row {
			if cell == g.guard {
				return coordinate{i, j}
			}
		}
	}

	return coordinate{-1, -1}
}

func (g *Grid) nextMoveOutOfBound(position coordinate, direction coordinate) bool {
	rowOutBound := position.x+direction.x < 0 || position.x+direction.x >= len(g.area[0])
	colOutBound := position.y+direction.y < 0 || position.y+direction.y >= len(g.area)

	return rowOutBound || colOutBound
}

func (g *Grid) get(position coordinate, direction coordinate) string {
	return g.area[position.y+direction.y][position.x+direction.x]
}

func (g *Grid) set(position coordinate, value string) {
	g.area[position.y][position.x] = value
}

func (g *Grid) next(position coordinate, direction coordinate) coordinate {
	return coordinate{position.y + direction.y, position.x + direction.x}
}

func (g *Grid) setVisited(position coordinate) {
	g.visited[position.string()] = g.visited[position.string()] + 1
}

func (g *Grid) getNextGuardTurn() string {
	switch g.guard {
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

func (g *Grid) getNextGuardMove() coordinate {
	switch g.guard {
	case "^":
		return coordinate{-1, 0}
	case ">":
		return coordinate{0, 1}
	case "v":
		return coordinate{1, 0}
	default:
		return coordinate{0, -1}
	}
}

type coordinate struct {
	y, x int
}

func (p coordinate) string() string {
	return fmt.Sprintf("%d,%d", p.y, p.x)
}
