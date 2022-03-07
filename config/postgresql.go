package config

import (
	"context"
	"noteapp/exception"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDatabase(c Config) *sqlx.DB {
	dsn := "host=" + c.Get("POSTGRES_HOST") +
		" port=" + c.Get("POSTGRES_PORT") +
		" user=" + c.Get("POSTGRES_USER") +
		" password=" + c.Get("POSTGRES_PASS") +
		" dbname=" + c.Get("POSTGRES_DB") +
		" sslmode=disable"
	// fmt.Println(dsn)
	db, err := sqlx.Connect("postgres", dsn)
	exception.CheckError(err)
	return db
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
