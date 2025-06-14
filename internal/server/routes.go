package server

import (
	"fmt"
	"net/http"

	"github.com/FilipeRSantos/planning-poker/cmd/web"
	"github.com/FilipeRSantos/planning-poker/internal/server/ws"
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	hub := ws.NewHub()
	go hub.Run()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(web.StartSession()).ServeHTTP(w, r)
	})

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	r.HandleFunc("/api/room", func(w http.ResponseWriter, r *http.Request) {
		uid := uuid.New().String()

		w.Header().Add("hx-Location", fmt.Sprintf("/room/%s", uid))

		w.WriteHeader(201)
	}).Methods("POST")

	r.HandleFunc("/room/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		templ.Handler(web.RoomById(id)).ServeHTTP(w, r)
	}).Methods("GET")

	return r
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
