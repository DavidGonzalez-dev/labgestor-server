package models

import "time"

type EtapaDetecciones struct {
	ID               int    `json:"id"`
	NombreEtapa      string `json:"nombreEtapa"`
	TiempoEtapa      string `json:"tiempoEtapa"`
	TemperaturaEtapa string `json:"temperaturaEtapa"`
}

type MonitoreosDeteccionesMicroorganismo struct {
	ID                        int             `gorm:"primaryKey autoincrement" json:"id"`
	VolumenMuestra            string          `json:"volumenMuestra"`
	NombreDiluyente           string          `json:"nombreDiluyente"`
	FechayhoraInicio          time.Time       `json:"fechayhoraInicio"`
	FechayhoraFinal           time.Time       `json:"fechayhoraFinal"`
	IdDeteccionMicroorganismo int             `json:"idDeteccionMicroorganismo"`
	IdEtapaDeteccion          int             `json:"idEtapaDeteccion"`
	EtapaDetecciones            *EtapaDetecciones `gorm:"foreignKey: IdEtapaDeteccion" json:"etapaDeteccion"`
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
