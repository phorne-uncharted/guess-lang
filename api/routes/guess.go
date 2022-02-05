package routes

import (
	"net/http"

	"github.com/phorne-uncharted/guess-lang/api/storage"
	"github.com/pkg/errors"
)

// GuessHandler processes a guess at the target word and returns the result.
func GuessHandler(storageCtor func() (*storage.Storage, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := getPostParameters(r)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to parse post parameters"))
			return
		}

		gameID := int(params["gameId"].(float64))
		word := params["word"].(string)

		data, err := storageCtor()
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to create storage"))
			return
		}

		game, err := data.LoadGame(gameID)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to load game from storage"))
			return
		}

		guesses, err := data.LoadGuesses(gameID)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to load guesses from storage"))
			return
		}

		for i := 0; i < len(guesses); i++ {
			_, _, err = game.Check(guesses[i])
			if err != nil {
				handleError(w, errors.Wrap(err, "Unable to process previous guesses"))
				return
			}
		}

		cr, tk, err := game.Check(word)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to process guess"))
			return
		}

		err = data.StoreGuess(gameID, word, len(guesses)+1)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to store guess in storage"))
			return
		}

		// marshal data
		err = handleJSON(w, map[string]interface{}{"check": cr, "knowledge": tk})
		if err != nil {
			handleError(w, errors.Wrap(err, "unable to marshal result into JSON"))
			return
		}
	}
}
