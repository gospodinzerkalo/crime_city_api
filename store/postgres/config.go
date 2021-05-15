package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	storeDb "github.com/gospodinzerkalo/crime_city_api/store"
	"log"
	"os"
)

type Config struct {
	Host             string
	Port             string
	User             string
	Password         string
	Database         string
	Params           string
	ConnectionString string
}

func getConnString(cfg Config) string {
	var connStr string
	if cfg.ConnectionString == "" {
		connStr = "postgres://"
		if cfg.Host == "" {
			cfg.Host = "localhost"
		}
		if cfg.Port == "" {
			cfg.Port = "5432"
		}
		connStr +=
			cfg.User + ":" +
				cfg.Password + "@" +
				cfg.Host + ":" +
				cfg.Port + "/" +
				cfg.Database + "?" +
				cfg.Params
		//connStr = "postgres://dev_dms:47EnI5eVngiW1j63@dev-postgresql-v965.c5uhwnscdeev.eu-west-1.rds.amazonaws.com/dms_business"
	} else {
		connStr = cfg.ConnectionString
	}

	return connStr
}

func getDbConn(connStr string, logger pgx.Logger) (*sql.DB, error) {
	connConfig, _ := pgx.ParseConfig(connStr)
	connConfig.Logger = logger
	connConfig.PreferSimpleProtocol = true
	connStr2 := stdlib.RegisterConnConfig(connConfig)
	db, err := sql.Open("pgx", connStr2)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type store struct {
	db  *sql.DB
	log logrus.Logger
}

var dbTables = []string{
	`CREATE TABLE IF NOT EXISTS crimes
	(
		id SERIAL,
		location_name TEXT NOT NULL,
		longitude NUMBER,
		latitude NUMBER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		description TEXT,
		image TEXT
	);`,
}

func New(cfg Config) (storeDb.Repository, error) {

	logr := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.Level(5),
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "02-01-06 15:04:05.000000",
		},
	}

	db, err := getDbConn(getConnString(cfg), logr)
	if err != nil {
		panic(fmt.Sprintf("fatal error config file: %s ", err))
	}

	return &store{db: db, log: *logr}, nil
}