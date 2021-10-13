package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/solarkennedy/AdventOfCode/utils"
)

type instruction struct {
	mask         string
	location     int
	preMaskValue int
}

func parseMemLocation(input string) int {
	first := strings.Split(input, "[")[1]
	second := strings.Split(first, "]")[0]
	return utils.Atoi(second)

}

func parseInstructions(input string) []instruction {
	ins := []instruction{}
	var mask string
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		tokens := strings.Split(s, " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
		} else {
			ins = append(ins, instruction{
				mask:         mask,
				location:     parseMemLocation(tokens[0]),
				preMaskValue: utils.Atoi(tokens[1]),
			})
		}
	}
	return ins
}

func bitSet(value int, place int) int {
	realMask := 1 << place
	return value | realMask
}

func bitClear(value int, place int) int {
	realMask := 1 << place
	return value &^ realMask
}

func computePostMaskValue(preValue int, mask string) int {
	value := preValue
	fmt.Printf("Value: %036b (%d)\n", value, value)
	fmt.Printf("Mask : %s\n", mask)
	for c := range mask {
		if mask[c] == 'X' {
			continue
		} else if mask[c] == '1' {
			value = bitSet(value, 35-c)
		} else if mask[c] == '0' {
			value = bitClear(value, 35-c)
		} else {
			panic(mask[c])
		}
	}
	fmt.Printf("Res  : %036b (%d)\n", value, value)
	return value
}

func computePostMaskValues(ins []instruction) int {
	m := map[int]int{}
	for i := range ins {
		fmt.Printf("Ins %d (%d):\n", i, ins[i].location)
		m[ins[i].location] = computePostMaskValue(ins[i].preMaskValue, ins[i].mask)
	}
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func partOne(input string) int {
	ins := parseInstructions(input)
	return computePostMaskValues(ins)
}

func partTwo(input string) int {
	ins := parseInstructions(input)
	addrs := parseInstructionAddressDecoders(ins)
	return sumOfAddresses(addrs)
}

func parseInstructionAddressDecoders(ins []instruction) map[string]int {
	addrs := map[string]int{}
	for _, ins := range ins {
		maskedString := parseSingleInstructionAddressDecode(ins)
		comboMap := replaceXsWithAllCombos(maskedString, ins.preMaskValue)
		for k, v := range comboMap {
			addrs[k] = v
		}
	}
	return addrs
}

func parseSingleInstructionAddressDecode(ins instruction) string {
	applied := ""
	vString := fmt.Sprintf("%036b", ins.location)
	// First apply the mask
	for i := 0; i < 36; i++ {
		m := ins.mask[i]
		if m == '1' {
			applied = applied + string(m)
		} else if m == '0' {
			applied = applied + string(vString[i])
		} else if m == 'X' {
			applied = applied + "X"
		} else {
			panic(m)
		}
	}
	fmt.Printf("Took %+v and turned it into `%s`\n", ins, applied)
	return applied
}

func replaceXsWithAllCombos(input string, value int) map[string]int {
	// Now we still have X's in there potentially
	ret := map[string]int{}
	Xs := strings.Count(input, "X")
	for i := 0; i < powInt(2, Xs); i++ {
		// Come up with every combo and set the mask
		comboString := fmt.Sprintf("%0*b", Xs, i)
		comboMaskedValue := input
		for _, b := range comboString {
			comboMaskedValue = strings.Replace(comboMaskedValue, "X", string(b), 1)
		}
		ret[comboMaskedValue] = value
	}
	fmt.Printf("Took `%s` and exploded it to all %d possibilities\n", input, len(ret))
	return ret
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func sumOfAddresses(addrs map[string]int) int {
	total := 0
	for _, v := range addrs {
		total += v
	}
	return total
}

func main() {
	input := utils.ReadInput()
	result := partOne(input)
	fmt.Printf("Answer to part one: %d\n", result)

	result2 := partTwo(input)
	fmt.Printf("Answer to part two: %d\n", result2)
}
