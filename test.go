package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	wth := sendHelpPls("c", 15, "Actually this is not bad")
	fmt.Println(wth)
}

/* Engine */
func compare2number(number1, number2 int) bool {
	if number1 > number2 {
		return true
	} else {
		return false
	}
}
func generate10k() int {
	var result int = 10000
	/* Just Loop to 10k */
	for i := 0; i < result; i++ {
		result = i + 1
	}
	return result
}
func playWithArray(name string, age int, address string) []string {
	return []string{
		name,
		strconv.Itoa(age),
		address,
	}
}

/* Custom data type, in this case OPPAI */
type oppai struct {
	OppaiTypes    string // A,B,C and etc
	OppaiDiameter int    // 46 centimeter
	Comment       string // example: "Please SEND ME HELP"
}

func sendHelpPls(types string, diameter int, comments string) string {
	result := &oppai{
		OppaiTypes:    types,
		OppaiDiameter: diameter,
		Comment:       comments,
	}
	endresult, _ := json.Marshal(result)
	return string(endresult)
}
