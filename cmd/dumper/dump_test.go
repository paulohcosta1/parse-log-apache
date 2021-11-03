package main

import (
	"paulo/parse-log-apache/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LogEntryRepositoryMock struct{}

var errBDMock func() error

func (l LogEntryRepositoryMock) BulkInsert(logs []domain.LogEntry) error {
	return errBDMock()
}

func TestDumpConcurrently(t *testing.T) {
	t.Run("should return an error when repository fails", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)
		errBDMock = func() error {
			return assert.AnError
		}
		err := dumpConcurrently("../../logfile.txt", repo)
		assert.Error(t, err)
	})

	t.Run("should return an error if file is not found", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)

		err := dumpConcurrently("invalid.txt", repo)
		assert.Error(t, err)
	})

	t.Run("should return nil if file is successfully dumped", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)
		errBDMock = func() error {
			return nil
		}
		err := dumpConcurrently("../../logfile.txt", repo)
		assert.NoError(t, err)
	})

}
func TestDumpSequentially(t *testing.T) {
	t.Run("should return an error when repository fails", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)
		errBDMock = func() error {
			return assert.AnError
		}
		err := dumpSequentially("../../logfile.txt", repo)
		assert.Error(t, err)
	})

	t.Run("should return an error if file is not found", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)

		err := dumpSequentially("invalid.txt", repo)
		assert.Error(t, err)
	})

	t.Run("should return nil if file is successfully dumped", func(t *testing.T) {
		repo := new(LogEntryRepositoryMock)
		errBDMock = func() error {
			return nil
		}
		err := dumpSequentially("../../logfile.txt", repo)
		assert.NoError(t, err)
	})

}
