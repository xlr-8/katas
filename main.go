package main

import (
	"fmt"
	"os"

	wordchains "github.com/xlr-8/katas/wordchains"
)

func main() {
	const wordList = "./wordlists/wordlist.txt"

	if len(os.Args) != 3 {
		fmt.Println("not enough arguments, or arguments invalid (length/diff)")
		os.Exit(1)
	}

	s, err := wordchains.New(wordList)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	chain, err := s.Search(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if chain != nil {
		fmt.Println("Found:", chain)
	} else {
		fmt.Println("Couldn't find anything matching")
		os.Exit(1)
	}
}
