package main

import (
	"log"
	"fmt"
)

func throwError(err string, lineno int, col int) {
	log.Fatal(fmt.Sprintf("%s, line: %d, col: %d", err, lineno, col))
}