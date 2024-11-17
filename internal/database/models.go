package database

import (
	"time"
)

type Client struct {
	UserID       int        `json:"user_id"`
	LastName     string     `json:"last_name"`
	FirstName    string     `json:"first_name"`
	UserName     string     `json:"user_name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Avatar       string     `json:"avatar"`
	BirthDate    time.Time  `json:"birth_date"`
	UserRole     string     `json:"user_role"`
	CreationDate time.Time  `json:"creation_date"`
	UpdateDate   time.Time  `json:"update_date"`
	DeletionDate *time.Time `json:"deletion_date"`
}

type UserSession struct {
	SessionID    string     `json:"session_id"`
	UserID       int        `json:"user_id"`
	UserRole     string     `json:"user_role"`
	UserName     string     `json:"user_name"`
	IsLoggedIn   bool       `json:"is_logged_in"`
	Expiration   time.Time  `json:"expiration"`
	CreationDate time.Time  `json:"creation_date"`
	UpdateDate   time.Time  `json:"update_date"`
	DeletionDate *time.Time `json:"deletion_date"`
	IsDeleted    bool       `json:"is_deleted"`
}

type Post struct {
	PostID       int        `json:"post_id"`
	AuthorID     int        `json:"author_id"`
	UserName     string     `json:"user_name"`
	Title        string     `json:"title"`
	Category     string     `json:"category"`
	Tags         string     `json:"tags"`
	Content      string     `json:"content"`
	CreationDate time.Time  `json:"creation_date"`
	UpdateDate   time.Time  `json:"update_date"`
	DeletionDate *time.Time `json:"deletion_date"`
	IsDeleted    bool       `json:"is_deleted"`
	Comments     []Comment
}

type Comment struct {
	CommentID    int       `json:"comment_id"`
	PostID       int       `json:"post_id"`
	UserID       int       `json:"user_id"`
	UserName     string    `json:"user_name"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
}