package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los usuarios en la base datos
type UsuarioRepository interface {
	ObtenerUsuarioID(id string) (*models.Usuario, error)
	CrearUsuario(usuario *models.Usuario) error
}

// Structura que implementa la interfaz anteriormente definida
type usuarioRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura usuarioRepository
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{DB: db}
}

// ---------------------------
// Metodos de la estructura
// ---------------------------

func (repo *usuarioRepository) ObtenerUsuarioID(id string) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := repo.DB.Preload("Rol").First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (repo *usuarioRepository) CrearUsuario(usuario *models.Usuario) error {
	return repo.DB.Create(&usuario).Error
}


