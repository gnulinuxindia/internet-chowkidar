package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

type IspRepository interface {
}

type ispRepositoryImpl struct {
	db *ent.Client
}
