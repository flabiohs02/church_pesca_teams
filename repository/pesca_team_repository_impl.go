package repository

import (
	"church_pesca_teams/domain"

	"gorm.io/gorm"
)

type GormPescaTeamRepository struct {
	db *gorm.DB
}

func NewGormPescaTeamRepository(db *gorm.DB) *GormPescaTeamRepository {
	return &GormPescaTeamRepository{db: db}
}

func (r *GormPescaTeamRepository) Save(pt *domain.PescaTeam) error {
	return r.db.Create(pt).Error
}

func (r *GormPescaTeamRepository) FindByID(id uint) (*domain.PescaTeam, error) {
	var pt domain.PescaTeam
	err := r.db.First(&pt, id).Error
	return &pt, err
}

func (r *GormPescaTeamRepository) FindAll() ([]domain.PescaTeam, error) {
	var pts []domain.PescaTeam
	err := r.db.Find(&pts).Error
	return pts, err
}

func (r *GormPescaTeamRepository) Update(pt *domain.PescaTeam) error {
	return r.db.Save(pt).Error
}

func (r *GormPescaTeamRepository) Delete(id uint) error {
	return r.db.Delete(&domain.PescaTeam{}, id).Error
}
