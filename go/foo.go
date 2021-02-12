package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    //TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		i++
	}
}

func decrementing() {
    //TODO: decrement i 1000000 times
	for k := 0; k < 1000000; k++ {
		i--
	}
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    
    // TODO: Spawn both functions as goroutines
	
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
