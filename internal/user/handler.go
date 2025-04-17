package user

import (
	"net/http"

	"github.com/fhva29/go-vault/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Signup(ctx *gin.Context) {
	var input SignupInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.SendError(ctx, 400, "Invalid input", utils.FormatValidationError(err))
		return
	}

	if err := h.service.Signup(input); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created succesfully",
	})
}
