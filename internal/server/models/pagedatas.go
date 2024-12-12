package models

import db "forum/internal/database"

type ViewPostPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
	Post		*db.Post
}

type LoginPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
}

type RegisterPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
}

type CreatePostPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
}

type HistoryPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
	Posts		[]*db.Post
}

type CategoryPostsPageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
	Category	*db.Category
	Posts		[]*db.Post
}

type HomePageData struct {
	Session		*db.UserSession // | null
	Categories	[]*db.Category
	Posts		[]*db.Post
}

type AdminPageData struct {
	Session		*db.UserSession
	Posts		[]*db.Post
	Categories	[]*db.Category
	Clients		[]db.Client
	Reports		[]db.Report
}

type AdminPostsPageData struct {
	Session		*db.UserSession
	Posts		[]*db.Post
}

type AdminClientsPageData struct {
	Session		*db.UserSession
	Clients		[]db.Client
}