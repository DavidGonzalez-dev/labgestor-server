package models

// Declaracion del modelo de referencia de la tabla rol_usuarios
type rolUsuario struct {
	ID        int `gorm:"primaryKey" json:"-"`
	NombreRol string
}

// Declaracion del modelo de referencia de la tabla usuarios
type Usuario struct {
	ID         string `gorm:"primaryKey"` // Llave primaria
	Nombres    string
	Apellidos  string
	Correo     string
	Contrasena string
	Firma      string
	Estado     bool
	RolID      int        `json:"-"`                 // Llave Secundaria
	Rol        rolUsuario `gorm:"foreignKey: RolID"` // Referecia a tabla llave foranea
}
