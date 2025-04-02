package models

// Declaracion del modelo de referencia de la tabla fabricantes
type Fabricante struct {
	ID        int `gorm:"primaryKey autoincrement"` // Llave primaria
	Nombre    string
	Direccion string
}
