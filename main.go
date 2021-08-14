package main

import (
	"fmt"
	"regexp"
)

func main() {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	// want := regexp.MustCompile(name)

	var err []string
	msg := "Glads" // errors.New("hhhhh")
	if !want.MatchString(msg) || err != nil {
		fmt.Printf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
	
}
