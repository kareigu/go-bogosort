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
    var i = 0
    for i < 10 {
	var value = random(0, 99999)
    	fmt.Println(value)
	i++
    }
}
