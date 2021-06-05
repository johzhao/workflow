package database

import "workflow/database/database/repository"

type Database interface {
	Close() error

	Begin() (Database, error)
	Rollback() error
	Commit() error

	GetActionRepository() repository.Action
}
