package utils

import (
	"io/ioutil"
)

func ReadInput() string {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return string(data)
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
