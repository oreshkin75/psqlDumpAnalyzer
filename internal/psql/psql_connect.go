package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (c *Creator) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.config.PsqlUser, c.config.PsqlPassword, c.config.PsqlDBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.logError.Print(err)
		return nil, err
	}
	c.db = db
	return db, nil
}

func (c *Creator) Disconnect(db *sql.DB) {
	db.Close()
}
