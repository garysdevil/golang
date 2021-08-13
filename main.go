package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	// "main/zero"
)

func main() {
	fmt.Println("--")
	// zero.ScanFunc1()
	fmt.Printf("%+v\n", 1)
	log.SetPrefix("main: ")
	// log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())
}
