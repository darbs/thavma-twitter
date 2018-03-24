package main

import (
	"fmt"

	"github.com/darbs/thavma-twitter/internal/fetch"
)

func main() {
	fmt.Printf("Main start\n")

	fetch.Get("$atrs")
}
