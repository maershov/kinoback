package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	api := NewMyHandler()
	r.HandleFunc("/", api.Root)
	r.HandleFunc("/signup", api.Signup)
	r.HandleFunc("/list", api.ListUsers)
	r.HandleFunc("/login", api.Login)
	r.HandleFunc("/logout", api.Logout)
	r.HandleFunc("/me", api.MyProfile)
	r.HandleFunc("/edit", api.EditUser)
	r.HandleFunc("/photodownload", api.GetPhoto)
	r.HandleFunc("/photo", api.mainPage)
	r.HandleFunc("/upload", api.uploadPage)

	http.ListenAndServe(":8080", r)
}