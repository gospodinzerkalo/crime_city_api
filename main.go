package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/gospodinzerkalo/crime_city_api/service"
	"github.com/gospodinzerkalo/crime_city_api/store/postgres"
	"github.com/gospodinzerkalo/crime_city_api/transport"
	"github.com/joho/godotenv"
	"github.com/oklog/oklog/pkg/group"
	"github.com/urfave/cli/v2"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath 				= "./"
	postgreDb				= ""
	postgreUser 			= ""
	postgrePass 			= ""
	postgreHost 			= ""
	postgrePort 			= "5432"

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config, c",
			Usage:       "path to .env config file",
			Aliases: []string{"c"},
			Destination: &configPath,
		},
	}
)

func parseEnv() {
	if configPath != "" {
		godotenv.Overload(configPath)
	}

	postgreHost = os.Getenv("POSTGRES_HOST")
	postgreDb = os.Getenv("POSTGRES_DATABASE")
	postgrePass = os.Getenv("POSTGRES_PASSWORD")
	postgreUser = os.Getenv("POSTGRES_USER")
}

func main() {
	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "start",
			Action: run,
			Flags: flags,
			Usage: "start server",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	httpAddr := flag.String("http-addr", ":8080", "HTTP listen address")
	flag.Parse()

	parseEnv()
	postgresConfig := postgres.Config{
		Host:             postgreHost,
		Port:             postgrePort,
		User:             postgreUser,
		Password:         postgrePass,
		Database:         postgreDb,
		Params:           "",
		ConnectionString: "",
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	store, err := postgres.New(postgresConfig)
	if err != nil {
		return err
	}

	svc := service.NewService(store, logger)
	endpoints := endpoint.NewEndpointsFactory(svc, logger)
	httpHandler := transport.MakeHTTPHandler(endpoints, logger)


	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", *httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
	return nil
}