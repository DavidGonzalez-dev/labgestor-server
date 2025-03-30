package models

// Declaracion del modelo de referencia de la tabla clientes
type Cliente struct {
	ID        int `gorm:"primaryKey" gorm:"autoincrement"` // Llave primaria
	Nombre    string
	Direccion string
}
