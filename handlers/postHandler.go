package handlers

import (
	"Forum/database"
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	//Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	//Check request method in POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Retreieve UserID from context
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	//Extract form values
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")
	tags := r.FormValue("tags")

	//Post validation options
	if title == "" || content == "" {
		http.Error(w, "Title and Content are require", http.StatusBadRequest)
		return
	}

	//Create new post object
	post := &database.Post{
		AuthorID:     userID,
		Title:        title,
		Content:      content,
		Category:     category,
		Tags:         tags,
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		IsDeleted:    false,
	}

	//call function to save the post in database
	if err := database.CreatePost(post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	//Respond with created post
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"psot":    post,
	})
}

func CreatePostFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/create_post.html")
	if err != nil {
		http.Error(w, "Unable to render template:"+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	userRole, roleOk := r.Context().Value(UserRoleKey).(string)
	userID, idOk := r.Context().Value(UserIDKey).(int)

	//Check authentication
	if !roleOk || !idOk {
		http.Error(w, "Unauthorizes", http.StatusUnauthorized)
		return
	}

	//Retrieve postID from URL parameters
	postIDstr := r.URL.Query().Get("postID")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil || postID == 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	//Fetch post to identify its author
	post, err := database.GetPostByID(postID)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	//Authorization check
	if userRole == "administrator" || userRole == "moderator" || post.AuthorID == userID {
		//Delete post if authorized
		err = database.DeletePost(postID)
		if err != nil {
			http.Error(w, "Failed to delete post", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized to delete this post", http.StatusInternalServerError)
		return
	}

	//Redirect to appropriate page after deletion
	if userRole == "administrator" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else if userRole == "moderator" {
		http.Redirect(w, r, "/moderator", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	userID, idOk := r.Context().Value(UserIDKey).(int)
	userRole, _ := r.Context().Value(UserRoleKey).(string)

	//Check authentication
	if !idOk {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//Retrieve postID from URL parameters
	postIDstr := r.URL.Query().Get("postID")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil || postID == 0 {
		http.Error(w, "Invalid postID", http.StatusBadRequest)
		return
	}

	//Fetch post to identify its author
	post, err := database.GetPostByID(postID)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	//Only Authors can edit their posts
	if post.AuthorID != userID {
		http.Error(w, "Unauthorized to edit this post", http.StatusForbidden)
		return
	}

	//Render the edit for with the post data
	tmpl, err := template.ParseFiles("web/templates/edit_post.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string
		Post       *database.Post
		IsLoggedIn bool
		UserRole   string
	}{
		Title:      "Edit Post",
		Post:       post,
		IsLoggedIn: true,
		UserRole:   userRole,
	}

	if err := tmpl.ExecuteTemplate(w, "home_page.html", data); err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
