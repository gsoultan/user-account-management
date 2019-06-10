package base

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID          *int64 `gorm:"primary_key;type:bigint;not null"`
	CreatedByID *int64 `gorm:"index;null"`
	CreatedAt   time.Time
	UpdateByID  *int64 `gorm:"index;null"`
	UpdatedAt   time.Time
	DeletedByID *int64 `gorm:"index;null"`
	DeletedAt   time.Time
}

func (this *Model) BeforeCreate(scope *gorm.Scope) error {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	err = scope.SetColumn("ID", node.Generate().Int64())
	return err
}
