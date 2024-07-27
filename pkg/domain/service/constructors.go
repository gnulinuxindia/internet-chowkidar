// Code generated by tools/update-providers.go; DO NOT EDIT.
package service

import "github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"

func ProvideCounterService(
	counterRepository repository.CounterRepository,
	emailService EmailService,
) CounterService {
	return &counterServiceImpl{
		counterRepository: counterRepository,

		emailService: emailService,
	}
}

func ProvideEmailService() EmailService {
	return &emailServiceImpl{}
}
