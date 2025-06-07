package models

type PruebaRecuento struct {
	ID                     int       `gorm:"primaryKey" json:"id,omitempty"`
	MetodoUsado            string    `json:"metodoUsado,omitempty"`
	Concepto               bool      `json:"concepto,omitempty"`
	Especificacion         string    `json:"especificacion,omitempty"`
	VolumenDiluyente       string    `json:"volumenDiluyente,omitempty"`
	TiempoDisolucion       string    `json:"tiempoDisolucion,omitempty"`
	CantidadMuestra        string    `json:"cantidadMuestra,omitempty"`
	Tratamiento            string    `json:"tratamiento,omitempty"`
	NombreRecuento         string    `json:"nombreRecuento,omitempty"`
	NumeroRegistroProducto string    `json:"numeroRegistroProducto,omitempty"`
	Estado                 string    `json:"estado,omitempty"`
	Producto               *Producto `gorm:"foreignKey:NumeroRegistroProducto" json:"-"`
}

func (pruebaRecuento *PruebaRecuento) ToMap() map[string]any {
	return map[string]any{
		"ID":                     pruebaRecuento.ID,
		"MetodoUsado":            pruebaRecuento.MetodoUsado,
		"Concepto":               pruebaRecuento.Concepto,
		"Especificacion":         pruebaRecuento.Especificacion,
		"VolumenDiluyente":       pruebaRecuento.VolumenDiluyente,
		"TiempoDisolucion":       pruebaRecuento.TiempoDisolucion,
		"CantidadMuestra":        pruebaRecuento.CantidadMuestra,
		"Tratamiento":            pruebaRecuento.Tratamiento,
		"NombreRecuento":         pruebaRecuento.NombreRecuento,
		"NumeroRegistroProducto": pruebaRecuento.NumeroRegistroProducto,
	}
}
