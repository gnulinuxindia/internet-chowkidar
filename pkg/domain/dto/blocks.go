package dto

import "time"

type BlockDto struct {
	IspID          int       `json:"isp_id"`
	SiteID         int       `json:"site_id"`
	ClientID	int       `json:"client_id"`
	LastReportedAt time.Time `json:"last_reported_at"`
	IsBlocked      bool      `json:"is_blocked"`
}
