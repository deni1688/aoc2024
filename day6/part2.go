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
	grid := newGrid(input, newGuard("^"))
	fmt.Println(grid.uniqueVisits())
}

func (g *Grid) uniqueVisits() int {
	for {
		currentPosition := g.findGuard()
		right := g.guard.getRightTurn()
		move := g.guard.getNextMove()

		g.setVisited(currentPosition)

		if g.nextMoveOutOfBound(currentPosition, move) {
			break
		}

		next := g.get(currentPosition, move)
		if next == "#" {
			g.set(currentPosition, right.string())
			g.guard = right
		} else {
			g.set(currentPosition, fmt.Sprintf("%s%s%s", Green, "o", Reset))
			g.set(g.next(currentPosition, move), g.guard.string())
		}
	}

	return len(g.visited)
}

type Grid struct {
	area    [][]string
	visited map[string]int
	guard   *Guard
}

func newGrid(input string, guard *Guard) *Grid {
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

func (g *Grid) findGuard() Position {
	for i, row := range g.area {
		for j, cell := range row {
			if g.guard.equals(cell) {
				return Position{i, j}
			}
		}
	}

	return Position{-1, -1}
}

func (g *Grid) nextMoveOutOfBound(position Position, direction Position) bool {
	rowOutBound := position.x+direction.x < 0 || position.x+direction.x >= len(g.area[0])
	colOutBound := position.y+direction.y < 0 || position.y+direction.y >= len(g.area)

	return rowOutBound || colOutBound
}

func (g *Grid) get(position Position, direction Position) string {
	return g.area[position.y+direction.y][position.x+direction.x]
}

func (g *Grid) set(position Position, value string) {
	g.area[position.y][position.x] = value
}

func (g *Grid) next(position Position, direction Position) Position {
	return Position{position.y + direction.y, position.x + direction.x}
}

func (g *Grid) setVisited(position Position) {
	g.visited[position.string()] = g.visited[position.string()] + 1
}

type Guard string

func newGuard(s string) *Guard {
	return (*Guard)(&s)
}

func (g *Guard) string() string {
	return string(*g)
}

func (g *Guard) getRightTurn() *Guard {
	switch *g {
	case "^":
		return newGuard(">")
	case ">":
		return newGuard("v")
	case "v":
		return newGuard("<")
	default:
		return newGuard("^")
	}
}

func (g *Guard) getNextMove() Position {
	switch *g {
	case "^":
		return Position{-1, 0}
	case ">":
		return Position{0, 1}
	case "v":
		return Position{1, 0}
	default:
		return Position{0, -1}
	}
}

func (g *Guard) equals(s string) bool {
	return string(*g) == s
}

type Position struct {
	y, x int
}

func (p Position) string() string {
	return fmt.Sprintf("%d,%d", p.y, p.x)
}
