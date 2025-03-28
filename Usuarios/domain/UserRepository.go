// domain/user_repository.go
package domain


type UserRepository interface {
	Save(user *User) error
	GetAll() ([]User, error)
	CountByGender() (map[string]int, error)
}
