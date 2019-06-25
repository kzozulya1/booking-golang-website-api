package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	roomsActions "app/pkg/controller/rooms"
	reservationsActions "app/pkg/controller/reservations"
	usersActions "app/pkg/controller/users"
	"app/pkg/middleware"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/rooms/", roomsActions.GetAll).Methods("GET","OPTIONS")
	router.HandleFunc("/api/v1/rooms/", roomsActions.Create).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/rooms/{id}", roomsActions.GetOne).Methods("GET","OPTIONS")
	router.HandleFunc("/api/v1/rooms/{id}", roomsActions.Update).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/v1/rooms/{id}", roomsActions.Delete).Methods("DELETE","OPTIONS")

	router.HandleFunc("/api/v1/reservations/room/{id}",reservationsActions.Create).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/reservations/{id}", reservationsActions.Delete).Methods("DELETE","OPTIONS")

	router.HandleFunc("/api/v1/users/", usersActions.Create).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/users/login", usersActions.Login).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/users/me", usersActions.MeGet).Methods("GET","OPTIONS")
	router.HandleFunc("/api/v1/users/me", usersActions.MePut).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/v1/users/{id}", usersActions.GetOne).Methods("GET","OPTIONS")

	router.Use(middleware.Cors) //attach CORS allow headers
	router.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("APP_API_PORT")
	fmt.Println("Listen ", port)
	err := http.ListenAndServe(":"+port, router) //Launch the app
	if err != nil {
		fmt.Print(err)
	}
}
