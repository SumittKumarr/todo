package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	"todo/helper"
	"todo/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var UserModel models.User
	err := json.NewDecoder(r.Body).Decode(&UserModel)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserModel.Password), 8)

	userDetails, err := helper.SignUp(UserModel.Name, string(hashedPassword))
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(userDetails)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, jsonErr = w.Write(jsonData)
	if jsonErr != nil {
		log.Printf("Error : %v", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	storedCredentials, err := helper.RetrieveCredentials(credentials.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
		w.WriteHeader(http.StatusInternalServerError)

	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCredentials.Password), []byte(credentials.Password)); err != nil {
		fmt.Printf("Error %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	// New session token created
	expiryAt := time.Now().Add(30 * time.Minute)

	SessionId, err := helper.CreateSession(storedCredentials.ID, expiryAt)
	if err != nil {
		log.Printf("Error : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(SessionId)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("Error : %v", wErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var UserModel models.User
	err := json.NewDecoder(r.Body).Decode(&UserModel)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.Context().Value("userId").(string)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserModel.Password), 8)

	err = helper.UpdateUser(UserModel.Name, string(hashedPassword), id)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("User Updated"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userId").(string)
	sessionId := r.Header.Get("sessionId")

	delErr := helper.DeleteAllTasks(id)
	if delErr != nil {
		log.Printf("Error %v", delErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	delErr = helper.DeleteSession(sessionId)
	if delErr != nil {
		log.Printf("Error %v", delErr)
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(delErr)
		return
	}
	delErr = helper.DeleteUser(id)
	if delErr != nil {
		log.Printf("Error %v", delErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := w.Write([]byte("Account Deleted"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.Task
	err := json.NewDecoder(r.Body).Decode(&taskModel)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.Context().Value("userId").(string)
	taskDetails, err := helper.CreateTask(taskModel.Name, id)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(taskDetails)
	if jsonErr != nil {
		log.Printf("Error %v", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var TaskModel models.Task
	jsonErr := json.NewDecoder(r.Body).Decode(&TaskModel)
	if jsonErr != nil {
		log.Printf("Error %v", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := helper.UpdateTask(TaskModel.ID)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("Task Updated"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

func FetchTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.Task
	jsonErr := json.NewDecoder(r.Body).Decode(&taskModel)
	if jsonErr != nil {
		log.Printf("Error %v", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskDetails, err := helper.FetchTask(taskModel.ID)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(taskDetails)
	if jsonErr != nil {
		log.Printf("Error %v", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.Task
	jsonErr := json.NewDecoder(r.Body).Decode(&taskModel)
	if jsonErr != nil {
		log.Printf("Error %v", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := helper.DeleteTask(taskModel.ID, taskModel.Name)
	if err != nil {
		log.Printf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Task Deleted"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	sessionId := r.Header.Get("sessionId")
	delErr := helper.DeleteSession(sessionId)
	if delErr != nil {
		log.Printf("Error %v", delErr)
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(delErr)
		return
	}

	_, err := w.Write([]byte("Logged out"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
