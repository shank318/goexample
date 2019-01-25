package models

import (
	"examplego/common"
	"github.com/jinzhu/gorm"
)

// Note data model
type URL struct {
	gorm.Model
	Url string
	CrawlTimeout int
	Frequency int
	FailureThreshold int
	Status string
}

// Serialize serializes note data
func (n URL) Serialize() common.JSON {
	return common.JSON{
		"id":         n.ID,
		"url":       n.Url,
		"crawl_timeout":   n.CrawlTimeout,
		"frequency": n.Frequency,
		"status": n.Status,
		"failure_threshold": n.FailureThreshold,
		"created_at": n.CreatedAt,
		"updated_at": n.UpdatedAt,

	}
}
