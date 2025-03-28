package application

import "recu/Usuarios/domain"


type ContarUsuariosPorGenero struct {
	repo domain.UserRepository
}


func NewContarUsuariosPorGenero(repo domain.UserRepository) *ContarUsuariosPorGenero {
	return &ContarUsuariosPorGenero{repo: repo}
}


func (uc *ContarUsuariosPorGenero) Ejecutar() (map[string]int, error) {
	return uc.repo.CountByGender()
}
