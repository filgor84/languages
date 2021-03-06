package main

import (
	"fmt"
	"io/ioutil"

	"github.com/filgor84/languages/arithspecial"
)

func main() {
	dirName := "/dev/shm/"
	fileName := "data/20MB.txt"
	data, err := ioutil.ReadFile(dirName + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := arithspecial.parseWhole(data, 1)
	fmt.Println(res)

}
