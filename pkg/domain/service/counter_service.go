package service

import (
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type CounterService interface {
}

type counterServiceImpl struct {
	counterRepository repository.CounterRepository
	emailService      EmailService
}
