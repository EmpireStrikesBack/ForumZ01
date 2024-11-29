package database

import (
	"time"
)

type Client struct {
	ID				int        `json:"id"`
	LastName		string     `json:"last_name"`
	FirstName		string     `json:"first_name"`
	UserName		string     `json:"user_name"`
	Email			string     `json:"email"`
	OauthProvider	string     `json:"oauth_provider"`
	OauthID			string     `json:"oauth_id"`
	Password		string     `json:"password"`
	Avatar			string     `json:"avatar"`
	BirthDate		time.Time  `json:"birth_date"`
	UserRole		string     `json:"user_role"`
	CreationDate	time.Time  `json:"creation_date"`
	UpdateDate		time.Time  `json:"update_date"`
	DeletionDate	*time.Time `json:"deletion_date"`
}

type UserSession struct {
	ID				string     `json:"id"`
	UserID			int        `json:"user_id"`
	UserRole		string     `json:"user_role"`
	UserName		string     `json:"user_name"`
	Expiration		time.Time  `json:"expiration"`
	CreationDate	time.Time  `json:"creation_date"`
	UpdateDate		time.Time  `json:"update_date"`
	DeletionDate	*time.Time `json:"deletion_date"`
	IsDeleted		bool       `json:"is_deleted"`
}

type Category struct {
	ID		int
	Name	string
}

type Post struct {
	ID				int			`json:"id"`
	AuthorID		int			`json:"author_id"`
	UserName		string		`json:"user_name"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`
	CreationDate	time.Time	`json:"creation_date"`
	UpdateDate		time.Time	`json:"update_date"`
	DeletionDate	*time.Time	`json:"deletion_date"`
	IsDeleted		bool		`json:"is_deleted"`
	Likes			int
	Dislikes		int
	Categories		[]PostCategory
	Comments		[]Comment	`json:"comments"`
}

type PostCategory struct {
	ID		int
	Name	string
	PostID	int
}

type Comment struct {
	ID				int       `json:"id"`
	PostID			int       `json:"post_id"`
	UserID			int       `json:"user_id"`
	UserName		string    `json:"user_name"`
	Content			string    `json:"content"`
	CreationDate	time.Time `json:"creation_date"`
}

type LikeDislike struct {
	ID			int
	PostID		int
	UserID		int
	Liked		bool
	UpdateDate	time.Time
}
