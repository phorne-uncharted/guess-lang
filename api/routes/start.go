package routes

import (
	"fmt"
	"net/http"

	"github.com/phorne-uncharted/guess-lang/api/storage"
	"github.com/phorne-uncharted/guess-lang/board"
	"github.com/pkg/errors"
)

// StartHandler creates a handler that starts games.
func StartHandler(storage *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := getPostParameters(r)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to parse post parameters"))
			return
		}

		letterCount := int(params["letterCount"].(float64))
		maxGuessCount := int(params["maxGuessCount"].(float64))
		lang := params["language"].(string)
		sourceFile := fmt.Sprintf("public/resource/words/%s.txt", lang)
		acceptedFile := fmt.Sprintf("public/resource/words/accepted-%s.txt", lang)

		game, err := board.NewGame(maxGuessCount, letterCount, sourceFile, acceptedFile)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to start the game"))
			return
		}

		gameID, err := storage.StoreGame(lang, game.Target(), sourceFile, acceptedFile, maxGuessCount)
		if err != nil {
			handleError(w, errors.Wrap(err, "Unable to store game in storage"))
			return
		}

		// marshal data
		err = handleJSON(w, map[string]interface{}{"gameId": gameID})
		if err != nil {
			handleError(w, errors.Wrap(err, "unable to marshal result into JSON"))
			return
		}
	}
}
