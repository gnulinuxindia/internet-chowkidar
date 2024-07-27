package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

type BlocksRepository interface {
}

type blocksRepositoryImpl struct {
	db *ent.Client
}
