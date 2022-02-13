package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/phorne-uncharted/guess-lang/board"
)

func main() {
	runGame()
}

func getBoolFromUser(message string) (bool, error) {
	fmt.Println(message)
	answer := captureWord()

	return answer == "yes" || answer == "y", nil
}

func getIntFromUser(message string) (int, error) {
	fmt.Println(message)
	lengthStr := captureWord()
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to read number from user")
	}

	return length, nil
}

func runGame() {
	more := true
	for more {
		length, err := getIntFromUser("how long should the word be?")
		if err != nil {
			fmt.Printf("%+v", err)
			return
		}
		g, _ := board.NewGame(5, length, "fr.txt", "fr.txt")
		solveWord(g)
		more, _ = getBoolFromUser("continue playing?")
	}
}

func solveWord(g *board.Game) {
	var err error
	for solved := false; !solved; {
		solved, err = attemptSolve(g)
		if err != nil {
			handleError(err)
		}
	}
}

func attemptSolve(g *board.Game) (bool, error) {
	wordTest := captureWord()
	r, tk, err := g.Check(wordTest)
	if err != nil {
		return false, err
	}

	for _, c := range r.Comparison {
		fmt.Printf("%v", string(c.SourceChar))
	}
	fmt.Printf("\t\t")
	for b := byte(65); b < 91; b++ {
		fmt.Printf("%v ", string(b))
	}
	fmt.Printf("\n")

	for _, c := range r.Comparison {
		fmt.Printf(mapResult(c.Result, " "))
	}
	fmt.Printf("\t\t")
	for b := 65; b < 91; b++ {
		fmt.Printf("%v ", mapResult(tk.Letters[b], "X"))
	}
	fmt.Printf("\n")

	return r.IsSolved(), nil
}

func mapResult(compare board.CompareResult, noMatch string) string {
	if compare == board.DontKnow {
		return " "
	} else if compare == board.InWord {
		return "-"
	} else if compare == board.AtPlace {
		return "+"
	}
	return noMatch
}

func handleError(err error) {
	fmt.Printf("%v\n", err.Error())
}

func captureWord() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
