package main

import (
	"os"
	"time"

	svg "github.com/eckertalex/gobytest/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
