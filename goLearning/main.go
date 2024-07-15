package main

import (
	"belajar/database"
	"fmt"
	"belajar/controller/auth"
	"belajar/controller/course"
	"belajar/controller/student"
	"log"
	"net/http"	

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()	
	fmt.Println("Hello World")

	router := mux.NewRouter()

	router.HandleFunc("/regis", auth.Registration).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	// Router handler courses
	router.HandleFunc("/courses", course.GetCourse).Methods("GET")
	router.HandleFunc("/courses", auth.JWTAuth(course.PostCourse)).Methods("POST")
	router.HandleFunc("/courses/{id}", auth.JWTAuth(course.PutCourse)).Methods("PUT")
	router.HandleFunc("/courses/{id}", auth.JWTAuth(course.DeleteCourse)).Methods("DELETE")

	// Router handler students
	router.HandleFunc("/students", student.GetStudent).Methods("GET")
	router.HandleFunc("/students", auth.JWTAuth(student.PostStudent)).Methods("POST")
	router.HandleFunc("/students/{id}", auth.JWTAuth(student.PutStudent)).Methods("PUT")
	router.HandleFunc("/students/{id}", auth.JWTAuth(student.DeleteStudent)).Methods("DELETE")

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://127.0.0.1:5500"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
        Debug: true,
    })
	
    handler := c.Handler(router)
	
	fmt.Println("Server is running on http://localhost:6789")
	log.Fatal(http.ListenAndServe(":6789", handler))


}