package usecase

import (
	"church_pesca_teams/domain"
	"church_pesca_teams/repository"
)

type PescaTeamService struct {
	repo repository.PescaTeamRepository
}

func NewPescaTeamService(repo repository.PescaTeamRepository) *PescaTeamService {
	return &PescaTeamService{repo: repo}
}

func (s *PescaTeamService) CreatePescaTeam(pt *domain.PescaTeam) error {
	return s.repo.Save(pt)
}

func (s *PescaTeamService) GetPescaTeamByID(id uint) (*domain.PescaTeam, error) {
	return s.repo.FindByID(id)
}

func (s *PescaTeamService) GetAllPescaTeams() ([]domain.PescaTeam, error) {
	return s.repo.FindAll()
}

func (s *PescaTeamService) UpdatePescaTeam(pt *domain.PescaTeam) error {
	return s.repo.Update(pt)
}

func (s *PescaTeamService) DeletePescaTeam(id uint) error {
	return s.repo.Delete(id)
}
