package application

import "recu/Usuarios/domain"


type ObtenerUsuarios struct {
	repo domain.UserRepository
}


func NewObtenerUsuarios(repo domain.UserRepository) *ObtenerUsuarios {
	return &ObtenerUsuarios{repo: repo}
}


func (uc *ObtenerUsuarios) Ejecutar() ([]domain.User, error) {
	return uc.repo.GetAll()
}
