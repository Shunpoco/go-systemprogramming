package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Fprintf formats data, then writes to io.Writer.
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
