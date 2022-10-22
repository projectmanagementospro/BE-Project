package repository

import (
	"BE-Project/model/domain"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	All() []domain.Project
	Create(r domain.Project) domain.Project
	Update(r domain.Project) domain.Project
	Delete(r domain.Project)
	FindById(id uint) (domain.Project, error)
}

type ProjectConnection struct {
	connection *gorm.DB
}

func NewProjectRepository(dbConn *gorm.DB) ProjectRepository {
	return &ProjectConnection{
		connection: dbConn,
	}
}

func (db *ProjectConnection) All() []domain.Project {
	var projects []domain.Project
	db.connection.Find(&projects)
	return projects
}

func (db *ProjectConnection) Create(r domain.Project) domain.Project {
	db.connection.Create(&r)
	return r
}

func (db *ProjectConnection) Update(r domain.Project) domain.Project {
	db.connection.Omit("created_at").Save(&r)
	return r
}

func (db *ProjectConnection) Delete(r domain.Project) {
	db.connection.Delete(&r)
}

func (db *ProjectConnection) FindById(id uint) (domain.Project, error) {
	var project domain.Project
	err := db.connection.First(&project, id).Error
	return project, err
}
