package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var input string

var Reset = "\033[0m"
var Green = "\033[32m"
var Trail = Green + "o" + Reset

func main() {
	area := parseInput()
	start := findStart(area)
	grid := newGrid(deepCopy(area), start)
	s := time.Now()
	grid.analyzeRoute()
	fmt.Println("Unique cells", len(grid.visitedMap), "Initial Route", time.Since(s))

	sem := make(Semaphore, 1000)
	posCh := make(chan Pair)
	result := 0

	go func() {
		for range posCh {
			result++
		}
	}()

	var wg sync.WaitGroup
	for _, v := range grid.visitedMap {
		wg.Add(1)
		sem.Acquire()

		go func() {
			defer sem.Release()
			defer wg.Done()

			g := newGrid(deepCopy(area), start)
			g.setCell(v.pos, "#")

			if g.analyzeRoute() {
				posCh <- v.pos
			}
		}()
	}
	wg.Wait()

	fmt.Println("Result", result, "Full Analysis", time.Since(s))
}

func findStart(area [][]string) Pair {
	for i, row := range area {
		for j, cell := range row {
			if cell == "^" {
				return Pair{i, j}
			}
		}
	}

	return Pair{-1, -1}
}

func deepCopy(input [][]string) [][]string {
	copySlice := make([][]string, len(input))

	for i, innerSlice := range input {
		copySlice[i] = make([]string, len(innerSlice))
		copy(copySlice[i], innerSlice)
	}

	return copySlice
}

func parseInput() [][]string {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	area := make([][]string, len(rows))
	for i, row := range rows {
		area[i] = strings.Split(row, "")
	}
	return area
}

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

type Visit struct {
	pos   Pair
	count int
}

type Grid struct {
	area       [][]string
	visitedMap map[string]Visit
	guard      *Guard
	current    Pair
}

func newGrid(area [][]string, start Pair) *Grid {
	return &Grid{area, make(map[string]Visit), newGuard("^"), start}
}

func (g *Grid) analyzeRoute() bool {
	for {
		g.visiting(g.current)

		next := g.guard.nextMove()
		if g.leaving(g.current, next) {
			break
		}

		nextValue := g.nextCellValue(g.current, next)
		if nextValue == "#" {
			g.guard.turnRight(g)
		} else {
			g.guard.move(g, next)
		}

		if g.loopDetected(g.current) {
			return true
		}
	}

	return false
}

func (g *Grid) loopDetected(pos Pair) bool {
	if v, ok := g.visitedMap[pos.String()]; ok {
		return v.count >= 4
	}

	return false
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

func (g *Grid) leaving(position Pair, direction Pair) bool {
	rowOutBound := position.x+direction.x < 0 || position.x+direction.x >= len(g.area[0])
	colOutBound := position.y+direction.y < 0 || position.y+direction.y >= len(g.area)

	return rowOutBound || colOutBound
}

func (g *Grid) nextCellValue(position Pair, direction Pair) string {

	return g.area[position.y+direction.y][position.x+direction.x]
}

func (g *Grid) setCell(position Pair, value string) {
	g.area[position.y][position.x] = value
}

func (g *Grid) nextCell(position Pair, direction Pair) Pair {
	return Pair{position.y + direction.y, position.x + direction.x}
}

func (g *Grid) visiting(pos Pair) {
	if v, ok := g.visitedMap[pos.String()]; ok {
		v.count++
		g.visitedMap[pos.String()] = v
	} else {
		g.visitedMap[pos.String()] = Visit{pos, 1}
	}

}

type Guard string

func newGuard(s string) *Guard {
	return (*Guard)(&s)
}

func (g *Guard) String() string {
	return string(*g)
}

func (g *Guard) turnRight(grid *Grid) {
	guard := func() *Guard {
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
	}()
	grid.setCell(grid.current, guard.String())
	grid.guard = guard
}

func (g *Guard) move(grid *Grid, next Pair) {
	grid.setCell(grid.current, Trail)
	cell := grid.nextCell(grid.current, next)
	grid.setCell(cell, grid.guard.String())
	grid.current = cell
}

func (g *Guard) nextMove() Pair {
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

type Pair struct {
	y, x int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d,%d", p.y, p.x)
}
