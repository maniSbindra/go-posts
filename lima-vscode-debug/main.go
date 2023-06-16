package main

import (
	"fmt"
	"syscall"
)

func main() {
	uts := &syscall.Utsname{}
	if err := syscall.Uname(uts); err != nil {
		fmt.Println("Error:", err)
		return
	}

	var releaseBytes []byte
	for _, b := range uts.Release {
		if b == 0 {
			break
		}
		releaseBytes = append(releaseBytes, byte(b))
	}

	kernelVersion := string(releaseBytes)
	fmt.Println("Kernel Version:", kernelVersion)
}

