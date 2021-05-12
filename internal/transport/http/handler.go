package http

import (
	"encoding/json"
	"fmt"
	"github.com/valdirmendesdev/go-rest-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to the comments service
type Handler struct {
	Router     *mux.Router
	Repository *services.CommentRepository
}

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// NewHandler - returns a pointer to a Handler
func NewHandler(repository *services.CommentRepository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

// SetupRoutes - sets up all the routes for the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comments", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := Response{
			Message: "I'm alive",
		}
		json, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(json)
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}

	comment, err := h.Repository.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving Comment By ID", err)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		sendErrorResponse(w, "Error enconding Comment json", err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Repository.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all comments", err)
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		sendErrorResponse(w, "Error enconding Comment json", err)
	}
}
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment services.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
	}

	comment, err := h.Repository.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		sendErrorResponse(w, "Error enconding Comment json", err)
	}

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	var comment services.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
	}
	comment, err = h.Repository.UpdateComment(uint(i), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update new comment", err)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		sendErrorResponse(w, "Error enconding Comment json", err)
	}
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}

	err = h.Repository.DeleteComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Failed to delete Comment By ID", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"});  err != nil{
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()});  err != nil{
		panic(err)
	}
}
