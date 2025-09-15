package main

import (
	"log"
	"os"

	"github.com/4ymane-code/mini-blog/internal/api"
	"github.com/4ymane-code/mini-blog/internal/auth"
	"github.com/4ymane-code/mini-blog/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file is not found")
	}

	db, err := store.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Migrate(db); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	postHandler := api.NewPostHandler(db)
	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePosts)
	r.GET("/secure/posts", auth.JWTMiddleware(), postHandler.GetPosts)
	r.POST("/posts", auth.JWTMiddleware(), postHandler.CreatePost)
	r.PUT("/posts/:id", auth.JWTMiddleware(), postHandler.UpdatePost)
	r.DELETE("/posts/:id", auth.JWTMiddleware(), postHandler.DeletePost)

	authHandler := api.NewAuthHandler(db)
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
	r.POST("/posts", auth.JWTMiddleware(), postHandler.CreatePost)
	r.PUT("/posts/:id", auth.JWTMiddleware(), postHandler.UpdatePost)
	r.DELETE("/posts/:id", auth.JWTMiddleware(), postHandler.DeletePost)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("server listening on port", port)
	r.Run(":" + port)
}
