package controllers

import (
	"recu/Usuarios/application"
	"recu/Usuarios/domain"
	"recu/Usuarios/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateUserHandler maneja la solicitud POST para crear un usuario
func CreateUserHandler(c *gin.Context) {
	log.Println("Método recibido: POST")

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error al decodificar JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	log.Printf("Usuario recibido: %+v", user)

	// Validar campos obligatorios
	if user.Nombre == "" || user.Edad == 0 || user.Sexo == nil {
		log.Println("Error: Nombre, edad y sexo son obligatorios")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre, edad y sexo son obligatorios"})
		return
	}

	// Crear una instancia del repositorio de MongoDB
	repo := infraestructure.NewUsuarioRepository()

	// Crear el caso de uso para la creación del usuario
	useCase := application.NewCreateUser(repo)

	// Ejecutar el caso de uso para guardar el usuario
	if err := useCase.Execute(&user); err != nil {
		log.Printf("Error al ejecutar caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario en la base de datos"})
		return
	}

	log.Println("Usuario insertado correctamente en MongoDB")
	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente"})
}

