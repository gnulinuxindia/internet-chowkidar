// Code generated by ogen, DO NOT EDIT.

package genapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateAbuseReport implements createAbuseReport operation.
//
// Create a new abuse report.
//
// POST /abuse-reports
func (UnimplementedHandler) CreateAbuseReport(ctx context.Context, req *AbuseReportInput) (r *AbuseReport, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateBlock implements createBlock operation.
//
// Create a new block.
//
// POST /blocks
func (UnimplementedHandler) CreateBlock(ctx context.Context, req *BlockInput) (r *Block, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateISP implements createISP operation.
//
// Create a new ISP.
//
// POST /isps
func (UnimplementedHandler) CreateISP(ctx context.Context, req *ISPInput) (r *ISP, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSite implements createSite operation.
//
// Create a new site.
//
// POST /sites
func (UnimplementedHandler) CreateSite(ctx context.Context, req *SiteInput) (r *Site, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSiteSuggestion implements createSiteSuggestion operation.
//
// Create a new site suggestion.
//
// POST /sites/suggestions
func (UnimplementedHandler) CreateSiteSuggestion(ctx context.Context, req *SiteSuggestionInput) (r *SiteSuggestion, _ error) {
	return r, ht.ErrNotImplemented
}

// ListAbuseReports implements listAbuseReports operation.
//
// List all abuse reports.
//
// GET /abuse-reports
func (UnimplementedHandler) ListAbuseReports(ctx context.Context) (r []AbuseReport, _ error) {
	return r, ht.ErrNotImplemented
}

// ListBlocks implements listBlocks operation.
//
// List all blocks.
//
// GET /blocks
func (UnimplementedHandler) ListBlocks(ctx context.Context) (r []Block, _ error) {
	return r, ht.ErrNotImplemented
}

// ListISPs implements listISPs operation.
//
// List all ISPs.
//
// GET /isps
func (UnimplementedHandler) ListISPs(ctx context.Context) (r []ISP, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSiteSuggestions implements listSiteSuggestions operation.
//
// List all site suggestions.
//
// GET /sites/suggestions
func (UnimplementedHandler) ListSiteSuggestions(ctx context.Context) (r []SiteSuggestion, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSites implements listSites operation.
//
// List all sites.
//
// GET /sites
func (UnimplementedHandler) ListSites(ctx context.Context) (r []Site, _ error) {
	return r, ht.ErrNotImplemented
}
