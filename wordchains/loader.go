package wordchains

import (
	"bufio"
	"os"
)

// loadFile takes the path of a list of word and will
// load all of them into memory. To save processing time
// on multiple call, the list is loaded into a map[int][]string
// which allows to easily access the list of word (value)
// of a given length (key).
func loadFile(path string) (map[int][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	words := make(map[int][]string)
	for scanner.Scan() {
		word := scanner.Text()
		words[len(word)] = append(words[len(word)], word)
	}
	return words, scanner.Err()
}
