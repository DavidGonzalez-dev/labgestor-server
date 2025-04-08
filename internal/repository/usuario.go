package repository

import (
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los usuarios en la base datos
type UsuarioRepository interface {
	ObtenerUsuarioID(id string) (*models.Usuario, error)
	CrearUsuario(usuario *models.Usuario) error
	ActualizarUsuario(usuario *models.Usuario) error
	ObtenerUsuarios() (*[]models.Usuario, error)
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
	if err := repo.DB.Create(&usuario).Error; err != nil {
		return err
	}
	return nil

}

func (repo *usuarioRepository) ActualizarUsuario(usuario *models.Usuario) error {
	if err := repo.DB.Save(usuario).Error; err != nil {
		return err
	}
	return nil
}

func (repo *usuarioRepository) ObtenerUsuarios() (*[]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := repo.DB.Preload("Rol").Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return &usuarios, nil

}
