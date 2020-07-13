package main

import (
	"flag"
	"fmt"
	"os"
)

const defaultCost = 10

func main() {
	password := flag.String("password", "", "Password to hash")
	cost := flag.Int("cost", defaultCost, "Cost between 4 and 31")

	flag.Usage = func() {
		fmt.Printf("\nUsage: go-hash [OPTIONS]\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n")
		fmt.Printf("  go-hash -password mypassword\n")
		fmt.Printf("  go-hash -password \"hello world\" -cost 20\n\n")
	}

	flag.Parse()

	// make sure a password was provided
	if *password == "" {
		fmt.Printf("error: a password must be provided with the -password flag\n")
		os.Exit(2)
	}

	// validate -cost flag meets allowed range
	if *cost < 4 || *cost > 31 {
		fmt.Printf("error: cost %v is outside allowed range (4,31)\n", *cost)
		os.Exit(2)
	}

	// hash the password
	hash, err := hash(*password, *cost)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(2)
	}

	// double check generated hash can be compared with original password
	if !verify(*password, hash) {
		fmt.Printf("error: password failed verification\n")
		os.Exit(2)
	}

	// display the hashed password
	fmt.Printf("hash: %v\n", hash)

	// output cost used when -cost flag not specified
	if *cost == defaultCost {
		fmt.Printf("cost: %v\n", *cost)
	}
}
