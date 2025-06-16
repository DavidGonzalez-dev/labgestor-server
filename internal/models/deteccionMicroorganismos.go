package models

type DeteccionesMicroorganismos struct {
	ID                     int    `gorm:"primaryKey autoincrement" json:"id"`
	NombreMicroorganismo   string `json:"nombreMicroorganismo,omitempty"`
	Especificacion         string `json:"especificacion,omitempty"`
	Concepto               bool   `json:"concepto"`
	Tratamiento            string `json:"tratamiento,omitempty"`
	MetodoUsado            string `json:"metodoUsado,omitempty"`
	CantidadMuestra        string `json:"cantidadMuestra,omitempty"`
	VolumenDiluyente       string `json:"volumenDiluyente,omitempty"`
	Resultado              string `json:"resultado"`
	NumeroRegistroProducto string `json:"numeroRegistroProducto,omitempty"`
}

func (deteccionMicroorganismos DeteccionesMicroorganismos) ToMap() map[string]any {
	return map[string]any{
		"ID":                     deteccionMicroorganismos.ID,
		"NombreMicroorganismo":   deteccionMicroorganismos.NombreMicroorganismo,
		"Especificacion":         deteccionMicroorganismos.Especificacion,
		"Concepto":               deteccionMicroorganismos.Concepto,
		"Tratamiento":            deteccionMicroorganismos.Tratamiento,
		"MetodoUsado":            deteccionMicroorganismos.MetodoUsado,
		"CantidadMuestra":        deteccionMicroorganismos.CantidadMuestra,
		"VolumenDiluyente":       deteccionMicroorganismos.VolumenDiluyente,
		"Resultado":              deteccionMicroorganismos.Resultado,
		"NumeroRegistroProducto": deteccionMicroorganismos.NumeroRegistroProducto,
	}
}
