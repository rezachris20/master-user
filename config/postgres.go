package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
)

func ReadPostgres() (*pg.DB, error) {
	return CreateDBConnectionPostgres(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER_POSTGRES"),
		os.Getenv("DB_PASSWORD_POSTGRES"),
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES")))
}

func WritePostgres() (*pg.DB, error) {
	return CreateDBConnectionPostgres(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER_POSTGRES"),
		os.Getenv("DB_PASSWORD_POSTGRES"),
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES")))
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {

	qq, _ := q.FormattedQuery()
	fmt.Println(string(qq))
	return nil
}

func CreateDBConnectionPostgres(url string) (*pg.DB, error) {

	opt, err := pg.ParseURL(url)
	if err != nil {
		log.Info("Invalid database URL!")
		log.Fatal(err)
	}

	opt.MaxConnAge = 100
	opt.IdleTimeout = 5
	opt.PoolSize = 50

	client := pg.Connect(opt)
	client.AddQueryHook(dbLogger{})

	err = client.Ping(context.TODO())
	if err != nil {
		log.Info("not connect database")
		log.Fatal(err)
		defer client.Close()
	}

	log.Info("Connected to Postgres!")
	return client, err

}
