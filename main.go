package main

import (
	"io/ioutil"
	"fmt"
)

const (
	_FILE_PATH = "input"
)

func main() {
	data, _ := ioutil.ReadFile(_FILE_PATH)
	nCount := [255]int{}
	for _, v := range data {
		nCount[v]++
	}
	for m, v := range nCount {
		if v != 0 {
			fmt.Println(m, string([]byte{byte(m)}), ":", v)
		}
	}
}