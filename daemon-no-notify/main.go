package main

import (
	"fmt"
	"os"
	"time"

	"github.com/okzk/sdnotify"
)

func main() {
	snap := os.Getenv("SNAP_NAME")
	sdnotify.SdNotify("READY=1")
	loop := true
	idx := -1
	for loop {
		idx += 1
		fmt.Printf("==== %s.watchdog-no-notify running. idx: %v\n", snap, idx)
		if idx == 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
