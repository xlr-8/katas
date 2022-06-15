package wordchains

import (
	"errors"
)

// Solution regroups all the various component required to
// solve the 'word-chains' kata.
type solution struct {
	// start is the words will start building neighbors from
	start string
	// end is the target to arrive to from start
	end string
	// words contains a dictionnary with length of words as keys
	words map[int][]string
	// seen avoids looping several time on a item and saves the origin of an added word
	seen map[string]string
	// neighbors represents all neighbors that we have found from the
	// start word, or from its/their own neighbors.
	// A neighbor is a word with a difference of one from another word.
	neighbors stack
}

func (s solution) areValid(start string, end string, words []string) bool {
	if len(start) != len(end) || start == end {
		return false
	}

	found := 0
	for _, w := range words {
		if start == w || end == w {
			found++
		}
	}
	return found == 2
}

// haveSeen verifies that the word has not yet been seen
func (s solution) haveSeen(w string) bool {
	if _, ok := s.seen[w]; ok {
		return true
	}
	return false
}

// areSimilar will return 'true' if the 2 given words differ by
// only 1 caracter or less.
func (s solution) areSimilar(word string, compared string) bool {
	diff := 0

	if len(word) == len(compared) {
		for i := 0; i < len(word); i++ {
			if word[i] != compared[i] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
	}
	return true
}

// formatWordChain will build the final list of words backward between
// the 'start' and 'end' word until we've reached the first word.
func (s solution) formatWordChain(word string) []string {
	result := []string{word}
	word, ok := s.seen[word]
	for ok != false {
		if word != "" {
			result = append(result, word)
		}
		word, ok = s.seen[word]
	}
	return result
}

// findShortest search words in a similar way as a graph exploration.
//
// It starts by the 'start' word, and looks for potential
// neighbors in the list of words.
//
// All of those are added into a list of neighbors to explore, and
// are assigned into the 'seen' map in order to know where this
// neighbor originated from.
//
// Once the full list of neighbors or neighbors of neighbors etc
// has been exausted, we know that the word can't be found.
func (s *solution) findShortest(neighbor string, end string, words []string) []string {
	if neighbor == s.start && len(s.neighbors) == 0 {
		s.seen = map[string]string{s.start: ""}
		s.neighbors = []string{s.start}
	}

	if len(s.neighbors) == 0 {
		return nil
	}

	for _, word := range words {
		if s.areSimilar(word, neighbor) && !s.haveSeen(word) {
			s.seen[word] = neighbor
			if end == word {
				return s.formatWordChain(word)
			}
			s.neighbors.push(word)
		}
	}
	return s.findShortest(s.neighbors.pop(), end, words)
}

// Search ensures that words given are valid, that they exist
// in the loaded list of words, and try find the shortest path from
func (s *solution) Search(start string, end string) ([]string, error) {
	words, ok := s.words[len(start)]
	if ok == false {
		return nil, errors.New("start word has an invalid length")
	}

	if s.areValid(start, end, words) == false {
		return nil, errors.New("some words have not been found in the list")
	}

	s.start = start
	s.end = end
	return s.findShortest(s.start, s.end, words), nil
}

func New(path string) (*solution, error) {
	words, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	return &solution{
		words: words,
	}, nil
}
