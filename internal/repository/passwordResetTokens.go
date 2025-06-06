package repository

import (
	"labgestor-server/internal/models"
	"time"

	"gorm.io/gorm"
)

// Definimos la interfaz
type PasswordResetTokenRepository interface {
    Create(token *models.PasswordResetToken) error
    FindByToken(tokenValue string) (*models.PasswordResetToken, error)
    MarkAsUsed(tokenID string) error
    DeleteExpired() error
    DeleteByUserID(userID string) error
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
	return repo.DB.Create(&token). Error
}

// Funcion para obtener un registro por token
func (repo *passwordResetTokenRepo) FindByToken(tokenValue string) (*models.PasswordResetToken, error) {
	var token models.PasswordResetToken
	
	if err := repo.DB.Where("token=?", tokenValue).First(&token). Error ; err != nil {
		return nil, err
	}
	return &token, nil
}

// Funcion para marcar un token como usado basado en el id del token
func (repo *passwordResetTokenRepo) MarkAsUsed(tokenID string) error {
	// Obtenemos el registro del token desde la base de datos
	var token models.PasswordResetToken
	if err := repo.DB.First(&token, tokenID).Error; err != nil{
		return err
	}

	// Lo marcamos como usado
	token.Used = true
	
	// Guardamos los cambios
	if err := repo.DB.Save(token).Error; err != nil{
		return err
	}
	
	return nil
}

// Funcion para Eliminar todos los tokens expirados
func (repo *passwordResetTokenRepo) DeleteExpired() error {
	return repo.DB.Exec("DELETE FROM password_reset_tokens where expired_timestamp > ?", time.Now().UTC()).Error
}

// Funcion para eliminar los tokens de un usario en especifico
func (repo *passwordResetTokenRepo) DeleteByUserID(userID string) error {
	return repo.DB.Exec("DELETE FROM password_reset_tokens WHERE id_usuario=?", userID).Error
}