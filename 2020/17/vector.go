package main

type vector3 struct {
	x int
	y int
	z int
}

type vector4 struct {
	x int
	y int
	z int
	w int
}

func place3to4(cubes []vector3) []vector4 {
	hyperplane := 0
	hypercubes := []vector4{}
	for _, cube := range cubes {
		hypercubes = append(hypercubes, vector4{x: cube.x, y: cube.y, z: cube.z, w: hyperplane})
	}
	return hypercubes
}
