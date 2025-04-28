package file

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fhva29/go-vault/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Upload(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		utils.SendError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	// Tem que fazer dessa forma pois o middleware/token pode ta salvando
	// em qualquer formato
	userID, _ := strconv.ParseUint(fmt.Sprintf("%v", userIDInterface), 10, 32)

	file, err := ctx.FormFile("file")
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "No file uploaded", utils.FormatValidationError(err))
		return
	}

	uploadsPath := "./uploads"
	os.MkdirAll(uploadsPath, os.ModePerm)

	dst := filepath.Join(uploadsPath, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Could not save file", utils.FormatValidationError(err))
		return
	}

	f := &File{
		UserID:   uint(userID),
		Filename: file.Filename,
		Path:     dst,
		Size:     file.Size,
	}

	if err := h.service.Save(f); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Could not save file metadata", utils.FormatValidationError(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
	})
}

func (h *Handler) ListMyFiles(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		utils.SendError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID, _ := strconv.ParseUint(fmt.Sprintf("%v", userIDInterface), 10, 32)

	files, err := h.service.ListByUser(uint(userID))
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Could not fetch files", utils.FormatValidationError(err))
		return
	}

	// Mapeia para FileResponse
	var filesResponse []FileResponse
	for _, f := range files {
		filesResponse = append(filesResponse, FileResponse{
			ID:        f.ID,
			Filename:  f.Filename,
			Size:      f.Size,
			CreatedAt: f.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"files": filesResponse})
}
