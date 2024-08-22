package main

import (
	"fmt"
	"time"
)

func main() {
	testLearn()
	testSpinner()
	testTrace()
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf(" Data type: %T\n", T)
	}
}

func testLearn() {
	x := 1
	y := true
	Learn(x)
	Learn(y)
}

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func testSpinner() {
	go Spinner(100 * time.Millisecond)
}

func Trace(name string) func() {
	var t = time.Now()
	fmt.Printf("Entering %s at %s\n", name, t)
	return func() {
		fmt.Printf("Exiting %s taken %s \n", name, time.Now().Sub(t))
	}
}

func testTrace() {
	longRunningFunc := func() {
		defer Trace("longRunningFunc()")()
		time.Sleep(time.Duration(1) * time.Second)
	}
	longRunningFunc()
}
