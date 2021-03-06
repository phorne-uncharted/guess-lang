package board

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	log "github.com/unchartedsoftware/plog"
)

// GameConfig defines basic configurations used for the game.
type GameConfig struct {
	maxGuessCount int
	letterCount   int
	sourceFile    string
	acceptedFile  string
	language      string
}

// Game is the main driver.
type Game struct {
	target       *Word
	config       *GameConfig
	allowedWords map[string]bool
	guessCount   int
	knowledge    *TargetKnowledge
}

// TargetKnowledge summarizes what is known about the target.
type TargetKnowledge struct {
	Results []*CheckResult        `json:"results"`
	Letters map[int]CompareResult `json:"letters"`
}

// NewTargetKnowledge creates a black target knowledge.
func NewTargetKnowledge(guessCount int) *TargetKnowledge {
	tk := &TargetKnowledge{
		Results: make([]*CheckResult, guessCount),
		Letters: map[int]CompareResult{},
	}

	for b := 65; b < 91; b++ {
		tk.Letters[b] = DontKnow
	}

	return tk
}

// AddCheckResult adds a result to the target target knowlegde.
func (t *TargetKnowledge) AddCheckResult(cr *CheckResult) {
	guessCount := 0
	for guessCount < len(t.Results) {
		if t.Results[guessCount] == nil {
			break
		}
		guessCount = guessCount + 1
	}

	t.Results[guessCount] = cr
	for _, r := range cr.Comparison {
		if r.Result == AtPlace {
			t.Letters[int(r.SourceChar)] = AtPlace
		} else if r.Result == InWord && t.Letters[int(r.SourceChar)] == DontKnow {
			t.Letters[int(r.SourceChar)] = InWord
		} else if t.Letters[int(r.SourceChar)] == DontKnow {
			t.Letters[int(r.SourceChar)] = NoMatch
		}
	}
}

// NewGame creates a new game, initializing with a target word and allowed words.
func NewGame(maxGuessCount int, letterCount int, sourceFile string, acceptedFile string) (*Game, error) {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		config: &GameConfig{
			maxGuessCount: maxGuessCount,
			letterCount:   letterCount,
			sourceFile:    sourceFile,
			acceptedFile:  acceptedFile,
		},
		knowledge: NewTargetKnowledge(maxGuessCount),
	}

	err := game.Reset(letterCount, sourceFile, acceptedFile)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// InitializeGame creates a new game, initializing with a target word and allowed words.
func InitializeGame(maxGuessCount int, letterCount int, sourceFile string, acceptedFile string, targetWord string) (*Game, error) {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		config: &GameConfig{
			maxGuessCount: maxGuessCount,
			letterCount:   letterCount,
			sourceFile:    sourceFile,
			acceptedFile:  acceptedFile,
		},
	}

	err := game.resetWithTarget(letterCount, sourceFile, acceptedFile, targetWord)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// Target gets the target word.
func (g *Game) Target() string {
	return g.target.word
}

// Knowledge gets the knowledge of the game.
func (g *Game) Knowledge() *TargetKnowledge {
	return g.knowledge
}

// CanGuess returns true if a guess can be made.
func (g *Game) CanGuess() bool {
	return g.knowledge.Results[g.config.maxGuessCount-1] == nil
}

// Check checks if a supplied word matches the target.
func (g *Game) Check(word string) (*CheckResult, *TargetKnowledge, error) {
	if g.knowledge.Results[g.config.maxGuessCount-1] != nil {
		return nil, nil, errors.Errorf("maximum number of guesses reached")
	}

	word = normalizeWord(word)
	res, err := g.target.Check(word)
	if err != nil {
		return nil, nil, err
	}
	if !g.allowedWords[word] {
		return nil, nil, errors.Errorf("'%s' not in dictionary of words loaded", word)
	}
	g.knowledge.AddCheckResult(res)
	return res, g.knowledge, nil
}

// Reset returns a game to the initial state, selecting a target and reading
// allowed words.
func (g *Game) Reset(letterCount int, sourceFile string, acceptedFile string) error {
	return g.resetWithTarget(letterCount, sourceFile, acceptedFile, "")
}

// Reset returns a game to the initial state, selecting a target and reading
// allowed words.
func (g *Game) resetWithTarget(letterCount int, sourceFile string, acceptedFile string, targetWord string) error {
	g.guessCount = 0
	g.config.letterCount = letterCount
	g.config.sourceFile = sourceFile
	g.config.acceptedFile = acceptedFile

	words, err := g.readWords(g.config.sourceFile, g.config.letterCount)
	if err != nil {
		return nil
	}

	targets, err := g.readWords(g.config.acceptedFile, g.config.letterCount)
	if err != nil {
		return nil
	}

	if targetWord == "" {
		index := int(rand.Float64() * float64(len(targets)))
		targetWord = targets[index]
	}
	target, err := NewWord(targetWord)
	if err != nil {
		return nil
	}
	g.target = target

	allowed := map[string]bool{}
	for _, w := range words {
		allowed[w] = true
	}
	g.allowedWords = allowed
	g.knowledge = NewTargetKnowledge(g.config.maxGuessCount)
	//fmt.Printf("TARGET: %v\n", g.target)
	fmt.Printf("LOADED %v WORDS OF LENGTH %v\n", len(allowed), g.config.letterCount)

	return nil
}

func (g *Game) readWords(filename string, count int) ([]string, error) {
	log.Infof("reading words with %d characters from %s", count, filename)
	wordsRaw, err := readWords(filename, count)
	if err != nil {
		return nil, err
	}

	proper := []string{}
	for _, w := range wordsRaw {
		if len(w) == count {
			proper = append(proper, normalizeWord(w))
		}
	}
	log.Infof("read %d words with %d characters", len(proper), count)

	return proper, nil
}

func readWords(sourceFile string, letterCount int) ([]string, error) {
	file, err := os.Open(sourceFile)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open source file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read lines")
	}
	return lines, nil
}
