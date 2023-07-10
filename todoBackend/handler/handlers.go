package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todolist/models"
	"todolist/repo"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetTaskList(w http.ResponseWriter, r *http.Request) {
	taskList, err := repo.GetAllTask()
	if len(taskList) == 0 && err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if len(taskList) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	respBody, err := json.Marshal(taskList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.ToDoList

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || reqBody == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	generatedId := uuid.New().String()
	task.ID = generatedId

	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repo.CreateTaskInDB(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := "Created Task With Id: " + generatedId
	w.Write([]byte(resp))
}

func MarkTaskAsCompleted(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := repo.SetTaskAsCompleted(id)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := repo.SetTaskAsNotCompleted(id)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := repo.DeleteSingleTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := "Deleted task " + id
	w.Write([]byte(resp))
}

func SetupHandlers() http.Handler {
	router := mux.NewRouter()

	// Enable CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3001"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Attach the handlers
	router.HandleFunc("/api/tasks", GetTaskList).Methods("GET")
	router.HandleFunc("/api/tasks", AddTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}/completed", MarkTaskAsCompleted).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}/undo", UndoTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE")

	// Wrap the router with the CORS middleware
	handler := cors(router)

	return handler
}
