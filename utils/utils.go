package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput() string {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ConvertIntoInts(input string) []int {
	ints := []int{}
	input = strings.TrimSpace(input)
	for _, line := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func RemoveDuplicatesStrings(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}
	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}
