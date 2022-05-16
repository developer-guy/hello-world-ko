package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("I'm running on OS %s and arch %s", runtime.GOOS, runtime.GOARCH)
}
