package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// StoreGuess stores a guess to storage.
func (s *Storage) StoreGuess(gameID int, word string, index int) error {
	sql := fmt.Sprintf("INSERT INTO %s (game_id, guess_count, word) VALUES ($1, $2, $3);", guessTableName)
	_, err := s.conn.Exec(context.Background(), sql, gameID, index, word)
	if err != nil {
		return errors.Wrapf(err, "unable to store the guess")
	}

	return nil
}

// LoadGuesses loads guesses from the database.
func (s *Storage) LoadGuesses(gameID int) ([]string, error) {
	sql := fmt.Sprintf("SELECT word FROM %s WHERE game_id = $%d ORDER BY guess_count;", guessTableName, 1)
	res, err := s.conn.Query(context.Background(), sql, gameID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to query for game data")
	}

	return parseGuess(res)
}

func parseGuess(rows pgx.Rows) ([]string, error) {
	guesses := []string{}
	for rows.Next() {
		var guess string
		err := rows.Scan(&guess)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read guesses")
		}
		guesses = append(guesses, guess)
	}
	err := rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error reading guesses")
	}

	return guesses, nil
}
