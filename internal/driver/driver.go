package driver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

// OpenDB opens a database connection
func OpenDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err!= nil {
        return nil, err
    }
	
	err = db.Ping()
	if err!= nil {
		fmt.Println(err)
        return nil, err
    }
	return db, nil
}
