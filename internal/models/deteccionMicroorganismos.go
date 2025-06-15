package models

type DeteccionesMicroorganismos struct {
	ID                     int    `gorm:"primaryKey autoincrement" json:"id"`
	NombreMicroorganismo   string `json:"nombre_microorganismo,omitempty"`
	Especificacion         string `json:"especificacion,omitempty"`
	Concepto               bool   `json:"concepto,omitempty"`
	Tratamiento            string `json:"tratamiento,omitempty"`
	MetodoUsado            string `json:"metodo_usado,omitempty"`
	CantidadMuestra        string `json:"cantidad_muestra,omitempty"`
	VolumenDiluyente       string `json:"volumen_diluyente,omitempty"`
	Resultado              string `json:"resultado,omitempty"`
	NumeroRegistroProducto string `json:"numero_registro_producto,omitempty"`
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
