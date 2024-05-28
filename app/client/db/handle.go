package db

type Repository interface {
	// methods of the clients/[db, kafka, redis, etc..]
	Start() error
	Stop() error
}
