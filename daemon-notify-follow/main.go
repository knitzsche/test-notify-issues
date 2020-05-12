package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	snap := os.Getenv("SNAP_NAME")
	loop := true
	idx := -1
	for loop {
		idx += 1
		fmt.Printf("==== %s.notify-follow running. idx: %v\n", snap, idx)
		time.Sleep(1 * time.Second)
		if idx == 5 {
			break
		}
	}
}
