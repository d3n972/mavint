package dao

import (
	"github.com/d3n972/mavint/domain/models"
	"github.com/d3n972/mavint/domain/repository"
	"gorm.io/gorm"
	"time"
)

type VPEEntryDAO struct {
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

func (V VPEEntryDAO) FromEntity(e models.VPEEntry) repository.DAO[models.VPEEntry] {
	return VPEEntryDAO{
		VPEHash:   e.VPEHash,
		From:      e.From,
		Until:     e.Until,
		Provider:  e.Provider,
		VPEID:     e.VPEID,
		EntryType: e.EntryType,
		VPEName:   e.VPEName,
		Status:    e.Status,
	}
}

func (V VPEEntryDAO) ToEntity() models.VPEEntry {
	//TODO implement me
	panic("implement me")
}
