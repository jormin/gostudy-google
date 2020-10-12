package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{
		p.i + r.i,
		p.j + r.j,
	}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func walk(maze [][]int, start point, end point) (success bool, steps [][]int, maxSteps int, path []point) {
	steps = make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	q := []point{start}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)
			//  maze at next is 0
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// steps at next is 0
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// next != start
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			// find the end
			if next == end {
				q = []point{}
				break
			}
			q = append(q, next)
		}
	}
	maxSteps, _ = end.at(steps)
	if maxSteps == 0 {
		return false, [][]int{}, 0, nil
	}
	path = append(path, end)
	tmpStep := maxSteps - 1
	cur := end
	for tmpStep >= 0 {
		for _, dir := range dirs {
			tmpCur := cur.add(dir)
			step, _ := tmpCur.at(steps)
			if step == tmpStep {
				cur = tmpCur
				path = append([]point{cur}, path...)
				tmpStep--
			}
		}
	}
	return true, steps, maxSteps, path
}

func main() {
	// 读取迷宫数据
	maze := readMaze("maze/maze.in")
	//for _, row := range maze {
	//	for _, col := range row {
	//		fmt.Printf("%d ", col)
	//	}
	//	fmt.Println()
	//}

	// 走迷宫
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	success, steps, maxSteps, path := walk(maze, start, end)
	if !success {
		fmt.Printf("There is no path from %v to %v\n", start, end)
		return
	}
	fmt.Printf("Find the way from %v to %v\n", start, end)
	fmt.Println("The steps is")
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
	fmt.Println("The max steps is", maxSteps)
	fmt.Println("The path is", path)
}
