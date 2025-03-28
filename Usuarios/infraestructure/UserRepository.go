// infraestructure/usuario_repository.go
package infraestructure

import (
	"recu/Usuarios/domain"
	"recu/core"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
	collection *mongo.Collection
}

// NewUsuarioRepository crea una nueva instancia del repositorio MongoDB para usuarios
func NewUsuarioRepository() *UsuarioRepository {
	client := core.GetMongoClient()
	collection := client.Database("api_hexa").Collection("Usuarios")
	return &UsuarioRepository{collection: collection}
}

// Save guarda un usuario en la base de datos
func (r *UsuarioRepository) Save(u *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insertar el usuario en la colección de MongoDB
	res, err := r.collection.InsertOne(ctx, u)
	if err != nil {
		log.Printf("Error al insertar el usuario: %v", err)
		return err
	}

	log.Printf("Usuario insertado con éxito con ID: %v", res.InsertedID)
	return nil
}

// GetAll obtiene todos los usuarios de la base de datos
func (r *UsuarioRepository) GetAll() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al obtener usuarios: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var usuarios []domain.User
	for cursor.Next(ctx) {
		var usuario domain.User
		if err := cursor.Decode(&usuario); err != nil {
			log.Printf("Error al decodificar un usuario: %v", err)
			continue
		}
		usuarios = append(usuarios, usuario)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error al recorrer los usuarios: %v", err)
		return nil, err
	}

	return usuarios, nil
}

// CountByGender cuenta cuántos hombres y cuántas mujeres hay en la base de datos
func (r *UsuarioRepository) CountByGender() (map[string]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Contar los hombres
	countHombres, err := r.collection.CountDocuments(ctx, bson.M{"genero": "hombre"})
	if err != nil {
		log.Printf("Error al contar hombres: %v", err)
		return nil, err
	}

	// Contar las mujeres
	countMujeres, err := r.collection.CountDocuments(ctx, bson.M{"genero": "mujer"})
	if err != nil {
		log.Printf("Error al contar mujeres: %v", err)
		return nil, err
	}

	return map[string]int{
		"hombres": int(countHombres),
		"mujeres": int(countMujeres),
	}, nil
}
