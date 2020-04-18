package main

import (
	 "fmt"
	 "math/rand"
         "time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    var value = 0
    for true {
	value = random(0, 99999)
    	fmt.Println(value)
    }
}
