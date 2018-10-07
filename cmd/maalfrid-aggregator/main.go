package main

import (
	"os"
	log "github.com/inconshreveable/log15"
	"github.com/namsral/flag"
	"github.com/nlnwa/pkg/logfmt"
	"github.com/nlnwa/maalfrid-aggregator/version"
	"github.com/nlnwa/maalfrid-aggregator/pkg/aggregator"
)

func main() {
	port := 8000
	dbHost := "localhost"
	dbPort := 28015
	dbName := "test"
	dbUser := "admin"
	dbPassword := ""
	debug := false

	flag.IntVar(&port, "port", port, "gRPC server listening port")
	flag.StringVar(&dbHost, "db-host", dbHost, "database host")
	flag.IntVar(&dbPort, "db-port", dbPort, "database port")
	flag.StringVar(&dbName, "db-name", dbName, "database name")
	flag.StringVar(&dbUser, "db-user", dbUser, "database user")
	flag.StringVar(&dbName, "db-password", dbPassword, "database password")
	flag.BoolVar(&debug, "debug", debug, "enable debugging")
	flag.Parse()

	logger := log.New()
	logHandler := log.CallerFuncHandler(log.StreamHandler(os.Stdout, logfmt.LogbackFormat()))
	if debug {
		logger.SetHandler(log.CallerStackHandler("%+v", logHandler))
	} else {
		logger.SetHandler(log.LvlFilterHandler(log.LvlInfo, logHandler))
	}

	logger.Info(version.String())


	api, err := aggregator.NewApi()
	srv, err := NewServer(api)
}
