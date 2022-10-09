package services

import (
	"context"
	"fmt"
	"github.com/d3n972/mavint/models/db"
	"gorm.io/gorm"
	"time"
)

type EngineDiscovery struct {
}

func (e *EngineDiscovery) FindByTrainNumber(ctx context.Context, tn string) (*db.EngineWorkday, error) {
	res := db.EngineWorkday{}
	fmt.Printf("[TN] %s\n", tn)
	if db := ctx.Value("db"); db != nil {
		tx := db.(*gorm.DB).
			Model(&res).
			Find(&res, "date=? AND train_number=?",
				time.Now().Format("2006-01-02"),
				tn)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	return &res, nil
}
