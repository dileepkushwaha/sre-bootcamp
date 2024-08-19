package repository

import (
	"github.com/dileepkushwaha/sre-bootcamp/internal/model"
	"github.com/jinzhu/gorm"
)

type StudentRepository interface {
	Create(student *model.Student) error
	FindAll() ([]model.Student, error)
	FindByID(id uint) (*model.Student, error)
	Update(student *model.Student) error
	Delete(student *model.Student) error
}

type studentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{DB: db}
}

func (r *studentRepository) Create(student *model.Student) error {
	return r.DB.Create(student).Error
}

func (r *studentRepository) FindAll() ([]model.Student, error) {
	var students []model.Student
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByID(id uint) (*model.Student, error) {
	var student model.Student
	err := r.DB.First(&student, id).Error
	return &student, err
}

func (r *studentRepository) Update(student *model.Student) error {
	return r.DB.Save(student).Error
}

func (r *studentRepository) Delete(student *model.Student) error {
	return r.DB.Delete(student).Error
}
