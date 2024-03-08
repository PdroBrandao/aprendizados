package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) save() {
	storage = append(storage, *u)
	logger.Println("Storage: ", storage)
}

func (u *User) Search(id int) (User, error) {
	for _, user := range storage {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

var storage []User

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

func main() {
	r := httprouter.New()
	r.GET("/:id", getUser)
	r.POST("/", newUser)
	logger.Println("Listening on port 9999")
	http.ListenAndServe(":9999", r)
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := readAndValidateRequest(ps.ByName("id"))
	if err != nil {
		errMsg := "Error reading request: " + err.Error()
		writeJSON(w, errMsg, http.StatusBadRequest)
		return
	}
	var user User
	data, err := user.Search(id)
	if err != nil {
		writeJSON(w, err, http.StatusNotFound)
		return
	}
	logger.Println("User found: ", data)
	writeJSON(w, data, http.StatusOK)
}

func readAndValidateRequest(id string) (int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	if idInt < 0 {
		return 0, errors.New("Invalid user ID")
	}
	return idInt, nil
}

func newUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := processAndValidateRequestBody(r.Body)
	if err != nil {
		errMsg := "Error processing request body: " + err.Error()
		writeJSON(w, errMsg, http.StatusBadRequest)
		return
	}
	user.save()
	writeJSON(w, user, http.StatusCreated)
}

func processAndValidateRequestBody(body io.Reader) (User, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return User{}, err
	}
	if user.ID < 0 {
		return User{}, errors.New("Invalid user ID")
	}
	return user, nil
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
