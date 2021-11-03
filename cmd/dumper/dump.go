package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"paulo/parse-log-apache/db/postgresql"
	"paulo/parse-log-apache/domain"
	"time"

	"github.com/spf13/cobra"
)

const maxLines = 7000

var parse = new(domain.LogEntry).Parse

func main() {
	db := postgresql.Connect()

	repo := postgresql.LogEntryRepository{Db: db}
	start := time.Now()

	var cmdDump = &cobra.Command{
		Use:   "dump  [log file to db] [conc, seq]",
		Short: "Dump apache log text files into database",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			if args[1] == "conc" {
				dumpConcurrently(args[0], repo)

			} else {
				dumpSequentially(args[0], repo)
			}

		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdDump)
	rootCmd.Execute()
	elapsed := time.Since(start)
	log.Printf("dump took %s", elapsed)
	db.Close()
}

func dumpConcurrently(fileDirectory string, repo domain.LogEntryRepository) error {
	file, err := os.Open(fileDirectory)

	if err != nil {
		log.Print(err)
		return errors.New("fail to open file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	logs := []domain.LogEntry{}
	cont := 0
	contRoutines := 0
	errorc := make(chan error)

	for scanner.Scan() {
		if cont >= maxLines {
			go func(logs []domain.LogEntry) {
				contRoutines++
				errorc <- repo.BulkInsert(logs)
			}(logs)
			cont = 0
			logs = []domain.LogEntry{}
		}

		cont++

		line, err := parse(scanner.Text())
		if err != nil {
			continue
		}

		logs = append(logs, line)
	}

	//insert remaining logs
	if len(logs) > 0 {
		contRoutines++
		go func(logs []domain.LogEntry) {
			errorc <- repo.BulkInsert(logs)
		}(logs)
	}

	errorsInserts := []error{}

	for i := 0; i < contRoutines; i++ {
		if err := <-errorc; err != nil {
			errorsInserts = append(errorsInserts, err)
		}
	}

	file.Close()

	if len(errorsInserts) > 0 {
		for _, e := range errorsInserts {
			log.Println(e)
		}
		return errors.New("fail to insert logs")
	}

	log.Println("Logs successfully inserted")
	return nil

}

func dumpSequentially(fileDirectory string, repo domain.LogEntryRepository) error {

	file, err := os.Open(fileDirectory)

	if err != nil {
		log.Print(err)
		return errors.New("fail to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	logs := []domain.LogEntry{}
	cont := 0
	errorsInserts := []error{}

	for scanner.Scan() {

		if cont >= maxLines {
			errorI := repo.BulkInsert(logs)
			if errorI != nil {
				errorsInserts = append(errorsInserts, errorI)
			}
			cont = 0
			logs = []domain.LogEntry{}
		}
		cont++

		line, err := parse(scanner.Text())

		if err != nil {
			continue
		}

		logs = append(logs, line)

	}
	file.Close()

	if len(errorsInserts) > 0 {
		for _, e := range errorsInserts {
			log.Println(e)
		}
		return errors.New("fail to insert logs")

	}

	log.Println("Logs successfully inserted")
	return nil

}
