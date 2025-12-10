package day1 

import (
	"fmt"
	"github.com/nishan-soni/advent2025/utils"
)


func Solve() {
	lines, err := utils.GetFileLines("input.txt")

	if err != nil {
		panic(err)
	}
	
	for line := range lines {
		fmt.Println(line)
	}

}

