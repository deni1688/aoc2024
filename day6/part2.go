package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
	grid := newGrid(input, newGuard("^"))
	// time it
	s := time.Now()
	res := grid.analyzeRoute(true)
	fmt.Println("Time", time.Since(s))

	sem := make(Semaphore, 7000)
	posCh := make(chan Pair)
	m := map[string]bool{}

	go func() {
		for v := range posCh {
			m[v.String()] = true
		}
	}()

	var wg sync.WaitGroup

	for _, v := range res.route {
		v := v
		wg.Add(1)
		sem.Acquire()

		go func() {
			defer sem.Release()
			defer wg.Done()

			g := newGrid(input, newGuard("^"))
			g.setCell(v, "#")
			r := g.analyzeRoute(false)
			if r.loop {
				posCh <- v
			}
		}()
	}
	wg.Wait()

	fmt.Println("Done", len(m))
}

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

type Grid struct {
	area       [][]string
	visitedMap map[string]int
	guard      *Guard
}

func newGrid(input string, guard *Guard) *Grid {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	area := make([][]string, len(rows))
	for i, row := range rows {
		area[i] = strings.Split(row, "")
	}

	return &Grid{area, make(map[string]int), guard}
}

type AnalyzeResult struct {
	route []Pair
	loop  bool
}

func (g *Grid) analyzeRoute(calcRoute bool) AnalyzeResult {
	for {
		currentPosition := g.findGuard()
		right := g.guard.getRightTurn()
		move := g.guard.getNextMove()

		g.setCellVisited(currentPosition, g.guard)

		if g.willLoop() {
			return AnalyzeResult{
				route: nil,
				loop:  true,
			}
		}

		if g.isLeaving(currentPosition, move) {
			break
		}

		next := g.getCellValue(currentPosition, move)
		if next == "#" {
			g.setCell(currentPosition, right.String())
			g.guard = right
		} else {
			g.setCell(currentPosition, Trail)
			g.setCell(g.nextCell(currentPosition, move), g.guard.String())
		}
	}

	if !calcRoute {
		return AnalyzeResult{
			route: nil,
			loop:  false,
		}
	}

	result := make([]Pair, 0)

	for k := range g.visitedMap {
		pair := strings.Split(k, ",")
		y, _ := strconv.Atoi(pair[0])
		x, _ := strconv.Atoi(pair[1])
		result = append(result, Pair{y, x})
	}

	return AnalyzeResult{
		route: result,
		loop:  false,
	}
}

func (g *Grid) willLoop() bool {
	for _, visited := range g.visitedMap {
		if visited > 1 {
			return true
		}
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
