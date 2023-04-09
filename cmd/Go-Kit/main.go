package main

import (
	"database/sql"
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"os"
)

const dbSource = "postgres://GoKit:GoKit@db/GoKit?sslmode=disable"

func main() {
	var _ = flag.String("http", "8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	_ = level.Info(logger).Log("msg", "service started")
	defer func(logger log.Logger, message ...interface{}) {
		_ = logger.Log(message)
	}(level.Info(logger), "msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbSource)
		if err != nil {
			_ = level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		err = db.Close()
		if err != nil {
			return
		}
	}
}
