package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"paulo/parse-log-apache/domain"
	"paulo/parse-log-apache/util/constants"
	"strings"
)

func Connect() *sql.DB {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		constants.DbHost,
		constants.DbPort,
		constants.DbUser,
		constants.DbPassword,
		constants.DbName,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Print("Postgresql successfully connected!")

	return db

}

func getArgsBulkInset(logs []domain.LogEntry) []interface{} {
	valueArgs := []interface{}{}

	for _, l := range logs {

		valueArgs = append(
			valueArgs,
			l.IP,
			l.Date,
			l.Method,
			l.Resource,
			l.Version,
			l.Status,
			l.Size,
			l.Referer,
			l.UserAgent,
		)
	}

	return valueArgs

}

func getQueryBulkInsert(qtyRows int) string {

	valueStrings := []string{}

	for i := 0; i < qtyRows; i++ {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			i*9+1,
			i*9+2,
			i*9+3,
			i*9+4,
			i*9+5,
			i*9+6,
			i*9+7,
			i*9+8,
			i*9+9,
		))
	}

	stmt := fmt.Sprintf(
		`INSERT INTO logs(ip, date, method, resource, version, status, size, referer, user_agent)
    VALUES %s`,
		strings.Join(valueStrings, ","),
	)

	return stmt
}
