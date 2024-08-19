package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dileepkushwaha/sre-bootcamp/config"
	"github.com/dileepkushwaha/sre-bootcamp/internal/controller"
	"github.com/dileepkushwaha/sre-bootcamp/internal/model"
	"github.com/dileepkushwaha/sre-bootcamp/internal/repository"
	"github.com/dileepkushwaha/sre-bootcamp/internal/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupRouter() *mux.Router {
	cfg := config.LoadConfig()
	db, _ := gorm.Open("postgres", cfg.DB_URL)
	db.AutoMigrate(&model.Student{})
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentController := controller.NewStudentController(studentService)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/students", studentController.CreateStudent).Methods("POST")
	return r
}

func TestCreateStudent(t *testing.T) {
	r := setupRouter()

	student := model.Student{Name: "John Doe", Age: 25, Email: "john@example.com"}
	body, _ := json.Marshal(student)
	req, _ := http.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestHealthCheck(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
