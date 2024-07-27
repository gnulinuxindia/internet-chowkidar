package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

type SitesRepository interface {
}

type sitesRepositoryImpl struct {
	db *ent.Client
}
