package models

import "time"

type ControlesNegativosMedio struct {
	ID                     int       `gorm:"primaryKey" json:"id"`
	MedioCultivo           string    `json:"medioCultivo"`
	FechayhoraIncubacion   time.Time `json:"fechayhoraIncubacion"`
	FechayhoraLectura      time.Time `json:"fechayhoraLectura"`
	Resultado              string    `json:"resultado"`
	NumeroRegistroProducto string    `json:"numeroRegistro"`
}

func (controlesNegativos *ControlesNegativosMedio) ToMap() map[string]any {
	return map[string]any{
		"ID":                     controlesNegativos.ID,
		"MedioCultivo":           controlesNegativos.MedioCultivo,
		"FechayhoraIncubacion":   controlesNegativos.FechayhoraIncubacion,
		"FechayhoraLectura":      controlesNegativos.FechayhoraLectura,
		"Resultado":              controlesNegativos.Resultado,
		"NumeroRegistroProducto": controlesNegativos.NumeroRegistroProducto,
	}
}
