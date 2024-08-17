package main

import (
	"fmt"
	"time"
)

func main() {
	longRunningFunc()

	x := 1
	y := true
	Learn(x)
	Learn(y)
}

func longRunningFunc() {
	defer Trace("longRunningFunc()")()
	time.Sleep(time.Duration(1) * time.Second)
}

func Trace(name string) func() {
	var t = time.Now()
	fmt.Printf("Entering %s at %s\n", name, t)
	return func() {
		fmt.Printf("Exiting %s taken %s \n", name, time.Now().Sub(t))
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf(" Data type: %T\n", T)
	}
}
