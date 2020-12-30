package main

import (
	"fmt"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type vector3 struct {
	x int
	y int
	z int
}

func parseGrid(input string) []vector3 {
	input = strings.TrimSpace(input)
	grid := []vector3{}
	for y, line := range strings.Split(input, "\n") {
		middleY := y - len(line)/2
		for x, char := range strings.Split(line, "") {
			middleX := x - len(line)/2
			if char == "#" {
				grid = append(grid, vector3{x: middleX, y: middleY, z: 0})
			}
		}
	}
	return grid
}

func problemSpace(max int) []vector3 {
	grid := []vector3{}
	for x := -max; x <= max; x++ {
		for y := -max; y <= max; y++ {
			for z := -max; z <= max; z++ {
				grid = append(grid, vector3{x: x, y: y, z: z})
			}
		}
	}
	return grid
}

// Each cube only ever considers its neighbors:
// any of the 26 other cubes where any of their coordinates differ by at most 1.
// For example, given the cube at x=1,y=2,z=3, its neighbors include the cube at x=2,y=2,z=2
// the cube at x=0,y=2,z=3, and so on.
func (cube vector3) isNeighborTo(q vector3) bool {
	return (q.x-1 <= cube.x && cube.x <= q.x+1) &&
		(q.y-1 <= cube.y && cube.y <= q.y+1) &&
		(q.z-1 <= cube.z && cube.z <= q.z+1) &&
		(cube != q)
}

func countNeighbors(cube vector3, grid []vector3) int {
	counter := 0
	for _, c := range grid {
		if c.isNeighborTo(cube) {
			counter++
		}
	}
	return counter
}

func isActive(cube vector3, grid []vector3) bool {
	for _, c := range grid {
		if cube == c {
			return true
		}
	}
	return false
}

func maxInt(numbers ...int) int {
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMaxDistance(grid []vector3) int {
	max := 0
	for _, c := range grid {
		max = maxInt(max, absInt(c.x), absInt(c.y), absInt(c.z))
	}
	return max
}

// cycle performs the following rules on a grid:
//
// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
func cycle(grid []vector3, cycle int) []vector3 {
	newGrid := []vector3{}
	max := getMaxDistance(grid)
	for _, cube := range problemSpace(max + 1) {
		numberOfNeighbors := countNeighbors(cube, grid)
		if isActive(cube, grid) && (numberOfNeighbors == 2 || numberOfNeighbors == 3) {
			newGrid = append(newGrid, cube)
		} else if !isActive(cube, grid) && numberOfNeighbors == 3 {
			newGrid = append(newGrid, cube)
		}
	}
	return newGrid
}

func printGrid(grid []vector3, cycle int) {
	fmt.Printf("After %d cycles:\n", cycle)
	fmt.Println(renderGrid(grid, cycle))
	fmt.Println()
}

func renderGrid(grid []vector3, cycle int) string {
	output := ""
	max := getMaxDistance(grid)
	for z := -max; z <= max; z++ {
		output += renderSlice(grid, max, z) + "\n"
	}
	return output
}

func renderSlice(grid []vector3, max int, z int) string {
	output := fmt.Sprintf("z=%d\n", z)
	for y := -max; y <= max; y++ {
		for x := -max; x <= max; x++ {
			if isActive(vector3{x: x, y: y, z: z}, grid) {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func partOne(grid []vector3) int {
	for i := 1; i <= 6; i++ {
		printGrid(grid, i)
		grid = cycle(grid, i)
	}
	return len(grid)
}

func main() {
	input := utils.ReadInput()
	initialGrid := parseGrid(input)

	result := partOne(initialGrid)
	fmt.Printf("Answer to part one: %d\n", result)
}
