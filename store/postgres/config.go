package postgres

import (
	"database/sql"
	"fmt"
	storeDb "github.com/gospodinzerkalo/crime_city_api/store"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
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

func getDbConn(connStr string, logger logrus.Logger) (*sql.DB, error) {
	connConfig, _ := pgx.ParseConfig(connStr)
	connConfig.Logger = logrusadapter.NewLogger(&logger)

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
		id SERIAL PRIMARY KEY,
		location_name TEXT NOT NULL,
		longitude NUMERIC,
		latitude NUMERIC,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		description TEXT,
		image TEXT
	);`,
	`CREATE TABLE IF NOT EXISTS homes
	(
		id INT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		user_name TEXT,
		longitude NUMERIC,
		latitude NUMERIC,
		image TEXT
	);`,
	// function that calculates between two lat and lon
	`CREATE OR REPLACE FUNCTION calculate_distance(lat1 float, lon1 float, lat2 float, lon2 float, units varchar)
		RETURNS float AS $dist$
			DECLARE
				dist float = 0;
				radlat1 float;
				radlat2 float;
				theta float;
				radtheta float;
			BEGIN
				IF lat1 = lat2 OR lon1 = lon2
					THEN RETURN dist;
				ELSE
					radlat1 = pi() * lat1 / 180;
					radlat2 = pi() * lat2 / 180;
					theta = lon1 - lon2;
					radtheta = pi() * theta / 180;
					dist = sin(radlat1) * sin(radlat2) + cos(radlat1) * cos(radlat2) * cos(radtheta);
		
					IF dist > 1 THEN dist = 1; END IF;
		
					dist = acos(dist);
					dist = dist * 180 / pi();
					dist = dist * 60 * 1.1515;
		
					IF units = 'K' THEN dist = dist * 1.609344; END IF;
					IF units = 'N' THEN dist = dist * 0.8684; END IF;
		
					RETURN dist;
				END IF;
			END;
		$dist$ LANGUAGE plpgsql;`,
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

	db, err := getDbConn(getConnString(cfg), *logr)
	if err != nil {
		panic(fmt.Sprintf("fatal error config file: %s ", err))
	}

	for _, q := range dbTables {
		_, err = db.Exec(q)
		if err != nil {
			return nil, err
		}
	}

	return &store{db: db, log: *logr}, nil
}