package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

var Reset = "\033[0m"
var Green = "\033[32m"

func main() {
	grid := newGrid(input, newGuard("^"))
	grid.analyzeRoute()

	sem := make(chan struct{}, 50)
	var wg sync.WaitGroup
	var m sync.Map
	for i, v := range grid.route {
		sem <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() {
				<-sem
			}()

			defer wg.Done()

			if work(v) {
				m.Store(v.String(), true)
			}
		}()

		fmt.Printf("Working on %d/%d\n", i, len(grid.route))
	}
	wg.Wait()

	size := 0
	m.Range(func(k, v any) bool {
		size++
		return true
	})
	fmt.Println("Done", size)
}

func work(v Pair) bool {
	g := newGrid(input, newGuard("^"))
	g.setCell(v, "#")
	g.analyzeRoute()

	return g.willLoop()
}

func (g *Grid) analyzeRoute() int {
	for {
		if g.willLoop() {
			break
		}
		currentPosition := g.findGuard()
		right := g.guard.getRightTurn()
		move := g.guard.getNextMove()

		g.setCellVisited(currentPosition, g.guard)

		if g.isLeaving(currentPosition, move) {
			break
		}

		next := g.getCellValue(currentPosition, move)
		if next == "#" {
			g.setCell(currentPosition, right.String())
			g.guard = right
		} else {
			g.setCell(currentPosition, fmt.Sprintf("%s%s%s", Green, "+", Reset))
			g.setCell(g.nextCell(currentPosition, move), g.guard.String())
		}
	}

	result := make(map[string]int)
	for k, v := range g.visitedMap {
		vals := strings.Split(k, ",")
		result[vals[0]+","+vals[1]] = v
	}

	return len(result)
}

func (g *Grid) willLoop() bool {
	for _, v := range g.visitedMap {
		if v > 1 {
			return true
		}
	}

	return false
}

type Grid struct {
	area       [][]string
	visitedMap map[string]int
	route      []Pair
	guard      *Guard
}

func newGrid(input string, guard *Guard) *Grid {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	area := make([][]string, len(rows))
	for i, row := range rows {
		area[i] = strings.Split(row, "")
	}

	return &Grid{area, make(map[string]int), make([]Pair, 0), guard}
}

func (g *Grid) print() {
	fmt.Printf("\033[0;0H")
	fmt.Println(g.String())
}

func (g *Grid) String() string {
	var sb strings.Builder
	for _, row := range g.area {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g *Grid) findGuard() Pair {
	for i, row := range g.area {
		for j, cell := range row {
			if g.guard.equals(cell) {
				return Pair{i, j}
			}
		}
	}

	return Pair{-1, -1}
}

func (g *Grid) isLeaving(position Pair, direction Pair) bool {
	rowOutBound := position.x+direction.x < 0 || position.x+direction.x >= len(g.area[0])
	colOutBound := position.y+direction.y < 0 || position.y+direction.y >= len(g.area)

	return rowOutBound || colOutBound
}

func (g *Grid) getCellValue(position Pair, direction Pair) string {

	return g.area[position.y+direction.y][position.x+direction.x]
}

func (g *Grid) setCell(position Pair, value string) {
	g.area[position.y][position.x] = value
}

func (g *Grid) nextCell(position Pair, direction Pair) Pair {
	return Pair{position.y + direction.y, position.x + direction.x}
}

func (g *Grid) setCellVisited(position Pair, guard *Guard) {
	g.route = append(g.route, position)
	key := position.String() + "," + guard.String()
	g.visitedMap[key] = g.visitedMap[key] + 1
}

type Guard string

func newGuard(s string) *Guard {
	return (*Guard)(&s)
}

func (g *Guard) String() string {
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

func (g *Guard) getNextMove() Pair {
	switch *g {
	case "^":
		return Pair{-1, 0}
	case ">":
		return Pair{0, 1}
	case "v":
		return Pair{1, 0}
	case "<":
		return Pair{0, -1}
	default:
		panic("Invalid guard")
	}
}

func (g *Guard) equals(s string) bool {
	return string(*g) == s
}

type Pair struct {
	y, x int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d,%d", p.y, p.x)
}
