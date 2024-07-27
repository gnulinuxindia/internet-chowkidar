// Code generated by ogen, DO NOT EDIT.

package genapi

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// CreateAbuseReport invokes createAbuseReport operation.
	//
	// Create a new abuse report.
	//
	// POST /abuse-reports
	CreateAbuseReport(ctx context.Context, request *AbuseReportInput) (*AbuseReport, error)
	// CreateBlock invokes createBlock operation.
	//
	// Create a new block.
	//
	// POST /blocks
	CreateBlock(ctx context.Context, request *BlockInput) (*Block, error)
	// CreateISP invokes createISP operation.
	//
	// Create a new ISP.
	//
	// POST /isps
	CreateISP(ctx context.Context, request *ISPInput) (*ISP, error)
	// CreateSite invokes createSite operation.
	//
	// Create a new site.
	//
	// POST /sites
	CreateSite(ctx context.Context, request *SiteInput) (*Site, error)
	// CreateSiteSuggestion invokes createSiteSuggestion operation.
	//
	// Create a new site suggestion.
	//
	// POST /site-suggestions
	CreateSiteSuggestion(ctx context.Context, request *SiteSuggestionInput) (*SiteSuggestion, error)
	// ListAbuseReports invokes listAbuseReports operation.
	//
	// List all abuse reports.
	//
	// GET /abuse-reports
	ListAbuseReports(ctx context.Context) ([]AbuseReport, error)
	// ListBlocks invokes listBlocks operation.
	//
	// List all blocks.
	//
	// GET /blocks
	ListBlocks(ctx context.Context) ([]Block, error)
	// ListISPs invokes listISPs operation.
	//
	// List all ISPs.
	//
	// GET /isps
	ListISPs(ctx context.Context) ([]ISP, error)
	// ListSiteSuggestions invokes listSiteSuggestions operation.
	//
	// List all site suggestions.
	//
	// GET /site-suggestions
	ListSiteSuggestions(ctx context.Context) ([]SiteSuggestion, error)
	// ListSites invokes listSites operation.
	//
	// List all sites.
	//
	// GET /sites
	ListSites(ctx context.Context) ([]Site, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// CreateAbuseReport invokes createAbuseReport operation.
//
// Create a new abuse report.
//
// POST /abuse-reports
func (c *Client) CreateAbuseReport(ctx context.Context, request *AbuseReportInput) (*AbuseReport, error) {
	res, err := c.sendCreateAbuseReport(ctx, request)
	return res, err
}

func (c *Client) sendCreateAbuseReport(ctx context.Context, request *AbuseReportInput) (res *AbuseReport, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createAbuseReport"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/abuse-reports"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateAbuseReport",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/abuse-reports"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateAbuseReportRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateAbuseReportResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateBlock invokes createBlock operation.
//
// Create a new block.
//
// POST /blocks
func (c *Client) CreateBlock(ctx context.Context, request *BlockInput) (*Block, error) {
	res, err := c.sendCreateBlock(ctx, request)
	return res, err
}

func (c *Client) sendCreateBlock(ctx context.Context, request *BlockInput) (res *Block, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createBlock"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/blocks"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateBlock",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/blocks"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateBlockRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateBlockResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateISP invokes createISP operation.
//
// Create a new ISP.
//
// POST /isps
func (c *Client) CreateISP(ctx context.Context, request *ISPInput) (*ISP, error) {
	res, err := c.sendCreateISP(ctx, request)
	return res, err
}

func (c *Client) sendCreateISP(ctx context.Context, request *ISPInput) (res *ISP, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createISP"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/isps"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateISP",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/isps"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateISPRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateISPResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSite invokes createSite operation.
//
// Create a new site.
//
// POST /sites
func (c *Client) CreateSite(ctx context.Context, request *SiteInput) (*Site, error) {
	res, err := c.sendCreateSite(ctx, request)
	return res, err
}

func (c *Client) sendCreateSite(ctx context.Context, request *SiteInput) (res *Site, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSite"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/sites"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateSite",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/sites"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateSiteRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateSiteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSiteSuggestion invokes createSiteSuggestion operation.
//
// Create a new site suggestion.
//
// POST /site-suggestions
func (c *Client) CreateSiteSuggestion(ctx context.Context, request *SiteSuggestionInput) (*SiteSuggestion, error) {
	res, err := c.sendCreateSiteSuggestion(ctx, request)
	return res, err
}

func (c *Client) sendCreateSiteSuggestion(ctx context.Context, request *SiteSuggestionInput) (res *SiteSuggestion, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSiteSuggestion"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/site-suggestions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateSiteSuggestion",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/site-suggestions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateSiteSuggestionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateSiteSuggestionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListAbuseReports invokes listAbuseReports operation.
//
// List all abuse reports.
//
// GET /abuse-reports
func (c *Client) ListAbuseReports(ctx context.Context) ([]AbuseReport, error) {
	res, err := c.sendListAbuseReports(ctx)
	return res, err
}

func (c *Client) sendListAbuseReports(ctx context.Context) (res []AbuseReport, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listAbuseReports"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/abuse-reports"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListAbuseReports",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/abuse-reports"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListAbuseReportsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListBlocks invokes listBlocks operation.
//
// List all blocks.
//
// GET /blocks
func (c *Client) ListBlocks(ctx context.Context) ([]Block, error) {
	res, err := c.sendListBlocks(ctx)
	return res, err
}

func (c *Client) sendListBlocks(ctx context.Context) (res []Block, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listBlocks"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/blocks"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListBlocks",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/blocks"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListBlocksResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListISPs invokes listISPs operation.
//
// List all ISPs.
//
// GET /isps
func (c *Client) ListISPs(ctx context.Context) ([]ISP, error) {
	res, err := c.sendListISPs(ctx)
	return res, err
}

func (c *Client) sendListISPs(ctx context.Context) (res []ISP, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listISPs"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/isps"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListISPs",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/isps"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListISPsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListSiteSuggestions invokes listSiteSuggestions operation.
//
// List all site suggestions.
//
// GET /site-suggestions
func (c *Client) ListSiteSuggestions(ctx context.Context) ([]SiteSuggestion, error) {
	res, err := c.sendListSiteSuggestions(ctx)
	return res, err
}

func (c *Client) sendListSiteSuggestions(ctx context.Context) (res []SiteSuggestion, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listSiteSuggestions"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/site-suggestions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListSiteSuggestions",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/site-suggestions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListSiteSuggestionsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListSites invokes listSites operation.
//
// List all sites.
//
// GET /sites
func (c *Client) ListSites(ctx context.Context) ([]Site, error) {
	res, err := c.sendListSites(ctx)
	return res, err
}

func (c *Client) sendListSites(ctx context.Context) (res []Site, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listSites"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/sites"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListSites",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/sites"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListSitesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
