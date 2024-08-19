package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/dileepkushwaha/sre-bootcamp/config"
	"github.com/dileepkushwaha/sre-bootcamp/internal/controller"
	"github.com/dileepkushwaha/sre-bootcamp/internal/model"
	"github.com/dileepkushwaha/sre-bootcamp/internal/repository"
	"github.com/dileepkushwaha/sre-bootcamp/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	log := logrus.New()

	db, err := gorm.Open("postgres", cfg.DB_URL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Student{})

	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentController := controller.NewStudentController(studentService)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/students", studentController.CreateStudent).Methods("POST")
	r.HandleFunc("/api/v1/students", studentController.GetAllStudents).Methods("GET")
	r.HandleFunc("/api/v1/students/{id}", studentController.GetStudentByID).Methods("GET")
	r.HandleFunc("/api/v1/students/{id}", studentController.UpdateStudent).Methods("PUT")
	r.HandleFunc("/api/v1/students/{id}", studentController.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/healthcheck", studentController.HealthCheck).Methods("GET")

	log.Info("Starting the server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
