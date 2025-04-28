package file

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
}
