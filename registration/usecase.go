package registration

import "registration/models"

type Usecase interface {
	NewClass(class *models.Class) error
	GetClassList() ([]models.Class, error)
	ImportStudentList(classID int, studentList []models.Student) error
	CheckIn(classID, employeeID int) (string, error)
	GetStudentList(classID int) ([]models.Student, error)
	DeleteClass(classID int) error
}
