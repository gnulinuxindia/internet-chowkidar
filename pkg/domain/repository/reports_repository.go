package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

type ReportsRepository interface {
}

type reportsRepositoryImpl struct {
	db *ent.Client
}
