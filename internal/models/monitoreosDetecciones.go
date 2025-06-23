package models

type MonitoreosDeteccionesMicroorganismo struct {
	ID                        int                         `gorm:"primaryKey autoincrement" json:"id"` // Llave primaria
	VolumenMuestra            string                      `json:"volumenMuestra"`
	NombreDiluyente           string                      `json:"nombreDiluyente"`
	FechayhoraInicio          string                      `json:"fechayhoraInicio"`
	FechayhoraFinal           string                      `json:"fechayhoraFinal"`
	IdEtapaDeteccion          int                         `json:"idEtapaDeteccion"`
	IdDeteccionMicroorganismo int                         `json:"idDeteccionMicroorganismo"`
	DeteccionMicroorganismo   *DeteccionesMicroorganismos `gorm:"foreignKey:IdDeteccionMicroorganismo" json:"deteccionMicroorganismo"`
}

func (monitoreosDetecciones MonitoreosDeteccionesMicroorganismo) ToMap() map[string]any {
	return map[string]any{
		"ID":                        monitoreosDetecciones.ID,
		"VolumenMuestra":            monitoreosDetecciones.VolumenMuestra,
		"NombreDiluyente":           monitoreosDetecciones.NombreDiluyente,
		"FechayhoraInicio":          monitoreosDetecciones.FechayhoraInicio,
		"FechayhoraFinal":           monitoreosDetecciones.FechayhoraFinal,
		"IdEtapaDeteccion":          monitoreosDetecciones.IdEtapaDeteccion,
		"IdDeteccionMicroorganismo": monitoreosDetecciones.IdDeteccionMicroorganismo,
	}
}
