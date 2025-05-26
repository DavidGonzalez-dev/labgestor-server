package models

type PruebaRecuento struct {
	ID                     int       `gorm:"primaryKey" json:"id"`
	MetodoUsado            string    `json:"metodoUsado"`
	Concepto               bool      `json:"concepto"`
	Especificacion         string    `json:"especificacion"`
	VolumenDiluyente       string    `json:"volumenDiluyente"`
	TiempoDisolucion       string    `json:"tiempoDisolucion"`
	CantidadMuestra        string    `json:"cantidadMuestra"`
	Tratamiento            string    `json:"tratamiento"`
	NombreRecuento         string    `json:"nombreRecuento"`
	NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
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
<<<<<<< HEAD
=======

>>>>>>> 0fee36fb23a0a4d5dd82e7dbec779d4f64aafe56
