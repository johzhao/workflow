package _interface

type Database interface {
	Close() error

	Begin() (Database, error)
	Rollback() error
	Commit() error

	GetActionRepository() ActionRepository
}
