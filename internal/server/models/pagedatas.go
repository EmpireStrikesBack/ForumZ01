package models

import "forum/internal/database"

type ViewPostPageData struct {
	Session		*database.UserSession
	Post		*database.Post
}

type CreatePostPageData struct {
	Session	*database.UserSession
}

type HomePageData struct {
	Posts		[]database.Post
	Session		*database.UserSession
}
