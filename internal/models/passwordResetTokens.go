package models

import "time"

// ? Declaraion del modelo de referencia de la tabla password_reset_tokens
type PasswordResetToken struct {
	ID					int
	Token               string
	CreatedTimestamp    time.Time
	ExpirationTimestamp time.Time
	Used                bool
	IdUsuario           string
	Usuario             Usuario `gorm:"foreignKey: IdUsuario"`
}