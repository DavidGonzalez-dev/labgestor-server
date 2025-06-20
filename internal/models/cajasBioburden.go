package models

type CajasBioburden struct {
	ID                   int             `gorm:"primaryKey autoincrement" json:"id"`
	Tipo                 string          `json:"tipo,omitempty"`
	Resultado            string          `json:"resultado"`
	MetodoSiembra        string          `json:"metodoSiembra,omitempty"`
	MedidaAritmetica     string          `json:"medidaAritmetica,omitempty"`
	FechayhoraIncubacion string          `json:"fechayhoraIncubacion,omitempty"`
	FechayhoraLectura    string          `json:"fechayhoraLectura,omitempty"`
	FactorDisolucion     string          `json:"factorDisolucion,omitempty"`
	IdPruebaRecuento     int             `json:"idPruebaRecuento,omitempty"`
	PruebaRecuento       *PruebaRecuento `gorm:"foreignKey:IdPruebaRecuento" json:"pruebaRecuento,omitempty"`
}

func (cajasBioburden CajasBioburden) ToMap() map[string]any {
	return map[string]any{
		"ID":                   cajasBioburden.ID,
		"Tipo":                 cajasBioburden.Tipo,
		"Resultado":            cajasBioburden.Resultado,
		"MetodoSiembra":        cajasBioburden.MetodoSiembra,
		"MedidaAritmetica":     cajasBioburden.MedidaAritmetica,
		"FechayhoraIncubacion": cajasBioburden.FechayhoraIncubacion,
		"FechayhoraLectura":    cajasBioburden.FechayhoraLectura,
		"FactorDisolucion":     cajasBioburden.FactorDisolucion,
		"IdPruebaRecuento":     cajasBioburden.IdPruebaRecuento,
	}
}
