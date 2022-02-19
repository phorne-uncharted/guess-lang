package routes

import (
	"net/http"

	"github.com/phorne-uncharted/guess-lang/api/storage"
	"github.com/pkg/errors"
	log "github.com/unchartedsoftware/plog"
)

// GuessHandler processes a guess at the target word and returns the result.
func GuessHandler(storage *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := getPostParameters(r)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to parse post parameters"))
			return
		}

		gameID := int(params["gameId"].(float64))
		word := params["word"].(string)

		game, err := storage.LoadGame(gameID)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to load game from storage"))
			return
		}
		log.Infof("loaded game %d from storage", gameID)

		guesses, err := storage.LoadGuesses(gameID)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to load guesses from storage"))
			return
		}
		log.Infof("loaded guesses for game %d from storage", gameID)

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

		err = storage.StoreGuess(gameID, word, len(guesses)+1)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to store guess in storage"))
			return
		}
		log.Infof("stored guess for game %d from storage", gameID)

		response := map[string]interface{}{"done": cr.IsSolved() || !game.CanGuess(), "solved": cr.IsSolved(), "check": cr, "knowledge": tk}
		if !game.CanGuess() {
			response["target"] = game.Target()
		}

		// marshal data
		err = handleJSON(w, response)
		if err != nil {
			handleError(w, errors.Wrap(err, "unable to marshal result into JSON"))
			return
		}
	}
}
