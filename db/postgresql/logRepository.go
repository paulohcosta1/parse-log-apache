package postgresql

import (
	"database/sql"
	"log"
	"paulo/parse-log-apache/domain"

	_ "github.com/lib/pq"
)

type LogEntryRepository struct {
	Db *sql.DB
}

func (r LogEntryRepository) BulkInsert(logs []domain.LogEntry) error {

	tx, err := r.Db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	query := getQueryBulkInsert(len(logs))
	args := getArgsBulkInset(logs)
	_, err = tx.Exec(query, args...)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
