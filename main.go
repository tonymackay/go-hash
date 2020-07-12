package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	password := flag.String("password", "", "Password to hash")
	cost := flag.Int("cost", 10, "Cost between 4 and 31")

	flag.Usage = func() {
		fmt.Printf("\nUsage: go-hash [OPTIONS]\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n")
		fmt.Printf("  go-hash -password mypassword\n")
		fmt.Printf("  go-hash -password \"hello world\" -cost 20\n\n")
	}

	flag.Parse()

	if *password == "" {
		fmt.Printf("error: %v\n", "a password must be provided with the -password flag")
		os.Exit(2)
	}

	fmt.Printf("hashing: %v\n", *password)

	hash, err := hash(*password, *cost)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(2)
	}
	fmt.Println(hash)
}
