package store

import "time"

type User struct {
	ID           uint      `grom:"primaryKey" json:"id"`
	Username     string    `grom:"unique" json:"username"`
	Email        string    `grom:"unique" json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Post struct {
	ID        uint      `grom:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uint      `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `json:"post_id"`
	AuthorID  uint      `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
