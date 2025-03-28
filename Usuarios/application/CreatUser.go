package application

import (
	"recu/Usuarios/domain"
)

// CrearUsuario es el caso de uso para crear un usuario
type CrearUsuario struct {
	repo domain.UserRepository
}

// NewCreateUser crea un nuevo caso de uso para crear un usuario
func NewCreateUser(repo domain.UserRepository) *CrearUsuario {
	return &CrearUsuario{repo: repo}
}

// Execute ejecuta la creación del usuario en el repositorio
func (uc *CrearUsuario) Execute(user *domain.User) error {
	return uc.repo.Save(user)
}
