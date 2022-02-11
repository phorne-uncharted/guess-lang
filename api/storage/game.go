package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/phorne-uncharted/guess-lang/board"
)

const (
	gameTableName  = "game"
	guessTableName = "guess"

	createGameTable = `CREATE TABLE %s (
    game_id SERIAL,
    language TEXT,
    target TEXT,
    source_file TEXT,
    max_guess_count INT
    );`

	createGuessTable = `CREATE TABLE %s (
      game_id INT,
      guess_count INT,
      word TEXT
      );`
)

// StoreGame stores a game to storage.
func (s *Storage) StoreGame(language string, target string, sourceFile string, maxGuessCount int) (int, error) {
	sql := fmt.Sprintf("INSERT INTO %s (language, target, source_file, max_guess_count) VALUES ($1, $2, $3, $4) RETURNING game_id;", gameTableName)
	row := s.conn.QueryRow(context.Background(), sql, language, target, sourceFile, maxGuessCount)

	var gameID int
	err := row.Scan(&gameID)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to read game id when storing a new game")
	}

	return gameID, nil
}

// LoadGame loads game data from the database.
func (s *Storage) LoadGame(gameID int) (*board.Game, error) {
	sql := fmt.Sprintf("SELECT language, target, source_file, max_guess_count FROM %s WHERE game_id = $%d;", gameTableName, 1)
	res, err := s.conn.Query(context.Background(), sql, gameID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to query for game data")
	}

	return parseGame(res)
}

func parseGame(rows pgx.Rows) (*board.Game, error) {
	if !rows.Next() {
		return nil, errors.Errorf("game does not exist")
	}

	var language string
	var target string
	var sourceFile string
	var maxGuessCount int
	err := rows.Scan(&language, &target, &sourceFile, &maxGuessCount)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read game data")
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error reading game data")
	}

	return board.InitializeGame(maxGuessCount, len(target), sourceFile, target)
}
