package repository

import (
	"church_pesca_teams/domain"
)

type PescaTeamRepository interface {
	Save(pt *domain.PescaTeam) error
	FindByID(id uint) (*domain.PescaTeam, error)
	FindAll() ([]domain.PescaTeam, error)
	Update(pt *domain.PescaTeam) error
	Delete(id uint) error
}
