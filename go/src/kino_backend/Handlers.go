package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint64, 0),
		usersAuth: map[string]*User{
			"testUser": {1, "testuser", "test"},
		},
		users: make([]User, 0),
		mu: &sync.Mutex{},
	}
}

func(api *MyHandler) Signup(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserIn := new(UserInput)
	//err = json.Unmarshal(bytes, newUserInput)
	err := decoder.Decode(newUserIn)

	if err != nil{
		log.Printf("error unmarshalling %s", err)
		w.Write([]byte("{}"))
		return
	}

	fmt.Println(newUserIn)
	api.mu.Lock()

	var id uint64 = 0
	if len(api.users) > 0{
		id = uint64(api.users[len(api.users)-1].ID + 1)
	}

	api.users = append(api.users, User{
		ID: id,
		Username: newUserIn.Username,
		Password: newUserIn.Password,
	})
	api.mu.Unlock()
}

func (api *MyHandler) ListUsers(w http.ResponseWriter, r *http.Request){
	encoder := json.NewEncoder(w)
	api.mu.Lock()
	err := encoder.Encode(api.users)
	//usersJson, err := json.Marshal(h.users)
	api.mu.Unlock()

	if err != nil {
		log.Printf("error marshal json %s", err)
		w.Write([]byte("{}"))
	}
}

func (api *MyHandler) MyProfile(w http.ResponseWriter, r *http.Request){
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		encoder := json.NewEncoder(w)
		w.Write([]byte("your profile"))
		api.mu.Lock()
		err := encoder.Encode(api.users[api.sessions[session.Value]])
		api.mu.Unlock()
		if err != nil {
			log.Printf("error marshal json %s", err)
		}
	} else {
		w.Write([]byte("not autrorized"))
	}
}

func (api *MyHandler) EditUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("autrorized"))
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	changedUser := new(UserEdit)
	err := decoder.Decode(changedUser)

	if err != nil{
		log.Printf("error unmarshalling %s", err)
		w.Write([]byte("{}"))
		return
	}

	ok := false
	var userID uint64

	for _, value := range api.users{
		if value.Username == changedUser.PrevUsername{
			if value.Password == changedUser.PrevPassword{
				if changedUser.NewPassword != "" || changedUser.NewPassword != ""{
					ok = true
					userID = value.ID
				}

			}
		}
	}

	if !ok {
		http.Error(w, `invalid login or password`, 404)
		return
	}
	api.users[userID].Username = changedUser.NewUsername
	api.users[userID].Password = changedUser.NewPassword
}

func (api *MyHandler) Login(w http.ResponseWriter, r *http.Request) {

	ok := false
	var user User
	for _, value := range api.users{
		if value.Username == r.FormValue("login"){
			if value.Password == r.FormValue("password"){
				ok = true
				user = value
			}
		}
	}
	if !ok {
		http.Error(w, `invalid login or password`, 404)
		return
	}

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
		HttpOnly:true,
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(SID))

}

func (api *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no sess`, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, `no sess`, 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (api *MyHandler) Root(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
