package board

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	letters = regexp.MustCompile(`[A-Z]+$`)
)

// CompareResult is the result of a comparison between characters.
type CompareResult int

const (
	// NoMatch enum for possible match results.
	NoMatch CompareResult = iota
	// InWord enum for possible match results.
	InWord
	// AtPlace enum for possible match results.
	AtPlace
)

// CharacterCompare captures a comparison between characters.
type CharacterCompare struct {
	index      int
	sourceChar byte
	result     CompareResult
}

// CheckResult is the result of a comparison between a word and the target word.
type CheckResult struct {
	word       string
	comparison []*CharacterCompare
}

// Word is word being used as target for the game.
type Word struct {
	word string
}

// NewWord creates checks the validity of a word and then creates it.
func NewWord(word string) (*Word, error) {
	if !IsValidWord(word) {
		return nil, errors.Errorf("'%s' is not a valid word", word)
	}

	return &Word{
		word: strings.ToUpper(word),
	}, nil
}

// Check compares the submitted word to the target word, responding on a
// character by character basis identifying the characters that are at the
// right place, in the word at the wrong place, or not in the word.
func (w *Word) Check(word string) (*CheckResult, error) {
	if !IsValidWord(word) {
		return nil, errors.Errorf("'%s' is not a valid word", word)
	}

	if len(word) != len(w.word) {
		return nil, errors.Errorf("word needs to be %d characters but '%s' has %d", len(w.word), word, len(word))
	}

	res := &CheckResult{
		word:       word,
		comparison: make([]*CharacterCompare, len(word)),
	}
	compWord := w.word
	// find exact matches, setting exact matches as empty to not double count
	for i := range word {
		if w.word[i] == compWord[i] {
			res.comparison[i] = &CharacterCompare{
				index:      i,
				sourceChar: word[i],
				result:     AtPlace,
			}
			compWord = compWord[:i-1] + " " + compWord[i:]
		}
	}

	// find letters in word but not at the right place
	for i := range word {
		if res.comparison[i] == nil {
			for j := range compWord {
				if compWord[j] == word[i] {
					res.comparison[i] = &CharacterCompare{
						index:      i,
						sourceChar: word[i],
						result:     InWord,
					}
					compWord = compWord[:i-1] + " " + compWord[i:]
					break
				}
			}
		}
	}

	// mark remaining as not matched
	for i := range res.comparison {
		if res.comparison[i] == nil {
			res.comparison[i] = &CharacterCompare{
				index:      i,
				sourceChar: word[i],
				result:     NoMatch,
			}
		}
	}

	return res, nil
}

// IsValidWord checks to make sure a word meets the constraints.
func IsValidWord(word string) bool {
	word = strings.ToUpper(word)
	if !letters.MatchString(word) {
		return false
	}

	return true
}
