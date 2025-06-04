package models

type DeteccionesMicroorganismos struct {
	ID                     int    `gorm:"primaryKey autoincrement" json:"id"`
	NombreMicroorganismo   string `json:"nombre_microorganismo"`
	Especificacion         string `json:"especificacion"`
	Concepto               bool   `json:"concepto"`
	Tratamiento            string `json:"tratamiento"`
	MetodoUsado            string `json:"metodo_usado"`
	CantidadMuestra        string `json:"cantidad_muestra"`
	VolumenDiluyente       string `json:"volumen_diluyente"`
	Resultado              string `json:"resultado"`
	NumeroRegistroProducto string `json:"numero_registro_producto"`
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
