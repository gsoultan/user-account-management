package database

type Database interface {
	GetConnection() interface{}
}
