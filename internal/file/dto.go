package file

import "time"

type FileResponse struct {
	ID        uint      `json:"id"`
	Filename  string    `json:"filename"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}
