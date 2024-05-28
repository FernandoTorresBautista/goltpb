package biz

import (
	"log"

	"goltpb/app/client/db"
)

// Biz structure with all escential elements
type Biz struct {
	logger *log.Logger
	db     db.Repository
}

// New return a new biz instance
func New(logger *log.Logger, db db.Repository) *Biz {
	return &Biz{
		logger: logger,
		db:     db,
	}
}

// Handle implementations ...
type Handle interface {
	// db
}

// Start all elements of the biz
func (b *Biz) Start() error {
	// initialize the Biz
	if b.db != nil {
		err := b.db.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop all elements of the biz
func (b *Biz) Stop() error {
	// Stop the Biz
	if b.db != nil {
		err := b.db.Stop()
		if err != nil {
			return err
		}
	}

	return nil
}

// implement functions
