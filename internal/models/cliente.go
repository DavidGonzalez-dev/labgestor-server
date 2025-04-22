package models

// Declaracion del modelo de referencia de la tabla clientes
type Cliente struct {
	ID        int `gorm:"primaryKey autoincrement"` // Llave primaria
	Nombre    string
	Direccion string
}

func (cliente Cliente) ToMap() map[string]any {
	return map[string]any {
		"Nombre": cliente.Nombre,
		"Direccion": cliente.Direccion,
	}
}