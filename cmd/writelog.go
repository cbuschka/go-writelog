package main

import (
	"fmt"
	"github.com/cbuschka/go-writelog/internal"
	"os"
)

func main() {
	err := internal.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed: %s", err.Error())
		os.Exit(1)
	}
}
