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
