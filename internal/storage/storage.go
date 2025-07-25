package storage

import "github.com/SSKSamarSinghKanojiya/student-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int)(int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student,error)
}
