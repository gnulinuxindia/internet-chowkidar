package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

type CounterRepository interface {
}

type counterRepositoryImpl struct {
	db *ent.Client
}
