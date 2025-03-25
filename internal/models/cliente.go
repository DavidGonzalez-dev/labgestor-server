package models

// Declaracion del modelo de referencia de la tabla clientes
type Cliente struct {
	ID        string `gorm:"primaryKey" gorm:"autoincrement"` // Llave primaria
	Nombres   string
	Direccion string
}
