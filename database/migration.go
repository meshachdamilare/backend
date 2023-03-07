package database

import (
	"context"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

// this is a custom migration function i use for sql dbs
func Migrate(db *gorm.DB) {
	const TOTAL_WORKERS = 1
	var (
		wg      sync.WaitGroup
		errorCh = make(chan error, TOTAL_WORKERS)
	)
	wg.Add(TOTAL_WORKERS)
	log.Println("running db migration :::::::::::::")

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
		defer cancel()
		// check if table exist before creating
		tableExist, err := checkTableExist(ctx, db, "users")
		if err != nil {
			errorCh <- err
		}
		if !tableExist {
			query := `CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`
			err := db.Exec(query).Error
			if err != nil {
				errorCh <- err
			}
		}
	}()

	// more go routines can be added here and number of TOTAL_WORKERS increased to handle other tables

	go func() {
		wg.Wait()
		close(errorCh)
	}()

	for err := range errorCh {
		if err != nil {
			panic(err)
		}
	}

	log.Println("complete db migration")
}

// check if a table exist in the pg db
func checkTableExist(ctx context.Context, db *gorm.DB, tableName string) (bool, error) {
	query := `
		SELECT EXISTS (
   SELECT FROM pg_tables
   WHERE  schemaname = 'public'
   AND    tablename  = $1
   );
	`
	row := db.Raw(query, tableName)
	var response bool
	_ = row.Scan(&response)
	return response, nil
}
