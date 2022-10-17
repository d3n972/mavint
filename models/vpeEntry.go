package models

import (
	"gorm.io/gorm"
	"time"
)

type VPEEntry struct {
	gorm.Model
	VPEHash   string `db:"hash"`
	From      time.Time
	Until     time.Time
	Provider  string
	VPEID     string
	EntryType string
	VPEName   string
	Status    string
}
