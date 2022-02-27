package config

import (
	"noteapp/exception"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(c Config) *sqlx.DB {
	dsn := "host=" + c.Get("POSTGRES_HOST") +
		" port=" + c.Get("POSTGRES_PORT") +
		" user=" + c.Get("POSTGRES_USER") +
		" password=" + c.Get("POSTGRES_PASS") +
		" dbname=" + c.Get("POSTGRES_DB") +
		" sslmode=disable"
	// fmt.Println(dsn)
	db, err := sqlx.Connect("postgres", dsn)
	exception.PanicIfNeeded(err)
	defer db.Close()
	return db
}
