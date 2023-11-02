package http

import (
	"net/http"

	"github.com/Yury132/Golang-Task-3/internal/transport/http/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Home).Methods(http.MethodGet)
	r.HandleFunc("/auth", h.Auth).Methods(http.MethodGet)
	r.HandleFunc("/callback", h.Callback).Methods(http.MethodGet)
	r.HandleFunc("/me", h.Me).Methods(http.MethodGet)
	r.HandleFunc("/logout", h.Logout).Methods(http.MethodGet)
	r.HandleFunc("/users-list", h.GetUsersList).Methods(http.MethodGet)
	//--
	//--
	//--
	// Новое
	r.HandleFunc("/get-chats", h.GetChat).Methods(http.MethodGet)
	r.HandleFunc("/create-chat", h.CreateChat).Methods(http.MethodGet)
	r.HandleFunc("/edit-chat", h.EditChat).Methods(http.MethodGet)
	r.HandleFunc("/message", h.Message).Methods(http.MethodGet)

	return r
}
