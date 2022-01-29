package board

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

// GameConfig defines basic configurations used for the game.
type GameConfig struct {
	maxGuessCount int
	letterCount   int
	sourceFile    string
}

// Game is the main driver.
type Game struct {
	target       *Word
	config       *GameConfig
	allowedWords map[string]bool
	guessCount   int
}

// NewGame creates a new game, initializing with a target word and allowed words.
func NewGame(maxGuessCount int, letterCount int, sourceFile string) (*Game, error) {
	game := &Game{
		config: &GameConfig{
			maxGuessCount: maxGuessCount,
			letterCount:   letterCount,
			sourceFile:    sourceFile,
		},
	}

	err := game.Reset(letterCount, sourceFile)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// Reset returns a game to the initial state, selecting a target and reading
// allowed words.
func (g *Game) Reset(letterCount int, sourceFile string) error {
	g.guessCount = 0
	g.config.letterCount = letterCount
	g.config.sourceFile = sourceFile

	words, err := g.readWords()
	if err != nil {
		return nil
	}

	index := int(rand.Float64() * float64(len(words)))
	target, err := NewWord(words[index])
	if err != nil {
		return nil
	}
	g.target = target

	allowed := map[string]bool{}
	for _, w := range words {
		allowed[w] = true
	}
	g.allowedWords = allowed

	return nil
}

func (g *Game) readWords() ([]string, error) {
	return readWords(g.config.sourceFile, g.config.letterCount)
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
