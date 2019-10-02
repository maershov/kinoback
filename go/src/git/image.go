package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func (api *MyHandler) mainPage(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		w.Write(uploadFormTmpl)
	} else {
		w.Write([]byte("not autrorized"))
	}
}

func (api *MyHandler) uploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
	session, err := r.Cookie("session_id")
	id := api.sessions[session.Value]
	strid := strconv.Itoa(int(id))
	error := Download(file, strid)
	if error != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (api *MyHandler) GetPhoto(w http.ResponseWriter, r *http.Request){


	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		id := api.sessions[session.Value]
		file, err := getPhoto(int(id))
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reader := bufio.NewReader(&file)
		bytes := make([]byte, 10<<20)
		_, err = reader.Read(bytes)

		w.Header().Set("content-type", "multipart/form-data;boundary=1")

		_, err = w.Write(bytes)
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(500)
			return
		}

		log.Println("Successfully Uploaded File")

	} else {
		w.Write([]byte("not autrorized"))
	}
}
