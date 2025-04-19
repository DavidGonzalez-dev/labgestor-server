package models

// Declaracion del modelo de referencia de la tabla rol_usuarios
type rolUsuario struct {
	ID        int `gorm:"primaryKey" json:"-"`
	NombreRol string
}

// Declaracion del modelo de referencia de la tabla usuarios
type Usuario struct {
	ID         string     `gorm:"primaryKey,omitempty"` // Llave primaria
	Nombres    string     `json:"nombres,omitempty"`
	Apellidos  string     `json:"apellidos,omitempty"`
	Correo     string     `json:"correo,omitempty"`
	Contrasena string     `json:"contrasena,omitempty"`
	Firma      string     `json:"firma,omitempty"`
	Estado     bool       `json:"estado,omitempty"`
	RolID      int        `json:"-"`                 // Llave Secundaria
	Rol        *rolUsuario `gorm:"foreignKey: RolID" json:"rol,omitempty"` // Referecia a tabla llave foranea
}

func (usuario Usuario) ToMap() map[string]any {
	return map[string]any{
		"ID":         usuario.ID,
		"Nombres":    usuario.Nombres,
		"Apellidos":  usuario.Apellidos,
		"Correo":     usuario.Correo,
		"Contrasena": usuario.Contrasena,
		"Firma":      usuario.Firma,
		"Estado":     usuario.Estado,
		"RolID":      usuario.RolID,
	}
}
