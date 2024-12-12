package handlers

import (
	"forum/internal/config"
	db "forum/internal/database"
	"forum/internal/server/models"
	"forum/internal/server/templates"
	"log"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	var	session		*db.UserSession
	var	categories	[]*db.Category
	var	posts		[]*db.Post
	var clients 	[]db.Client
	var	userID		int
	var	err			error

	session, _ = r.Context().Value(config.SessionKey).(*db.UserSession)
	if session == nil {
		userID = 0
	} else {
		userID = session.UserID
	}

	posts, err = db.GetAllPosts(userID)
	if err != nil {
		log.Printf("Failed to retrieve posts: %v", err)
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	if categories, err = db.GetGlobalCategories(); err != nil {
		http.Error(
			w, "Error at fetching categories", http.StatusInternalServerError,
		)
	}

	clients, err = db.GetAllClients()
	if err != nil {
		log.Printf("Failed to retrieve clients: %v", err)
		http.Error(w, "Failed to retrieve clients", http.StatusInternalServerError)
		return
	}

	data := models.AdminPageData{
		Session: session,
		Posts: posts,
		Categories: categories,
		Clients: clients,
	}

	templates.RenderTemplate(w, "admin", data)
}

func AdminPostsHandler(w http.ResponseWriter, r*http.Request) {
	var	session		*db.UserSession
	var	posts		[]*db.Post
	var err 		error

	session, _ = r.Context().Value(config.SessionKey).(*db.UserSession)

	posts, err = db.GetAllPosts(1)
	if err != nil {
		log.Printf("Failed to retrieve posts: %v", err)
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	data := models.AdminPostsPageData{
		Session: session,
		Posts: posts,
	}

	templates.RenderTemplate(w, "admin-posts", data)
}

func AdminClientsHandler(w http.ResponseWriter, r*http.Request) {
	var	session		*db.UserSession
	var	clients		[]db.Client
	var err 		error

	session, _ = r.Context().Value(config.SessionKey).(*db.UserSession)

	clients, err = db.GetAllClients()
	if err != nil {
		log.Printf("Failed to retrieve clients: %v", err)
		http.Error(w, "Failed to retrieve clients", http.StatusInternalServerError)
		return
	}

	data := models.AdminClientsPageData{
		Session: session,
		Clients: clients,
	}

	templates.RenderTemplate(w, "admin-users", data)
}