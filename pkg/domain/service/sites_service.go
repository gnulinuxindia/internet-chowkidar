package service

import (
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type SitesService interface {
}

type sitesServiceImpl struct {
	repo repository.SitesRepository
}
