package models

// Declaracion del modelo de referencia de la tabla fabricantes
type Fabricante struct {
	ID        int `gorm:"primaryKey autoincrement"` // Llave primaria
	Nombre    string
	Direccion string
}

func (fabricante Fabricante) ToMap()  map[string]any{
	return map[string]any{
		"ID" : fabricante.ID,
		"Nombre" : fabricante.Nombre,
		"Direccion" :fabricante.Direccion,
	}
}