package router

import (
    "tudoo-backend/middleware"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/board/{id}", middleware.GetBoard).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/board", middleware.GetAllBoard).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newboard", middleware.CreateBoard).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/board/{id}", middleware.UpdateBoard).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteboard/{id}", middleware.DeleteBoard).Methods("DELETE", "OPTIONS")

    return router
}