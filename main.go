package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/gospodinzerkalo/crime_city_api/middleware"
	"github.com/gospodinzerkalo/crime_city_api/pb"
	"github.com/gospodinzerkalo/crime_city_api/service"
	"github.com/gospodinzerkalo/crime_city_api/store/postgres"
	"github.com/gospodinzerkalo/crime_city_api/transport"
	"github.com/joho/godotenv"
	"github.com/oklog/oklog/pkg/group"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
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
			Name: "rest",
			Action: rest,
			Flags: flags,
			Usage: "start rest server",
		},
		&cli.Command{
			Name: "rpc",
			Action: rpc,
			Flags: flags,
			Usage: "start rpc server",
		},
		&cli.Command{
			Name: "amqp",
			Action: rabbitMQ,
			Flags: flags,
			Usage: "start amqp server",
		},
	}
	fmt.Println(app.Run(os.Args))
}

func rest(c *cli.Context) error {
	parseEnv()
	httpAddr := flag.String("http-addr", ":8080", "HTTP listen address")
	flag.Parse()

	endpoints, logger, err := initConfig()
	if err != nil {
		return err
	}

	httpHandler := transport.MakeHTTPHandler(*endpoints, logger)

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

func rpc(c *cli.Context) error {
	parseEnv()
	endpoints, logger, err := initConfig()
	if err != nil {
		return nil
	}

	grpcServer := transport.NewGRPCServer(*endpoints, logger)


	var g group.Group
	{
		// The gRPC listener mounts the Go kit gRPS handler we created.
		grpcListener, err := net.Listen("tcp", ":50051")
		if err != nil {
			logger.Log("during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			baseServer := grpc.NewServer()
			pb.RegisterCrimeServiceServer(baseServer, grpcServer)
			level.Info(logger).Log("msg", "Server started successfully 🚀")
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
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

func rabbitMQ(c *cli.Context) error {
	parseEnv()
	endpoints, logger, err := initConfig()
	if err != nil {
		return err
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_ = transport.NewRabbitMqServer(*endpoints, logger, transport.RabbitMQConfig{
		Channel: ch,
	})
	var g group.Group
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

func initConfig() (*endpoint.Endpoints, log.Logger, error) {
	postgresConfig := postgres.Config{
		Host:             postgreHost,
		Port:             postgrePort,
		User:             postgreUser,
		Password:         postgrePass,
		Database:         postgreDb,
		Params:           "",
		ConnectionString: "",
	}
	var fieldKeys = []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "CrimeService",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "CrimeService",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "CrimeService",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	store, err := postgres.New(postgresConfig)
	if err != nil {
		return nil, nil, err
	}

	svc := service.NewService(store, logger)
	svc = middleware.LoggingMiddleWare{
		Logger: logger,
		Next:   svc,
	}

	svc = middleware.InstrumentingMiddleware{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
		CountResult:    countResult,
		Next:           svc,
	}
	
	endpoints := endpoint.NewEndpointsFactory(svc, logger)

	return &endpoints, logger, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		os.Exit(1)
	}
}