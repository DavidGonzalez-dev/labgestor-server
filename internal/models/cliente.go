package models


// Declaracion del modelo de referencia de la tabla clientes
type Cliente struct {
	ID        int    `gorm:"primaryKey autoincrement" json:"id"` // Llave primaria
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}

func (cliente Cliente) ToMap() map[string]any {
	return map[string]any{
		"Nombre":    cliente.Nombre,
		"Direccion": cliente.Direccion,
	}
}
