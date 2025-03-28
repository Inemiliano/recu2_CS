package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	ID    primitive.ObjectID
	Edad int
	Nombre string
	Sexo *bool

}