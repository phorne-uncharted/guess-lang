package storage

import (
	"context"
	"fmt"
	"sync"

	pool "github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	log "github.com/unchartedsoftware/plog"
)

var (
	mu      = &sync.Mutex{}
	clients map[string]*pool.Pool
)

func init() {
	clients = make(map[string]*pool.Pool)
}

// Storage is used to query the stored game data.
type Storage struct {
	conn *pool.Pool
}

// NewDataStorage returns a constructor for a data storage.
func NewDataStorage(clientCtor func() (*pool.Pool, error)) func() (*Storage, error) {
	return func() (*Storage, error) {
		client, err := clientCtor()
		if err != nil {
			return nil, err
		}

		return &Storage{
			conn: client,
		}, nil
	}
}

// InitializeDatabase sets up the database tables.
func (s *Storage) InitializeDatabase() error {
	gameSQL := fmt.Sprintf(createGameTable, fmt.Sprintf("IF NOT EXISTS %s", gameTableName))
	_, err := s.conn.Exec(context.Background(), gameSQL)
	if err != nil {
		return errors.Wrapf(err, "unable to init game table")
	}

	guessSQL := fmt.Sprintf(createGuessTable, fmt.Sprintf("IF NOT EXISTS %s", guessTableName))
	_, err = s.conn.Exec(context.Background(), guessSQL)
	if err != nil {
		return errors.Wrapf(err, "unable to init guess table")
	}

	return nil
}

// NewClient instantiates and returns a new postgres client constructor.  Log level is one
// of none, info, warn, error, debug.
func NewClient(host string, port int, user string, password string, database string) func() (*pool.Pool, error) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d",
		user, password, host, port, database, 1)
	return NewClientFromConnectionString(connString)
}

// NewClientFromConnectionString instantiates and returns a new postgres client constructor.
func NewClientFromConnectionString(connectionString string) func() (*pool.Pool, error) {
	return func() (*pool.Pool, error) {

		mu.Lock()
		defer mu.Unlock()

		// see if we have an existing connection
		pgxClient, ok := clients[connectionString]
		if !ok {
			poolConfig, err := pool.ParseConfig(connectionString)
			if err != nil {
				return nil, errors.Wrap(err, "unable to parse postgres config")
			}
			endpoint := fmt.Sprintf("%s:%d", poolConfig.ConnConfig.Host, poolConfig.ConnConfig.Port)
			log.Infof("Creating new Postgres connection to connection %s", endpoint)
			poolConfig.LazyConnect = false
			poolConfig.MaxConns = 1
			// BuildStatementCache set to nil prevents the caching of queries
			// This does slow down performance when multiple of the same query is ran
			// However, this also causes issues when types are changing and the caches are not updated
			// One solution would be to reset all pool connection every time a type is changed (but for now this seems to be the best way)
			poolConfig.ConnConfig.BuildStatementCache = nil
			//TODO: Need to close the pool eventually. Not sure how to hook that in.
			pgxClient, err = pool.ConnectConfig(context.Background(), poolConfig)
			if err != nil {
				return nil, errors.Wrap(err, "Postgres client init failed")
			}
			log.Infof("Postgres connection established to endpoint %s", endpoint)
			clients[endpoint] = pgxClient
		}
		return pgxClient, nil
	}
}
