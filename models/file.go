package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID uuid.UUID	`gorm:"type:uuid;primaryKey;" json:"id"`
	Name string	`json:"name"` 
	Size int64 `json:"size"`
	Path string `json:"path"`
	CreatedAt time.Time  `json:"created_at"`
}

func (f *File) GenerateUUID(tx *gorm.DB) (err error){
	f.ID = uuid.New()
	return nil
}