package service

import (
	"BE-Project/model/domain"
	"BE-Project/model/web"
	"BE-Project/repository"

	"github.com/mashingan/smapping"
)

type ProjectService interface {
	All() []domain.Project
	Create(b web.ProjectCreateRequest) (domain.Project, error)
	FindById(id uint) (domain.Project, error)
	Update(b web.ProjectUpdateRequest) (domain.Project, error)
	Delete(id uint) error
}

type projectService struct {
	projectRepository repository.ProjectRepository
}

func NewUserService(projectRepository repository.ProjectRepository) ProjectService {
	return &projectService{projectRepository: projectRepository}
}

func (s *projectService) All() []domain.Project {
	return s.projectRepository.All()
}

func (s *projectService) Create(b web.ProjectCreateRequest) (domain.Project, error) {
	project := domain.Project{}
	err := smapping.FillStruct(&project, smapping.MapFields(&b))
	if err != nil {
		return domain.Project{}, err
	}
	return s.projectRepository.Create(project), nil
}

func (s *projectService) FindById(id uint) (domain.Project, error) {
	project, err := s.projectRepository.FindById(uint(id))
	if err != nil {
		return domain.Project{}, err
	}
	return project, nil
}

func (s *projectService) Update(b web.ProjectUpdateRequest) (domain.Project, error) {
	project := domain.Project{}
	_, err := s.projectRepository.FindById(uint(b.ID))
	if err != nil {
		return project, err
	}
	err = smapping.FillStruct(&project, smapping.MapFields(&b))
	if err != nil {
		return project, err
	}
	return s.projectRepository.Update(project), nil
}

func (s *projectService) Delete(id uint) error {
	project, err := s.projectRepository.FindById(uint(id))
	if err != nil {
		return err
	}
	s.projectRepository.Delete(project)
	return nil
}
