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
	// DontKnow enum for possible match results.
	DontKnow CompareResult = iota
	// NoMatch enum for possible match results.
	NoMatch
	// InWord enum for possible match results.
	InWord
	// AtPlace enum for possible match results.
	AtPlace
)

// CharacterCompare captures a comparison between characters.
type CharacterCompare struct {
	Index      int           `json:"index"`
	SourceChar byte          `json:"sourceChar"`
	Result     CompareResult `json:"result"`
}

// CheckResult is the result of a comparison between a word and the target word.
type CheckResult struct {
	Word       string              `json:"word"`
	Comparison []*CharacterCompare `json:"comparison"`
}

// Word is word being used as target for the game.
type Word struct {
	word string
}

// IsSolved checks if the word was matched.
func (cr *CheckResult) IsSolved() bool {
	countAtPlace := 0
	for _, c := range cr.Comparison {
		if c.Result == AtPlace {
			countAtPlace++
		}
	}

	return countAtPlace == len(cr.Word)
}

// NewWord creates checks the validity of a word and then creates it.
func NewWord(word string) (*Word, error) {
	if !IsValidWord(word) {
		return nil, errors.Errorf("'%s' is not a valid word", word)
	}

	return &Word{
		word: normalizeWord(word),
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

	word = normalizeWord(word)

	res := &CheckResult{
		Word:       word,
		Comparison: make([]*CharacterCompare, len(word)),
	}
	compWord := w.word
	// find exact matches, setting exact matches as empty to not double count
	for i := range word {
		if word[i] == compWord[i] {
			res.Comparison[i] = &CharacterCompare{
				Index:      i,
				SourceChar: word[i],
				Result:     AtPlace,
			}
			compWord = updateChar(compWord, " ", i)
			//fmt.Printf("NEW COMP AFTER MATCH AT %d: %v\n", i, compWord)
		}
	}

	// find letters in word but not at the right place
	for i := range word {
		if res.Comparison[i] == nil {
			for j := range compWord {
				if compWord[j] == word[i] {
					res.Comparison[i] = &CharacterCompare{
						Index:      i,
						SourceChar: word[i],
						Result:     InWord,
					}
					compWord = updateChar(compWord, " ", j)
					//fmt.Printf("NEW COMP AFTER FIND AT %d: %v\n", j, compWord)
					break
				}
			}
		}
	}

	// mark remaining as not matched
	for i := range res.Comparison {
		if res.Comparison[i] == nil {
			res.Comparison[i] = &CharacterCompare{
				Index:      i,
				SourceChar: word[i],
				Result:     NoMatch,
			}
		}
	}

	return res, nil
}

// IsValidWord checks to make sure a word meets the constraints.
func IsValidWord(word string) bool {
	word = normalizeWord(word)
	if !letters.MatchString(word) {
		return false
	}

	return true
}

func updateChar(current string, newValue string, index int) string {
	if index == 0 {
		return newValue + current[1:]
	} else if index == len(current)-1 {
		return current[:len(current)-1] + newValue
	}

	return current[:index] + newValue + current[index+1:]
}

func normalizeWord(word string) string {
	// TODO: remove accents

	// case insensitve
	word = strings.ToUpper(word)

	return word
}
