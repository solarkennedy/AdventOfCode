package main

var adapterPathsAlreadyTried = make(map[int]int)

func howManyAdaptersWorkStartingAt(adapters []int, startAdapter int) int {
	workingPaths := 0
	if startAdapter+3 >= len(adapters) {
		// We have overshot at this point
		return 1
	} else if _, ok := adapterPathsAlreadyTried[adapters[startAdapter]]; ok {
		// Fast path, if we have been here before, we already know how many paths
		return adapterPathsAlreadyTried[adapters[startAdapter]]
	}

	for i := startAdapter + 1; i <= startAdapter+3; i++ {
		lower := adapters[startAdapter]
		upper := adapters[i]
		if (upper-lower) >= 1 && (upper-lower) <= 3 {
			workingPaths += howManyAdaptersWorkStartingAt(adapters, i)
		}
	}

	// Cache what we have seen already into the global map
	adapterPathsAlreadyTried[adapters[startAdapter]] = workingPaths
	return workingPaths
}

func partTwo(input string) int {
	adapterPathsAlreadyTried = make(map[int]int)
	adapters := parseAdapters(input)
	builtIn := adapters[len(adapters)-1] + 3
	wall := []int{0}
	adapters = append(wall, adapters...)
	adapters = append(adapters, builtIn)
	return howManyAdaptersWorkStartingAt(adapters, 0)
}
