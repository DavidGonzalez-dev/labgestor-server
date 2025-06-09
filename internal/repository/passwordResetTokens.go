package repository

import (
	"errors"
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Definimos la interfaz
type PasswordResetTokenRepository interface {
	Create(token *models.PasswordResetToken) error
	MarkAsUsed(tokenID int) error
	DeleteByUserID(userID string) error
	GetMostRecentTokenByUserID(userID string) (*models.PasswordResetToken ,error)
}

type passwordResetTokenRepo struct {
	DB *gorm.DB
}

// Funcion para instanciar el repositorio
func NewPasswordResetTokenRepo(db *gorm.DB) PasswordResetTokenRepository {
	return &passwordResetTokenRepo{DB: db}
}

// Funcion para crear un token en la base de datos
func (repo *passwordResetTokenRepo) Create(token *models.PasswordResetToken) error {
	return repo.DB.Create(&token).Error
}

// Funcion para marcar un token como usado basado en el id del token
func (repo *passwordResetTokenRepo) MarkAsUsed(tokenID int) error {
	// Obtenemos el registro del token desde la base de datos
	var token models.PasswordResetToken
	if err := repo.DB.First(&token, tokenID).Error; err != nil {
		return err
	}

	// Lo marcamos como usado
	token.Used = true

	// Guardamos los cambios
	if err := repo.DB.Save(token).Error; err != nil {
		return err
	}

	return nil
}

// Funcion para eliminar los tokens de un usario en especifico
func (repo *passwordResetTokenRepo) DeleteByUserID(userID string) error {
	return repo.DB.Exec("DELETE FROM password_reset_tokens WHERE id_usuario=?", userID).Error
}

// Funcion para obtener los tokens de un usuario
func (repo *passwordResetTokenRepo) GetMostRecentTokenByUserID(userID string) (*models.PasswordResetToken ,error) {
	var token models.PasswordResetToken

	if err := repo.DB.Where("id_usuario=?", userID).Order("created_timestamp DESC").First(&token).Error; err != nil{
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &token, nil
}