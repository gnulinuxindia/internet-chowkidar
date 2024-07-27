package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
)

type ReportsHandler interface {
	ListAbuseReports(ctx context.Context) ([]genapi.AbuseReport, error)
	CreateAbuseReport(ctx context.Context, req *genapi.AbuseReportInput) (*genapi.AbuseReport, error)
}

type reportsHandlerImpl struct {
}

func (r *reportsHandlerImpl) ListAbuseReports(ctx context.Context) ([]genapi.AbuseReport, error) {
	panic("not implemented")
}

func (r *reportsHandlerImpl) CreateAbuseReport(ctx context.Context, req *genapi.AbuseReportInput) (*genapi.AbuseReport, error) {
	panic("not implemented")
}
