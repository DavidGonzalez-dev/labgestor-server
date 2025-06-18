package models

type CajasBioburden struct {
	ID                   int             `gorm:"primaryKey autoincrement" json:"id"`
	Tipo                 string          `json:"tipo"`
	Resultado            string          `json:"resultado"`
	MetodoSiembra        string          `json:"metodoSiembra"`
	MedidaAritmetica     string          `json:"metodoAritmetica"`
	FechayhoraIncubacion string          `json:"fechayhoraIncubacion"`
	FechayhoraLectura    string          `json:"fechayhoraLectura"`
	FactorDisolucion     string          `json:"factorDisolucion"`
	IdPruebaRecuento     int             `json:"idPruebaRecuento"`
	PruebaRecuento       *PruebaRecuento `gorm:"foreignKey:IdPruebaRecuento" json:"pruebaRecuento"`
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
