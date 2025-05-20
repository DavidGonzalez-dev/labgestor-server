package models

type PruebaRecuento struct {
	ID                     int    `gorm:"primaryKey" json:"id"`
	MetodoUsado            string `json:"metodoUsado"`
	Concepto               bool   `json:"concepto"`
	Especificacion         string `json:"especificacion"`
	VolumenDiluyente       string `json:"volumenDiluyente"`
	TiempoDisolucion       string `json:"tiempoDisolucion"`
	CantidadMuestra        string `json:"cantidadMuestra"`
	Tratamiento            string `json:"tratamiento"`
	NombreRecuento         string `json:"nombreRecuento"`
	NumeroRegistroProducto string `json:"numeroRegistroProducto"`
}
