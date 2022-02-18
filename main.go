package main

import (
	"net/http"
	"os"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	pool "github.com/jackc/pgx/v4/pgxpool"
	log "github.com/unchartedsoftware/plog"
	"github.com/zenazn/goji/graceful"
	goji "goji.io/v3"
	"goji.io/v3/pat"

	"github.com/phorne-uncharted/guess-lang/api/env"
	"github.com/phorne-uncharted/guess-lang/api/middleware"
	"github.com/phorne-uncharted/guess-lang/api/routes"
	"github.com/phorne-uncharted/guess-lang/api/storage"
)

func registerRoute(mux *goji.Mux, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	log.Infof("Registering GET route %s", pattern)
	mux.HandleFunc(pat.Get(pattern), handler)
}

func registerRoutePost(mux *goji.Mux, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	log.Infof("Registering POST route %s", pattern)
	mux.HandleFunc(pat.Post(pattern), handler)
}

func main() {
	// load config from env
	config, err := env.LoadConfig()
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}
	log.Infof("%+v", spew.Sdump(config))

	var clientCtor func() (*pool.Pool, error)
	if config.DatabaseURL != "" {
		clientCtor = storage.NewClientFromConnectionString(config.DatabaseURL)
	} else {
		clientCtor = storage.NewClient(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDatabase)
	}

	client, err := clientCtor()
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}

	log.Infof("read configuration from the environment")
	storageCtor := storage.NewDataStorage(client)
	stor, err := storageCtor()
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}

	s, err := storageCtor()
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}
	err = s.InitializeDatabase()
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}
	log.Infof("connected to postgres database")

	// register routes
	mux := goji.NewMux()
	mux.Use(middleware.Log)
	mux.Use(middleware.Gzip)
	registerRoutePost(mux, "/game/start", routes.StartHandler(stor))
	registerRoutePost(mux, "/game/guess", routes.GuessHandler(stor))

	registerRoute(mux, "/*", routes.FileHandler("./dist"))

	// catch kill signals for graceful shutdown
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)

	// kick off the server listen loop
	log.Infof("Listening on port %s", config.AppPort)
	err = graceful.ListenAndServe(":"+config.AppPort, mux)
	if err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}

	// wait until server gracefully exits
	graceful.Wait()
}
