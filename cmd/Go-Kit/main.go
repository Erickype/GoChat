package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/Erickype/Go-Kit/account"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbSource = "postgres://gokit:gokit@:5432/GoKit?sslmode=disable"

func main() {
	var httpAddress = flag.String("http", ":8080", "http listen address")
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
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbSource)
		if err != nil {
			_ = level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()

	var srv account.Service
	{
		repository := account.NewRepository(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%c", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddress)
		handler := account.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddress, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
