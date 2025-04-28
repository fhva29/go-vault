package file

import "gorm.io/gorm"

type Repository interface {
	Save(file *File) error
	ListByUser(userID uint) ([]File, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&File{})
	return &repository{db}
}

func (r *repository) Save(file *File) error {
	return r.db.Create(file).Error
}

func (r *repository) ListByUser(userID uint) ([]File, error) {
	var files []File
	err := r.db.Where("user_id = ?", userID).Find(&files).Error
	return files, err
}
