package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type Adapter struct {
	//	index  int
	rating int
}

type AdapterChain struct {
	adapters []Adapter
	unused   []Adapter
}

func parseAdapters(input string) []Adapter {
	ints := []int{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		ints = append(ints, utils.Atoi(s))
	}
	adapters := []Adapter{}
	sort.Ints(ints)
	for _, i := range ints {
		adapters = append(adapters, Adapter{rating: i})
	}
	return adapters
}

func (a Adapter) IsCompatibleWithInputVoltage(voltage int) bool {
	// adapters can only connect to a source 1-3 jolts lower than its rating,
	return a.rating-voltage >= 1 && a.rating-voltage <= 3
}

func findSuitableAdapters(adapters []Adapter, voltage int) []int {
	ret := []int{}
	for i, a := range adapters {
		if a.IsCompatibleWithInputVoltage(voltage) {
			ret = append(ret, i)
			// Hack for speed, always the first one we find???
			return ret
		}
	}
	return ret
}

func getLastVoltage(adapters []Adapter) int {
	if len(adapters) == 0 {
		// The charging outlet has an effective rating of 0 jolts
		return 0
	}
	return adapters[len(adapters)-1].rating
}

func findWorkingChains(c AdapterChain) []AdapterChain {
	if len(c.unused) == 0 {
		return []AdapterChain{c}
	}
	fmt.Printf("%sDepth %d\n", strings.Repeat(" ", len(c.adapters)), len(c.adapters))
	workingChains := []AdapterChain{}
	voltage := getLastVoltage(c.adapters)
	suitableAdapters := findSuitableAdapters(c.unused, voltage)
	for _, i := range suitableAdapters {
		suitableAdapter := c.unused[i]
		fmt.Printf("Going to try adding %+v to a chain of %+v\n", suitableAdapter, c)
		newUnused := removeAdapter(c.unused, i)
		newAdapters := append(c.adapters, suitableAdapter)
		newChain := AdapterChain{
			adapters: newAdapters,
			unused:   newUnused,
		}
		// fmt.Printf("   New chain: %+v\n", newChain)
		newChains := findWorkingChains(newChain)
		workingChains = append(workingChains, newChains...)
	}
	// if len(workingChains) == 0 {
	// 	fmt.Printf("%sNo valid chains found after %+v\n", strings.Repeat(" ", len(c.adapters)), c)
	// }
	return workingChains
}

func removeAdapter(s []Adapter, index int) []Adapter {
	ret := make([]Adapter, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func findWorkingChain(adapters []Adapter) AdapterChain {
	fmt.Printf("Building chains using %d different adapters...\n%+v\n", len(adapters), adapters)
	chains := findWorkingChains(AdapterChain{unused: adapters})
	if len(chains) != 1 {
		// Shouldn't happen
		return AdapterChain{}
	}
	fmt.Printf("FOUND IT: %+v\n", chains[0])
	return chains[0]
}

func countJoltDifferences(adapters []Adapter, diff int) int {
	total := 0
	for i := range adapters {
		if i == len(adapters)-1 {
			continue
		}
		differential := adapters[i+1].rating - adapters[i].rating
		if differential == diff {
			total++
		}
	}
	// Last built-in adapter is always 3 higher than the first
	if diff == 3 {
		total++
	}
	// Wall has 0
	wallDiff := adapters[0].rating - 0
	if wallDiff == diff {
		total++
	}
	return total
}

func partOne(input string) int {
	adapters := parseAdapters(input)
	chain := findWorkingChain(adapters)
	ones := countJoltDifferences(chain.adapters, 1)
	threes := countJoltDifferences(chain.adapters, 3)
	return ones * threes
}

func partTwo(input string) int {
	return 42
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
