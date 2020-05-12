package main

import (
	"fmt"
	"os"
	"time"

	"github.com/okzk/sdnotify"
)

func main() {
	snap := os.Getenv("SNAP_NAME")
	loop := true
	idx := -1
	for loop {
		idx += 1
		fmt.Printf("==== %s.watchdog-notify running. idx: %v\n", snap, idx)
		if idx == 5 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	sdnotify.SdNotify("READY=1")
	time.Sleep(2 * time.Second)

	fmt.Printf("==== %s.watchdog-notify ending.\n", snap)

}
