package api

import (
	"log/slog"
	"net/http"

	"github.com/4ymane-code/mini-blog/internal/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	DB *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{DB: db}
}

func (h *PostHandler) CreatePosts(ctx *gin.Context) {
	var input store.Post
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to fetch posts"})
		return
	}

	if err := h.DB.Create(&input); err != nil {
		slog.Error("faild to create post", "error:", err)
		return
	}

	ctx.JSON(http.StatusCreated, input)
}

func (h *PostHandler) GetPosts(ctx *gin.Context) {
	var posts []store.Post
	if err := h.DB.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch post"})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
