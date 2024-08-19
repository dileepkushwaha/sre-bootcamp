package service

import (
	"github.com/dileepkushwaha/sre-bootcamp/internal/model"
	"github.com/dileepkushwaha/sre-bootcamp/internal/repository"
)

type StudentService interface {
	CreateStudent(student *model.Student) error
	GetAllStudents() ([]model.Student, error)
	GetStudentByID(id uint) (*model.Student, error)
	UpdateStudent(student *model.Student) error
	DeleteStudent(id uint) error
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(student *model.Student) error {
	return s.repo.Create(student)
}

func (s *studentService) GetAllStudents() ([]model.Student, error) {
	return s.repo.FindAll()
}

func (s *studentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.repo.FindByID(id)
}

func (s *studentService) UpdateStudent(student *model.Student) error {
	return s.repo.Update(student)
}

func (s *studentService) DeleteStudent(id uint) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(student)
}
